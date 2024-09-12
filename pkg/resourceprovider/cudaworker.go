//go:build cuda
// +build cuda

package resourceprovider

import (
	"context"
	_ "embed"
	"fmt"
	"math/big"
	"os"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gorgonia.org/cu"
)

//go:embed cudaminer/keccak.ptx
var keccakPtx string

const entry_point = "kernel_lilypad_pow"

var MaybeCudaOrCpu = NewGpuWorker

func DefaultWorkerNum() int {
	return 1
}

type GpuWorker struct {
	cfg     *WorkerConfig
	state   atomic.Int32
	entryFn cu.Function
	cuCtx   *cu.Ctx

	quit chan chan struct{}
}

func NewGpuWorker(cfg *WorkerConfig) (Worker, error) {
	if GetGpuNumber() == 0 {
		log.Warn().Msg("No gpu found and fallback to cpu")
		return NewCpuWorker(cfg)
	}
	//TODO use first gpu for now, plan to support multiple gpu in future
	cuCtx, err := setupGPU()
	if err != nil {
		return nil, err
	}

	fs, err := os.CreateTemp(os.TempDir(), "*")
	if err != nil {
		return nil, err
	}

	//TODO no LoadData for Cuda Ctx, maybe cu author forget or just not supprot
	_, err = fs.WriteString(keccakPtx)
	if err != nil {
		return nil, err
	}
	myModule, err := cuCtx.Load(fs.Name())
	if err != nil {
		return nil, err
	}

	entryFn, err := myModule.Function(entry_point)
	if err != nil {
		return nil, err
	}

	return &GpuWorker{
		cfg:     cfg,
		cuCtx:   cuCtx,
		entryFn: entryFn,
		quit:    make(chan chan struct{}, 1),
	}, nil
}
func (w *GpuWorker) Stop() {
	if w.state.Load() == 0 {
		return
	}

	//cancel previous task
	waitChan := make(chan struct{})
	select {
	case w.quit <- waitChan: // may already exit
		<-waitChan
	default:
	}
}
func (w *GpuWorker) FindSolution(ctx context.Context, task *Task) {
	w.state.Store(1)
	defer w.state.Store(0)

	nonce := task.From.Clone()
	startTime := time.Now()

	hashesCompleted := uint64(0)
	ticker := time.NewTicker(time.Second * hashUpdateSecs)
	defer ticker.Stop()
	batch_size := uint64(w.cfg.gridSize * w.cfg.blockSize / 32)
OUT:
	for {
		select {
		case <-ctx.Done():
			break OUT
		case respCh := <-w.quit:
			respCh <- struct{}{}
			return
		case <-ticker.C:
			w.cfg.updateHashes <- hashesCompleted
			hashesCompleted = 0
		default:
			// Non-blocking select to fall through
		}

		if nonce.Cmp(task.End) >= 0 {
			return
		}

		maybeNonce, err := Kernel_lilypad_pow_with_ctx(w.cuCtx, w.entryFn, task.Challenge, nonce.ToBig(), task.Difficulty.ToBig(), w.cfg.gridSize, w.cfg.blockSize, w.cfg.hashsPerThread)
		if err != nil {
			log.Err(err).Msg("InvokeGpu fail")
			continue
		}
		hashesCompleted += batch_size
		nonce = nonce.Add(nonce, uint256.NewInt(batch_size))

		if maybeNonce.Int64() == 0 {
			continue
		}

		data, err := formatMinerArgs(task.Challenge, maybeNonce)
		if err != nil {
			log.Err(err).Msg("Generate hash data")
			continue
		}
		result := crypto.Keccak256Hash(data)

		hashNumber := new(uint256.Int).SetBytes(result[:])
		// Check if the hash is below the target difficulty
		if hashNumber.Cmp(task.Difficulty) == -1 {
			w.cfg.updateHashes <- hashesCompleted // report hash to avoid too fast
			hashesCompleted = 0
			log.Info().Int("WorkerID", w.cfg.id).Str("Elapsed Time", time.Since(startTime).String()).
				Str("challenge", new(big.Int).SetBytes(task.Challenge[:]).String()).
				Str("Nonce", maybeNonce.String()).
				Str("HashNumber", hashNumber.String()).
				Msg("Success!")
			select {
			case w.cfg.resultCh <- TaskResult{
				Id:    task.Id,
				Nonce: uint256.MustFromBig(maybeNonce),
			}:
			default: //avoid deadlock
			}
		} else {
			log.Error().Msg("This branch should never happen, only when cuda algo may have error")
		}
	}
}

func GetGpuNumber() int {
	devices, err := cu.NumDevices()
	if err != nil {
		log.Warn().Msgf("Cannot detect gpu numbers %v", err)
		return 0
	}
	return devices
}

func setupGPU() (*cu.Ctx, error) {
	devices, _ := cu.NumDevices()

	if devices == 0 {
		return nil, errors.Errorf("NoDevice")
	}

	dev := cu.Device(0)
	return cu.NewContext(dev, cu.SchedAuto), nil
}

func Kernel_lilypad_pow_with_ctx(cuCtx *cu.Ctx, fn cu.Function, challenge [32]byte, startNonce *big.Int, difficulty *big.Int, grid, block int, hashPerThread int) (*big.Int, error) {
	dIn1, err := cuCtx.MemAllocManaged(32, cu.AttachGlobal)
	if err != nil {
		return nil, err
	}

	dIn2, err := cuCtx.MemAllocManaged(32, cu.AttachGlobal)
	if err != nil {
		return nil, err
	}

	dIn3, err := cuCtx.MemAllocManaged(32, cu.AttachGlobal)
	if err != nil {
		return nil, err
	}

	dOut, err := cuCtx.MemAllocManaged(32, cu.AttachGlobal)
	if err != nil {
		return nil, err
	}

	batch := int64(grid * block)
	//(BYTE* indata,	 WORD inlen,	 BYTE* outdata,	 WORD n_batch,	 WORD KECCAK_BLOCK_SIZE)
	args := []unsafe.Pointer{
		unsafe.Pointer(&dIn1),
		unsafe.Pointer(&dIn2),
		unsafe.Pointer(&dIn3),
		unsafe.Pointer(&batch),
		unsafe.Pointer(&hashPerThread),
		unsafe.Pointer(&dOut),
	}

	cuCtx.MemcpyHtoD(dIn1, unsafe.Pointer(&challenge[0]), 32)

	startNonceBytes := math.U256Bytes(startNonce)
	cuCtx.MemcpyHtoD(dIn2, unsafe.Pointer(&startNonceBytes[0]), 32)

	difficutyBytes := math.U256Bytes(difficulty)
	cuCtx.MemcpyHtoD(dIn3, unsafe.Pointer(&difficutyBytes[0]), 32)

	cuCtx.LaunchKernel(fn, grid, 1, 1, block, 1, 1, 1, cu.Stream{}, args)
	if err = cuCtx.Error(); err != nil {
		return nil, fmt.Errorf("launch kernel fail maybe decrease threads help %w", err)
	}
	cuCtx.Synchronize()

	hOut := make([]byte, 32)
	cuCtx.MemcpyDtoH(unsafe.Pointer(&hOut[0]), dOut, 32)

	cuCtx.MemFree(dIn1)
	cuCtx.MemFree(dIn2)
	cuCtx.MemFree(dIn3)
	cuCtx.MemFree(dOut)

	return new(big.Int).SetBytes(hOut), nil
}

package resourceprovider

import (
	"context"
	_ "embed"
	"math/big"
	"os"
	"sync/atomic"
	"time"
	"unsafe"

	"github.com/holiman/uint256"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"gorgonia.org/cu"
)

//go:embed cudaminer/keccak.ptx
var keccakPtx string

const entry_point = "kernel_keccak_hash"
const batch_size = 1000

type GpuWorker struct {
	cfg     *WorkerConfig
	state   atomic.Int32
	entryFn cu.Function
	cuCtx   *cu.Ctx
	quit    chan chan struct{}
}

func NewGpuWorker(cfg *WorkerConfig) (*GpuWorker, error) {
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
	module, err := cuCtx.Load(fs.Name())
	if err != nil {
		return nil, err
	}

	entryFn, err := module.Function(entry_point)
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

		//aggregate input
		inputs := make([][64]byte, batch_size)
		for i := 0; i < batch_size; i++ {
			data, err := formatMinerArgs(task.Challenge, nonce.ToBig())
			if err != nil {
				log.Err(err).Msg("Generate hash data")
				continue
			}
			inputs[i] = [64]byte(data)
			nonce.Add(nonce, bigOne)
		}

		results, err := cuda_batch_keccak(w.cuCtx, w.entryFn, inputs)
		if err != nil {
			log.Err(err).Msg("InvokeGpu fail")
			continue
		}

		/*
			for index, result := range results {
				hash := crypto.Keccak256Hash(inputs[index][:])
				if !bytes.Equal(hash[:], result[:]) {
					panic("hash not match")
				}
			}
		*/
		hashesCompleted += batch_size
		for _, result := range results {
			hashNumber := new(uint256.Int).SetBytes(result[:])
			// Check if the hash is below the target difficulty
			if hashNumber.Cmp(task.Difficulty) == -1 {
				log.Info().Str("Elapsed Time", time.Since(startTime).String()).
					Str("challenge", new(big.Int).SetBytes(task.Challenge[:]).String()).
					Str("Nonce", nonce.String()).
					Str("HashNumber", hashNumber.String()).
					Msg("Success!")
				w.cfg.resultCh <- TaskResult{
					Id:    task.Id,
					Nonce: nonce.Clone(),
				}
			}
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

func cuda_batch_keccak(cuCtx *cu.Ctx, fn cu.Function, hIn [][64]byte) ([][32]byte, error) {
	inNum := int64(len(hIn))

	dIn, err := cuCtx.MemAllocManaged(64*inNum, cu.AttachGlobal)
	if err != nil {
		return nil, err
	}

	dOut, err := cuCtx.MemAllocManaged(32*inNum, cu.AttachGlobal)
	if err != nil {
		return nil, err
	}

	inLen := int64(64)
	block_size := int64(256 >> 3)
	//(BYTE* indata,	 WORD inlen,	 BYTE* outdata,	 WORD n_batch,	 WORD KECCAK_BLOCK_SIZE)
	args := []unsafe.Pointer{
		unsafe.Pointer(&dIn),
		unsafe.Pointer(&inLen),
		unsafe.Pointer(&dOut),

		unsafe.Pointer(&inNum),
		unsafe.Pointer(&block_size),
	}

	cuCtx.MemcpyHtoD(dIn, unsafe.Pointer(&hIn[0]), 64*inNum)

	thread := 256
	block := (int(inNum) + thread - 1) / thread //todo this argument maybe need to change
	cuCtx.LaunchKernel(fn, thread, 1, 1, block, 1, 1, 1, cu.Stream{}, args)
	cuCtx.Synchronize()
	hOut := make([][32]byte, inNum)
	cuCtx.MemcpyDtoH(unsafe.Pointer(&hOut[0]), dOut, 32*inNum)

	cuCtx.MemFree(dIn)
	cuCtx.MemFree(dOut)
	return hOut, nil
}

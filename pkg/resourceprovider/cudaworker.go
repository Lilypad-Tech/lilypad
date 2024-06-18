package resourceprovider

import (
	"context"
	_ "embed"
	"math/big"
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
	quit    chan chan struct{}
}

func NewGpuWorker(cfg *WorkerConfig) (*GpuWorker, error) {
	//TODO use first gpu for now, plan to support multiple gpu in future
	_, _, err := setupGPU()
	if err != nil {
		return nil, err
	}
	module, err := cu.LoadData(keccakPtx)
	if err != nil {
		return nil, err
	}

	entryFn, err := module.Function(entry_point)
	if err != nil {
		return nil, err
	}
	return &GpuWorker{
		cfg:     cfg,
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

		results, err := cuda_batch_keccak(w.entryFn, inputs)
		if err != nil {
			log.Err(err).Msg("InvokeGpu fail")
			continue
		}
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

func setupGPU() (dev cu.Device, ctx cu.CUContext, err error) {
	devices, _ := cu.NumDevices()

	if devices == 0 {
		err = errors.Errorf("NoDevice")
		return
	}

	dev = cu.Device(0) //todo support multiple gpu
	if ctx, err = dev.MakeContext(cu.SchedAuto); err != nil {
		return
	}
	return
}

func cuda_batch_keccak(fn cu.Function, hIn [][64]byte) ([][32]byte, error) {
	inNum := int64(len(hIn))

	dIn, err := cu.MemAlloc(64 * inNum)
	if err != nil {
		return nil, err
	}

	dOut, err := cu.MemAlloc(32 * inNum)
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

	if err = cu.MemcpyHtoD(dIn, unsafe.Pointer(&hIn[0]), 64*inNum); err != nil {
		return nil, err
	}

	thread := 256
	block := (int(inNum) + thread - 1) / thread
	if err = fn.LaunchAndSync(thread, 1, 1, block, 1, 1, 1, cu.Stream{}, args); err != nil {
		return nil, err
	}

	hOut := make([][32]byte, inNum)
	if err = cu.MemcpyDtoH(unsafe.Pointer(&hOut[0]), dOut, 32*inNum); err != nil {
		return nil, err
	}

	cu.MemFree(dIn)
	cu.MemFree(dOut)
	return hOut, nil
}

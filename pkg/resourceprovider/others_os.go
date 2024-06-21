//go:build !cuda
// +build !cuda

package resourceprovider

import "runtime"

var MaybeCudaOrCpu = NewCpuWorker

func DefaultWorkerNum() int {
	return runtime.NumCPU() * 2
}

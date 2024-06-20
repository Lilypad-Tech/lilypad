//go:build !linux
// +build !linux

package resourceprovider

import "runtime"

var MaybeCudaOrCpu = NewCpuWorker

func DefaultWorkerNum() int {
	return runtime.NumCPU() * 2
}

//go:build !linux
// +build !linux

package resourceprovider

var MaybeCudaOrCpu = NewCpuWorker

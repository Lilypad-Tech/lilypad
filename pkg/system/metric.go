package system

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/process"
	"go.opentelemetry.io/otel/metric"
)

// Use the official OTel semantic conventions where it makes sense to:
// - System: https://opentelemetry.io/docs/specs/semconv/system/system-metrics/
// - Process: https://opentelemetry.io/docs/specs/semconv/system/process-metrics/
// - Network: https://opentelemetry.io/docs/specs/semconv/system/system-metrics/#network-metrics
func NewMetrics(meter metric.Meter) (func() error, error) {
	// General
	uptime, err := meter.Int64ObservableCounter(
		"system.uptime",
		metric.WithDescription("System uptime."),
		metric.WithUnit("s"),
	)
	if err != nil {
		return nil, err
	}
	load1, err := meter.Float64ObservableGauge(
		"system.linux.cpu.load_1m",
		metric.WithDescription("The average number of processes in the queue during the last minute."),
	)
	if err != nil {
		return nil, err
	}
	load5, err := meter.Float64ObservableGauge(
		"system.linux.cpu.load_5m",
		metric.WithDescription("The average number of processes in the queue during the last five minutes."),
	)
	if err != nil {
		return nil, err
	}
	load15, err := meter.Float64ObservableGauge(
		"system.linux.cpu.load_15m",
		metric.WithDescription("The average number of processes in the queue during the last fifteen minutes."),
	)
	if err != nil {
		return nil, err
	}

	// CPU
	cpuUsedPercent, err := meter.Float64ObservableGauge(
		"system.cpu.used_percent",
		metric.WithDescription("The average percentage of CPU usage over all CPUs."),
		metric.WithUnit("%"),
	)
	if err != nil {
		return nil, err
	}
	cpuCount, err := meter.Int64ObservableGauge(
		"system.cpu.logical.count",
		metric.WithDescription("The number of logical CPU cores."),
	)
	if err != nil {
		return nil, err
	}

	// Memory
	memoryUsedPercent, err := meter.Float64ObservableGauge(
		"system.memory.used_percent",
		metric.WithDescription("Percentage of RAM used by programs."),
		metric.WithUnit("%"),
	)
	if err != nil {
		return nil, err
	}
	memoryUsed, err := meter.Int64ObservableGauge(
		"system.memory.used",
		metric.WithDescription("RAM used by programs."),
		metric.WithUnit("By"),
	)
	if err != nil {
		return nil, err
	}
	memoryAvailable, err := meter.Int64ObservableGauge(
		"system.memory.available",
		metric.WithDescription("RAM available for programs to allocate."),
		metric.WithUnit("By"),
	)
	if err != nil {
		return nil, err
	}
	memoryTotal, err := meter.Int64ObservableGauge(
		"system.memory.limit",
		metric.WithDescription("Total amount of RAM on this system."),
		metric.WithUnit("By"),
	)
	if err != nil {
		return nil, err
	}

	// Network
	networkConnectionsCount, err := meter.Int64ObservableGauge(
		"network.connections.count",
		metric.WithDescription("The number of connections the system is actively using."),
	)
	if err != nil {
		return nil, err
	}
	networkBytesSent, err := meter.Int64ObservableGauge(
		"network.io.transmit",
		metric.WithDescription("The number of bytes sent by the system."),
		metric.WithUnit("By"),
	)
	if err != nil {
		return nil, err
	}
	networkBytesReceived, err := meter.Int64ObservableGauge(
		"network.io.receive",
		metric.WithDescription("The number of bytes received by the system."),
		metric.WithUnit("By"),
	)
	if err != nil {
		return nil, err
	}

	// Filesystem
	filesystemUsedPercent, err := meter.Float64ObservableGauge(
		"system.filesystem.used_percent",
		metric.WithDescription("Used filesystem space percent."),
		metric.WithUnit("%"),
	)
	if err != nil {
		return nil, err
	}
	filesystemUsed, err := meter.Int64ObservableGauge(
		"system.filesystem.used",
		metric.WithDescription("Used filesystem space."),
		metric.WithUnit("By"),
	)
	if err != nil {
		return nil, err
	}
	filesystemFree, err := meter.Int64ObservableGauge(
		"system.filesystem.free",
		metric.WithDescription("Free filesystem space."),
		metric.WithUnit("By"),
	)
	if err != nil {
		return nil, err
	}
	filesystemTotal, err := meter.Int64ObservableGauge(
		"system.filesystem.limit",
		metric.WithDescription("Total filesystem space."),
		metric.WithUnit("By"),
	)
	if err != nil {
		return nil, err
	}
	inodesUsedPercent, err := meter.Float64ObservableGauge(
		"system.filesystem.inodes.used_percent",
		metric.WithDescription("Inodes used percent."),
		metric.WithUnit("%"),
	)
	if err != nil {
		return nil, err
	}

	// Process
	processUptime, err := meter.Int64ObservableCounter(
		"process.uptime",
		metric.WithDescription("The time this process has been running."),
		metric.WithUnit("s"),
	)
	if err != nil {
		return nil, err
	}
	processMemoryUsed, err := meter.Float64ObservableGauge(
		"process.memory.used",
		metric.WithDescription("The percent of the total RAM this process uses."),
	)
	if err != nil {
		return nil, err
	}
	processCPUUsed, err := meter.Float64ObservableGauge(
		"process.cpu.used",
		metric.WithDescription("The percent of CPU time this process uses."),
	)
	if err != nil {
		return nil, err
	}
	processConnectionsCount, err := meter.Int64ObservableGauge(
		"process.connections.count",
		metric.WithDescription("The number of connections this process is actively using."),
	)
	if err != nil {
		return nil, err
	}
	processThreadsCount, err := meter.Int64ObservableGauge(
		"process.thread.count",
		metric.WithDescription("The number of threads this process is using."),
	)
	if err != nil {
		return nil, err
	}

	// Register callback to observe metrics regularly
	registration, err := meter.RegisterCallback(
		func(_ context.Context, o metric.Observer) error {
			// General
			u, err := host.Uptime()
			if err != nil {
				log.Warn().Msgf("failed to collect uptime metric: %s", err)
				return err
			}
			o.ObserveInt64(uptime, int64(u))

			load, err := load.Avg()
			if err != nil {
				log.Warn().Msgf("failed to collect average load metric: %s", err)
				return err
			}
			o.ObserveFloat64(load1, load.Load1)
			o.ObserveFloat64(load5, load.Load5)
			o.ObserveFloat64(load15, load.Load15)

			// CPU
			cpus, err := cpu.Percent(0, false)
			if err != nil {
				log.Warn().Msgf("failed to collect CPU percent metric: %s", err)
				return err
			}
			o.ObserveFloat64(cpuUsedPercent, cpus[0])

			count, err := cpu.Counts(true)
			if err != nil {
				log.Warn().Msgf("failed to collect CPU count: %s", err)
				return err
			}
			o.ObserveInt64(cpuCount, int64(count))

			// Memory
			memory, err := mem.VirtualMemory()
			if err != nil {
				log.Warn().Msgf("failed to collect virtual memory metrics: %s", err)
				return err
			}
			o.ObserveInt64(memoryTotal, int64(memory.Total))
			o.ObserveInt64(memoryAvailable, int64(memory.Available))
			o.ObserveInt64(memoryUsed, int64(memory.Used))
			o.ObserveFloat64(memoryUsedPercent, memory.UsedPercent)

			// Network
			networkConnections, err := net.Connections("all")
			if err != nil {
				log.Warn().Msgf("failed to collect network connections metric: %s", err)
				return err
			}
			// net.IOCounters returns combined counters in one entry when pernic false
			networkIOCounters, err := net.IOCounters(false)
			if err != nil {
				log.Warn().Msgf("failed to collect network I/O counters: %s", err)
				return err
			}
			o.ObserveInt64(networkConnectionsCount, int64(len(networkConnections)))
			o.ObserveInt64(networkBytesSent, int64(networkIOCounters[0].BytesSent))
			o.ObserveInt64(networkBytesReceived, int64(networkIOCounters[0].BytesRecv))

			// Filesystem
			disk, err := disk.Usage("/")
			if err != nil {
				log.Warn().Msgf("failed to collect disk usage metrics: %s", err)
				return err
			}
			o.ObserveInt64(filesystemTotal, int64(disk.Total))
			o.ObserveInt64(filesystemFree, int64(disk.Free))
			o.ObserveInt64(filesystemUsed, int64(disk.Used))
			o.ObserveFloat64(filesystemUsedPercent, disk.UsedPercent)
			o.ObserveFloat64(inodesUsedPercent, disk.InodesUsedPercent)

			// Process
			pid := os.Getpid()
			p, err := process.NewProcess(int32(pid))
			if err != nil {
				log.Warn().Msgf("failed to collect process info for pid %d: %s", pid, err)
				return err
			}
			pCreateTime, err := p.CreateTime()
			if err != nil {
				log.Warn().Msgf("failed to collect create time for pid %d: %s", pid, err)
				return err
			}
			// Compute uptime, convert and truncate to seconds
			pUptime := (time.Now().UnixMilli() - pCreateTime) / 1000
			pMemoryUsed, err := p.MemoryPercent()
			if err != nil {
				log.Warn().Msgf("failed to collect memory percent for pid %d: %s", pid, err)
				return err
			}
			pCPUUsed, err := p.CPUPercent()
			if err != nil {
				log.Warn().Msgf("failed to collect CPU percent for pid %d: %s", pid, err)
				return err
			}
			pConnections, err := p.Connections()
			if err != nil {
				log.Warn().Msgf("failed to collect connections metrics for pid %d: %s", pid, err)
				return err
			}
			pThreadsCount, err := p.NumThreads()
			if err != nil {
				log.Warn().Msgf("failed to collect num threads metric for pid %d: %s", pid, err)
				return err
			}
			o.ObserveInt64(processUptime, pUptime)
			o.ObserveFloat64(processMemoryUsed, float64(pMemoryUsed))
			o.ObserveFloat64(processCPUUsed, float64(pCPUUsed))
			o.ObserveInt64(processConnectionsCount, int64(len(pConnections)))
			o.ObserveInt64(processThreadsCount, int64(pThreadsCount))

			return nil
		},
		uptime,
		load1,
		load5,
		load15,
		cpuUsedPercent,
		cpuCount,
		memoryTotal,
		memoryAvailable,
		memoryUsed,
		memoryUsedPercent,
		filesystemTotal,
		filesystemFree,
		filesystemUsed,
		filesystemUsedPercent,
		inodesUsedPercent,
		networkConnectionsCount,
		networkBytesSent,
		networkBytesReceived,
		processUptime,
		processMemoryUsed,
		processCPUUsed,
		processConnectionsCount,
		processThreadsCount,
	)
	if err != nil {
		fmt.Println("failed to register system metrics callback")
		panic(err)
	}

	return registration.Unregister, nil
}

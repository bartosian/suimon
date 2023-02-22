package utility

import (
	"context"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/dariubs/percent"
	"github.com/docker/docker/client"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

const gb = 1024 * 1024 * 1024

type (
	usageData struct {
		Total          int
		Free           int
		Used           int
		PercentageUsed int
	}
	DiskUsage    usageData
	MemoryUsage  usageData
	CPUUsage     usageData
	NetworkUsage struct {
		Recv float64
		Sent float64
	}
)

func GetDiskUsage() (*DiskUsage, error) {
	var (
		stat syscall.Statfs_t
		err  error
	)

	if err = syscall.Statfs("/", &stat); err != nil {
		return nil, err
	}

	total := int(stat.Blocks*uint64(stat.Bsize)) / gb
	free := int(stat.Bfree*uint64(stat.Bsize)) / gb
	used := total - free
	percentageUsed := int(percent.PercentOf(used, total))

	return &DiskUsage{
		Total:          total,
		Free:           free,
		Used:           used,
		PercentageUsed: percentageUsed,
	}, nil
}

func GetDirSize(path string) (float64, error) {
	var size float64

	err := filepath.Walk(path, func(subPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += float64(info.Size())
		}
		return nil
	})

	if size != 0 {
		size = size / gb
	}

	return size, err
}

func GetVolumeSize(volumeName string) (float64, error) {
	var size float64

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return size, err
	}

	volume, err := cli.VolumeInspect(context.Background(), volumeName)
	if err != nil {
		return size, err
	}

	return float64(volume.UsageData.Size) / gb, nil
}

func GetNetworkUsage() (*NetworkUsage, error) {
	stat, err := net.IOCounters(false)
	if err != nil {
		return nil, err
	}

	recv := float64(stat[0].BytesRecv) / gb
	sent := float64(stat[0].BytesSent) / gb

	return &NetworkUsage{
		Recv: recv,
		Sent: sent,
	}, nil
}

func GetMemoryUsage() (*MemoryUsage, error) {
	var (
		stat *mem.VirtualMemoryStat
		err  error
	)

	if stat, err = mem.VirtualMemory(); err != nil {
		return nil, err
	}

	total := int(stat.Total) / gb
	free := int(stat.Free) / gb
	used := int(stat.Used) / gb
	percentageUsed := int(percent.PercentOf(used, total))

	return &MemoryUsage{
		Total:          total,
		Free:           free,
		Used:           used,
		PercentageUsed: percentageUsed,
	}, nil
}

func GetCPUUsage() (*CPUUsage, error) {
	var (
		cores      int
		percentage []float64
		err        error
	)

	if cores, err = cpu.Counts(false); err != nil {
		return nil, err
	}

	if percentage, err = cpu.Percent(time.Second, false); err != nil {
		return nil, err
	}

	pct := int(percentage[0])

	total := cores * 100
	free := total - int(percent.Percent(pct, total))
	used := total - free

	return &CPUUsage{
		Total:          total,
		Free:           free,
		Used:           used,
		PercentageUsed: pct,
	}, nil
}

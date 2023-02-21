package utility

import (
	"context"
	"os"
	"path/filepath"
	"syscall"

	"github.com/dariubs/percent"
	"github.com/docker/docker/client"
)

const gb = 1024 * 1024 * 1024

type DiskUsage struct {
	Total          int
	Free           int
	Used           int
	PercentageUsed int
}

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

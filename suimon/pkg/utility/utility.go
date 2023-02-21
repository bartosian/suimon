package utility

import (
	"syscall"

	"github.com/dariubs/percent"
)

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

	total := int(stat.Blocks*uint64(stat.Bsize)) / (1024 * 1024 * 1024)
	free := int(stat.Bfree*uint64(stat.Bsize)) / (1024 * 1024 * 1024)
	used := total - free
	percentageUsed := int(percent.PercentOf(used, total))

	return &DiskUsage{
		Total:          total,
		Free:           free,
		Used:           used,
		PercentageUsed: percentageUsed,
	}, nil
}

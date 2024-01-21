package utility

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"time"

	"github.com/dariubs/percent"
	"github.com/docker/docker/client"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

const gb = 1024 * 1024 * 1024
const cpuMax = 100
const millisecondsInSecond = 1000
const secondsInMinute = 60

type UsageData struct {
	Total          int
	Free           int
	Used           int
	PercentageUsed int
}

type NetworkUsage struct {
	Recv float64
	Sent float64
}

// GetDiskUsage returns the disk usage information for the root directory.
// It calculates the total, free, used, and percentage used space on the disk.
// It returns a pointer to UsageData and an error if any.
func GetDiskUsage() (*UsageData, error) {
	var (
		stat *disk.UsageStat
		err  error
	)

	if stat, err = disk.Usage("/"); err != nil {
		return nil, err
	}

	total := int(stat.Total) / gb
	free := int(stat.Free) / gb
	used := int(stat.Used) / gb
	percentageUsed := int(percent.PercentOf(used, total))

	return &UsageData{
		Total:          total,
		Free:           free,
		Used:           used,
		PercentageUsed: percentageUsed,
	}, nil
}

// GetDirSize returns the size of the directory at the specified path.
// It takes the path of the directory as input and returns the size of the directory in gigabytes and an error if any.
// It walks through the directory and sums up the size of all the files within it.
// If an error occurs during the process, it returns the error.
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
		size /= gb
	}

	return size, err
}

// GetVolumeSize returns the size of the specified volume.
// It takes the volumeName as input and returns the size of the volume in gigabytes and an error if any.
// It uses the Docker client to inspect the volume and retrieve its size.
// If the volume does not exist or if there is an error during the process, it returns the error.
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

	if volume.UsageData == nil {
		return size, nil
	}

	return float64(volume.UsageData.Size) / gb, nil
}

// GetNetworkUsage returns the network usage statistics.
// It retrieves the number of bytes received and sent over the network interface.
// It returns a pointer to NetworkUsage and an error if any.
// It uses the net package to fetch the network I/O counters.
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

// GetMemoryUsage returns the memory usage statistics.
// It retrieves the total, free, used memory in gigabytes, and the percentage of memory used.
// It returns a pointer to UsageData and an error if any.
// It uses the mem package to fetch the virtual memory statistics.
func GetMemoryUsage() (*UsageData, error) {
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

	return &UsageData{
		Total:          total,
		Free:           free,
		Used:           used,
		PercentageUsed: percentageUsed,
	}, nil
}

// GetCPUUsage returns the CPU usage statistics.
// It retrieves the total number of CPU cores and the percentage of CPU usage.
// It returns a pointer to UsageData and an error if any.
func GetCPUUsage() (*UsageData, error) {
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

	total := cores * cpuMax
	free := total - int(percent.Percent(pct, total))
	used := total - free

	return &UsageData{
		Total:          total,
		Free:           free,
		Used:           used,
		PercentageUsed: pct,
	}, nil
}

// FormatDate formats a time.Time value to a string using a specific date and time layout string and the specified time zone.
func FormatDate(date time.Time, timeZone string) string {
	loc, err := time.LoadLocation(timeZone)
	if err != nil {
		loc = time.UTC
	}

	localTime := date.In(loc)

	return localTime.Format("01/02/2006 03:04:05 PM")
}

// StringMsToDuration converts a string representing a duration in milliseconds to a time.Duration value.
func StringMsToDuration(durationString string) (time.Duration, error) {
	durationInt, err := strconv.ParseInt(durationString, 0, 64)
	if err != nil {
		return 0, err
	}

	duration := time.Duration(durationInt) * time.Millisecond

	return duration, nil
}

// ParseEpochTime parses a string representing an epoch time and returns a pointer to a time.Time value.
func ParseEpochTime(epoch string) (*time.Time, error) {
	epochInt, err := strconv.ParseInt(epoch, 10, 64)
	if err != nil {
		return nil, err
	}

	epochTime := time.Unix(epochInt/millisecondsInSecond, (epochInt%millisecondsInSecond)*int64(time.Millisecond))

	return &epochTime, nil
}

// DurationToHoursAndMinutes converts a duration in milliseconds to a formatted string representing hours and minutes.
func DurationToHoursAndMinutes(duration time.Duration) string {
	hours := int64(duration.Hours())
	minutes := int64(duration.Minutes()) % secondsInMinute

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

// GetDurationTillTime calculates the duration between a specified start time and the current time and returns it as a time.Duration value.
func GetDurationTillTime(start time.Time, duration time.Duration) (time.Duration, error) {
	endTime := start.Add(duration)
	currentTime := time.Now()

	if endTime.Before(currentTime) {
		return 0, fmt.Errorf("end time is before current time")
	}

	timeLeft := endTime.Sub(currentTime)

	return timeLeft, nil
}

// ParseIntFromString extracts the integer part from a string that may have a suffix.
// The suffix can be any non-digit characters at the end of the string.
// If the string cannot be parsed, an error is returned.
func ParseIntFromString(str string) (int, error) {
	re := regexp.MustCompile(`^\d+`)

	match := re.FindString(str)
	if match == "" {
		return 0, fmt.Errorf("could not parse integer from string: %s", str)
	}

	return strconv.Atoi(match)
}

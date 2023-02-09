package main

import (
	"flag"
	"time"

	"github.com/bartosian/sui_helpers/sui-monitor/cmd/checker"
	"github.com/bartosian/sui_helpers/sui-monitor/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/sui-monitor/pkg/log"

	"github.com/schollz/progressbar/v3"
)

var (
	filePath = flag.String("f", "", "path to node config file")
	network  = flag.String("n", "devnet", "network name")
)

func main() {
	flag.Parse()

	logger := log.NewLogger()

	if *filePath == "" {
		logger.Error("provide path to the config file by using -f option")

		return
	}

	network, err := enums.NetworkTypeFromString(*network)
	if err != nil {
		logger.Error("provide valid network type by using -n option")

		return
	}

	progressCH := make(chan struct{})
	progressTicker := time.NewTicker(100 * time.Millisecond)

	go func() {
		bar := newProgressBar()

		for {
			select {
			case <-progressCH:
				bar.Clear()

				return
			case <-progressTicker.C:
				for i := 0; i < 500; i++ {
					bar.Add(1)

					time.Sleep(5 * time.Millisecond)
				}
			}
		}
	}()

	checker, err := checker.NewChecker(*filePath, network)
	if err != nil {
		logger.Error("failed to create peers checker: ", err)

		return
	}

	if err := checker.ParseData(); err != nil {
		logger.Error("failed to parse data: ", err)

		return
	}

	progressTicker.Stop()
	progressCH <- struct{}{}

	checker.GenerateSystemTable()
	checker.GenerateNodeTable()
	checker.GeneratePeersTable()

	checker.DrawTable()
}

func newProgressBar() *progressbar.ProgressBar {
	return progressbar.NewOptions(1000,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[blue]🔄 [ PROCESSING DATA... ][reset] "),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[blue]=[reset]",
			SaucerHead:    "[blue]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
}

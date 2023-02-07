package main

import (
	"flag"
	"github.com/schollz/progressbar/v3"
	"time"

	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker"
	"github.com/bartosian/sui_helpers/sui-peer-checker/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/sui-peer-checker/pkg/log"
)

var (
	filePath = flag.String("f", "", "path to node config file")
	network  = flag.String("n", "devnet", "network name")
)

func main() {
	flag.Parse()

	logger := log.NewLogger()

	if *filePath == "" {
		logger.Error("provide path to the config file by using -f option", nil)

		return
	}

	network, err := enums.NetworkTypeFromString(*network)
	if err != nil {
		logger.Error("provide valid network type by using -n option", nil)

		return
	}

	progressCH := make(chan struct{})
	progressTicker := time.NewTicker(100 * time.Millisecond)

	go func() {
		bar := newProgressBar()

		for {
			select {
			case <-progressCH:
				bar.Finish()
				bar.Clear()

				return
			case <-progressTicker.C:
				for i := 0; i < 500; i++ {
					bar.Add(1)

					time.Sleep(4 * time.Millisecond)
				}
			}
		}
	}()

	checker, err := checker.NewChecker(*filePath, network)
	if err != nil {
		logger.Error("failed to create peers checker: ", err)

		return
	}

	progressTicker.Stop()
	progressCH <- struct{}{}

	checker.GenerateTableConfig()
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

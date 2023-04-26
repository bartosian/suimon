package progress

import (
	"fmt"
	"os"
	"time"

	"github.com/schollz/progressbar/v3"
)

type Color string

const (
	ColorReset Color = "[reset]"
	ColorWhite Color = "[white]"
	ColorBlue  Color = "[blue]"
	ColorRed   Color = "[red]"
	ColorGreen Color = "[green]"
)

func NewProgressBar(action string, color Color) chan<- struct{} {
	progressTicker := time.NewTicker(100 * time.Millisecond)
	progressChan := make(chan struct{})

	bar := progressbar.NewOptions(1000,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetElapsedTime(false),
		progressbar.OptionShowBytes(false),
		progressbar.OptionClearOnFinish(),
		progressbar.OptionSetWidth(30),
		progressbar.OptionSetDescription(fmt.Sprintf("%s [ %s... ] [reset]", color, action)),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "=",
			SaucerHead:    ">",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	go func() {
		for {
			select {
			case <-progressChan:
				progressTicker.Stop()

				if err := bar.Clear(); err != nil {
					os.Exit(1)
				}

				return
			case <-progressTicker.C:
				for i := 0; i < 100; i++ {
					if err := bar.Add(1); err != nil {
						os.Exit(1)
					}

					time.Sleep(15 * time.Millisecond)
				}
			}
		}
	}()

	return progressChan
}

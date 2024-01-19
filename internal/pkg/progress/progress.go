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

// NewProgressBar creates a new progress bar with the specified action and color.
// It takes the action string and color as input and returns a channel for controlling the progress bar.
// The progress bar is updated at regular intervals and can be stopped by closing the returned channel.
// The color parameter specifies the color of the progress bar.
// Example usage:
//
//	progressChan := NewProgressBar("Downloading", ColorBlue)
//	// Perform download operation
//	close(progressChan) // Stop the progress bar
//
// Note: It is important to close the returned channel to stop the progress bar and free resources.
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
		defer progressTicker.Stop()

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

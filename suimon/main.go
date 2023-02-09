// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"time"

	"github.com/bartosian/sui_helpers/suimon/cmd/checker"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/pkg/log"

	"github.com/schollz/progressbar/v3"
)

var (
	filePath = flag.String("f", "", "(required) path to node config file fullnode.yaml")
	network  = flag.String("n", "devnet", "(optional) network name, possible values: testnet, devnet")
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

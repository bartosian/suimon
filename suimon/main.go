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
	"os"

	checkerBuilder "github.com/bartosian/sui_helpers/suimon/cmd/checker"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/config"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

var (
	suimonConfigPath = flag.String("s", "", "(optional) path to the suimon config file, can use SUIMON_CONFIG_PATH env variable instead")
	nodeConfigPath   = flag.String("f", "", "(optional) path to the node config file, can use SUIMON_NODE_CONFIG_PATH variable instead")
	network          = flag.String("n", "", "(optional) network name, possible values: testnet, devnet")
	watch            = flag.Bool("w", false, "(optional) flag to enable a dynamic dashboard to monitor node metrics in real-time")
)

func main() {
	flag.Parse()

	log.PrintLogo("SUIMON", "banner3", "red")

	var (
		logger        = log.NewLogger()
		checker       *checkerBuilder.Checker
		suimonConfig  *config.SuimonConfig
		nodeConfig    *config.NodeConfig
		networkConfig enums.NetworkType
		err           error
	)

	// parse suimon.yaml config file
	if suimonConfig, err = config.ParseSuimonConfig(suimonConfigPath); err != nil {
		return
	}

	// parse fullnode.yaml or validator.yaml config files
	if nodeConfig, err = config.ParseNodeConfig(nodeConfigPath, suimonConfig.NodeConfigPath); err != nil {
		return
	}

	// parse network config
	if networkConfig, err = config.ParseNetworkConfig(suimonConfig, network); err != nil {
		return
	}

	// create checker instance
	if checker, err = checkerBuilder.NewChecker(*suimonConfig, *nodeConfig, networkConfig); err != nil {
		logger.Error("failed to create suimon instance: ", err)

		return
	}

	// initialize checker instance with seed data
	if err = checker.Init(); err != nil {
		logger.Error("failed to init suimon instance: ", err)

		return
	}

	if *watch {
		// initialize realtime dashboard with styles
		checker.InitDashboard()

		// draw initialized dashboard to the terminal
		checker.DrawDashboards()
	} else {
		// initialize tables with the styles
		checker.InitTables()

		// draw initialized tables to the terminal
		checker.DrawTables()
	}

	defer func() {
		if err := recover(); err != nil {
			checker.DashboardBuilder.Terminal.Close()
			checker.DashboardBuilder.Ctx.Done()

			logger.Error("failed to execute suimon, please check an issue: ", err)

			os.Exit(1)
		}

		return
	}()
}

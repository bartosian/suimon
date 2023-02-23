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

	"github.com/bartosian/sui_helpers/suimon/cmd/checker"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/config"
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/enums"
	"github.com/bartosian/sui_helpers/suimon/pkg/log"
)

var (
	suimonConfigPath = flag.String("sf", "", "(optional) path to the suimon config file, can use SUIMON_CONFIG_PATH env variable instead")
	nodeConfigPath   = flag.String("nf", "", "(optional) path to the node config file, can use SUIMON_NODE_CONFIG_PATH variable instead")
	network          = flag.String("n", "", "(optional) network name, possible values: testnet, devnet")
	watch            = flag.Bool("w", false, "(optional) flag to enable watch mode with dynamic monitoring")
)

func main() {
	flag.Parse()
	
	log.PrintLogo("SUIMON", "banner3", "red")

	var (
		logger        = log.NewLogger()
		check         *checker.Checker
		suimonConfig  *config.SuimonConfig
		nodeConfig    *config.NodeConfig
		networkConfig enums.NetworkType
		err           error
	)

	// parse suimon.yaml config file
	if suimonConfig, err = config.ParseSuimonConfig(suimonConfigPath); err != nil {
		return
	}

	// parse fullnode/validator.yaml config file
	if nodeConfig, err = config.ParseNodeConfig(nodeConfigPath, suimonConfig.NodeConfigPath); err != nil {
		return
	}

	// parse network flag
	if networkConfig, err = config.ParseNetworkConfig(suimonConfig, network); err != nil {
		return
	}

	// create checker instance
	if check, err = checker.NewChecker(*suimonConfig, *nodeConfig, networkConfig); err != nil {
		logger.Error("failed to create suimon instance: ", err)

		return
	}

	// initialize checker instance with seed data
	if err = check.Init(); err != nil {
		logger.Error("failed to init suimon instance: ", err)

		return
	}

	switch *watch {
	case true:
		// initialize realtime dashboard with styles
		check.InitDashboard()

		// draw initialized dashboard to the terminal
		check.DrawDashboards()
	default:
		// initialize tables with the styles
		check.InitTables()

		// draw initialized tables to the terminal
		check.DrawTables()
	}

	defer func() {
		if err := recover(); err != nil {
			check.DashboardBuilder.Terminal.Close()
			check.DashboardBuilder.Ctx.Done()

			logger.Error("failed to execute suimon, please check an issue: ", err)

			os.Exit(1)
		}

		return
	}()
}

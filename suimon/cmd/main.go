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

	"github.com/bartosian/sui_helpers/suimon/cmd/config"
	"github.com/bartosian/sui_helpers/suimon/internal/core/controller"
	"github.com/bartosian/sui_helpers/suimon/internal/pkg/log"
)

var (
	suimonConfigPath = flag.String("s", "", "(optional) path to the suimon config file, can use SUIMON_CONFIG_PATH env variable instead")
	watch            = flag.Bool("w", false, "(optional) flag to enable a dynamic dashboard to monitor node metrics in real-time")
)

func main() {
	flag.Parse()

	var (
		logger            = log.NewLogger()
		checkerController *controller.CheckerController
		suimonConfig      *config.SuimonConfig
		err               error
	)
	// parse suimon.yaml config file
	if suimonConfig, err = config.ParseSuimonConfig(suimonConfigPath); err != nil {
		return
	}

	// create checker instance
	if checkerController, err = controller.NewCheckerController(*suimonConfig); err != nil {
		logger.Error("failed to create suimon instance: ", err)

		return
	}

	defer handlePanic(&logger)

	// initialize checker instance with seed data
	if err = checkerController.ParseData(); err != nil {
		logger.Error("failed to init suimon instance: ", err)

		return
	}

	if *watch {
		// initialize realtime dashboard with styles
		if err = checkerController.InitDashboards(); err != nil {
			logger.Error("failed to init dashboard: ", err)

			return
		}

		// draw initialized dashboard to the terminal
		if err = checkerController.RenderDashboards(); err != nil {
			logger.Error("failed to render dashboard: ", err)

			return
		}

		defer func() {
			checkerController.DashboardBuilder.Terminal.Close()
			checkerController.DashboardBuilder.Ctx.Done()
		}()
	} else {
		// initialize tables with the styles
		if err = checkerController.InitTables(); err != nil {
			logger.Error("failed to init tables: ", err)

			return
		}

		// draw initialized tables to the terminal
		if err = checkerController.RenderTables(); err != nil {
			logger.Error("failed to render tables: ", err)

			return
		}
	}
}

func handlePanic(logger *log.Logger) {
	if r := recover(); r != nil {
		logger.Error("failed to execute suimon, please check an issue: ", r)
		
		os.Exit(1)
	}
}

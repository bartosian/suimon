package controller

import (
	"sync"
	"time"

	"github.com/mum4k/termdash"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
)

// RenderDashboards draws dashboards for each type of table in the Checker struct.
// This function does not accept any parameters.
// This function does not return anything.
func (checker *CheckerController) RenderDashboards() error {
	var (
		monitorsConfig    = checker.suimonConfig.MonitorsConfig
		processLaunchType = checker.suimonConfig.ProcessLaunchType
		dashboardBuilder  = checker.DashboardBuilder
		dashCells         = dashboardBuilder.Cells
		ticker            = time.NewTicker(watchHostsTimeout)
		logsCH            = make(chan string)
		wg                sync.WaitGroup
	)

	defer ticker.Stop()

	var streamLogs = func() {
		defer wg.Done()

		var err error

		if processLaunchType.ServiceName != "" {
			if err = checker.logger.StreamFromService(processLaunchType.ServiceName, logsCH); err == nil {
				return
			}
		}

		if processLaunchType.DockerImageName != "" {
			if err = checker.logger.StreamFromContainer(processLaunchType.DockerImageName, logsCH); err == nil {
				return
			}
		}

		if processLaunchType.ScreenSessionName != "" {
			if err = checker.logger.StreamFromScreen(processLaunchType.ScreenSessionName, logsCH); err == nil {
				return
			}
		}
	}

	var render = func(hosts []host.Host) {
		defer wg.Done()

		doneCH := make(chan struct{}, len(hosts))

		for {
			select {
			case <-ticker.C:
				for _, hostToRender := range hosts {
					go func(host host.Host) {
						for idx, dashCell := range dashCells {
							cellName := enums.CellName(idx)

							metric := checker.getMetricForDashboardCell(cellName)
							options := checker.getOptionsForDashboardCell(cellName)

							dashCell.Write(metric, options)
						}

						doneCH <- struct{}{}
					}(hostToRender)
				}

				for i := 0; i < len(hosts); i++ {
					<-doneCH
				}
			case log := <-logsCH:
				dashCell := dashCells[enums.CellNameNodeLogs]
				options := checker.getOptionsForDashboardCell(enums.CellNameNodeLogs)

				dashCell.Write(log+"\n", options)
			case <-dashboardBuilder.Ctx.Done():
				close(doneCH)
				close(logsCH)

				return
			}
		}
	}

	wg.Add(1)

	go func() {
		defer wg.Done()

		checker.Watch()
	}()

	if monitorsConfig.NodeTable.Display && len(checker.node) > 0 {
		wg.Add(2)

		go streamLogs()
		go render(checker.node)
	}

	if err := termdash.Run(dashboardBuilder.Ctx, dashboardBuilder.Terminal, dashboardBuilder.Dashboard, termdash.KeyboardSubscriber(dashboardBuilder.Quitter)); err != nil {
		panic(err)
	}

	wg.Wait()

	return nil
}

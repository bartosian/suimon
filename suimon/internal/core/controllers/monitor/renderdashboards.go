package monitor

//import (
//	"sync"
//	"time"
//
//	"github.com/mum4k/termdash"
//
//	"github.com/bartosian/sui_helpers/suimon/internal/core/controllers"
//	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
//	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/host"
//)
//
//func (c *Controller) RenderDashboards() error {
//	monitorsConfig := c.config.MonitorsConfig
//	dashboardBuilder := c.builders.dynamic[enums.TableTypeNode]
//
//	ticker := time.NewTicker(controllers.watchHostsTimeout)
//
//	var wg sync.WaitGroup
//
//	defer ticker.Stop()
//
//	var render = func(hosts []host.Host) {
//		defer wg.Done()
//
//		doneCH := make(chan struct{}, len(hosts))
//
//		for {
//			select {
//			case <-ticker.C:
//				for _, hostToRender := range hosts {
//					go func(host host.Host) {
//						for idx, dashCell := range dashCells {
//							cellName := enums.CellName(idx)
//
//							metric := checker.getMetricForDashboardCell(cellName)
//							options := checker.getOptionsForDashboardCell(cellName)
//
//							dashCell.Write(metric, options)
//						}
//
//						doneCH <- struct{}{}
//					}(hostToRender)
//				}
//
//				for i := 0; i < len(hosts); i++ {
//					<-doneCH
//				}
//			case <-dashboardBuilder.Ctx.Done():
//				close(doneCH)
//
//				return
//			}
//		}
//	}
//
//	wg.Add(1)
//
//	go func() {
//		defer wg.Done()
//
//		checker.Watch()
//	}()
//
//	if monitorsConfig.NodeTable.Display && len(checker.hosts.node) > 0 {
//		wg.Add(2)
//
//		go render(checker.hosts.node)
//	}
//
//	if err := termdash.Run(dashboardBuilder.Ctx, dashboardBuilder.Terminal, dashboardBuilder.Dashboard, termdash.KeyboardSubscriber(dashboardBuilder.Quitter)); err != nil {
//		panic(err)
//	}
//
//	wg.Wait()
//
//	return nil
//}

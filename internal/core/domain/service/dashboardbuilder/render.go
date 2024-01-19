package dashboardbuilder

import (
	"fmt"
	"os"
	"time"

	"github.com/mum4k/termdash"
	"golang.org/x/sync/errgroup"

	"github.com/bartosian/suimon/internal/core/domain/service/dashboardbuilder/dashboards"
)

const (
	renderInterval = 200 * time.Millisecond
	queryInterval  = 2500 * time.Millisecond
)

// Render renders the dashboard by starting the query and rerender loops,
// and waiting for them to complete. It returns an error if any of the loops
// encounter an error.
func (db *Builder) Render() (err error) {
	defer func() {
		if r := recover(); r != nil {
			db.cliGateway.Error(fmt.Sprintf("panic: %v", r))
			db.tearDown()
			os.Exit(1)
		} else if err != nil {
			db.tearDown()
		}
	}()

	var errGroup errgroup.Group

	queryTicker, renderTicker := startTickers()
	defer stopTickers(queryTicker, renderTicker)

	errGroup.Go(queryMetricsLoop(db, queryTicker))
	errGroup.Go(rerenderLoop(db, renderTicker))
	errGroup.Go(runDashboard(db))

	return errGroup.Wait()
}

func startTickers() (queryTicker *time.Ticker, renderTicker *time.Ticker) {
	queryTicker = time.NewTicker(queryInterval)
	renderTicker = time.NewTicker(renderInterval)
	return
}

func stopTickers(tickers ...*time.Ticker) {
	for _, ticker := range tickers {
		ticker.Stop()
	}
}

// queryMetricsLoop fetches the metrics from the host at regular intervals.
// It uses the provided ticker to trigger the fetch and returns an error if
// the fetch encounters an error or if the context is done.
// It returns a function that can be used to start the loop.
// The loop can be stopped by canceling the context.
// The function signature is compatible with the errgroup.Group.Go method.
// The loop stops when the context is done.
func queryMetricsLoop(db *Builder, ticker *time.Ticker) func() error {
	return func() error {
		for {
			select {
			case <-ticker.C:
				if err := db.host.GetMetrics(); err != nil {
					return err
				}
			case <-db.ctx.Done():
				return nil
			}
		}
	}
}

// rerenderLoop continuously fetches the latest column values from the host at regular intervals.
// It uses the provided ticker to trigger the fetch and updates the cells with the latest values.
// The loop stops when the context is done.
// The function signature is compatible with the errgroup.Group.Go method.
// It returns a function that can be used to start the loop.
func rerenderLoop(db *Builder, ticker *time.Ticker) func() error {
	return func() error {
		for {
			select {
			case <-ticker.C:
				columnValues, err := dashboards.GetColumnsValues(db.tableType, db.host)
				if err != nil {
					return err
				}

				for columnName, cell := range db.cells {
					columnValue, ok := columnValues[columnName]
					if !ok {
						return fmt.Errorf("failed to get metric for column %s", columnName)
					}

					if err := cell.Write(columnValue); err != nil {
						return err
					}
				}
			case <-db.ctx.Done():
				return nil
			}
		}
	}
}

// runDashboard runs the dashboard using termdash.Run method.
// It takes a Builder instance as input and returns a function that can be used to start the dashboard.
// The returned function can be used to start the dashboard and handle keyboard events using the provided quitter function.
// It returns an error if the dashboard fails to run.
// The dashboard is run using the termdash.Run method with the provided context, terminal, dashboard, and keyboard subscriber.
// The returned error indicates any failure during the dashboard run.
// The function signature is compatible with the errgroup.Group.Go method.
// It returns a function that can be used to start the dashboard.
func runDashboard(db *Builder) func() error {
	return func() error {
		return termdash.Run(db.ctx, db.terminal, db.dashboard, termdash.KeyboardSubscriber(db.quitter))
	}
}

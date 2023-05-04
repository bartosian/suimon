package checker

import (
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/dashboardbuilder"
)

func (checker *Checker) InitDashboard() error {
	var (
		dashboard *dashboardbuilder.DashboardBuilder
		err       error
	)

	if dashboard, err = dashboardbuilder.NewDashboardBuilder(); err != nil {
		return err
	}

	checker.DashboardBuilder = dashboard

	return nil
}

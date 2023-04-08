package controllers

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/dashboardbuilder"
)

// InitDashboards initializes the CheckerController's internal state for the dashboard by creating and setting the appropriate DashboardMetric and DashboardOptions objects for each dashboard cell.
// The function populates the CheckerController's internal state with information about the available hosts and their corresponding rpcgw.
// Returns an error if the initialization process fails for any reason.
func (checker *CheckerController) InitDashboards() error {
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

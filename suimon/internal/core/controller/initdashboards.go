package controller

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/dashboardbuilder"
)

// InitDashboards initializes the dashboard of the "Checker" struct instance passed as a pointer receiver.
// Parameters: None.
// Returns: an error, if any occurred during dashboard initialization.
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

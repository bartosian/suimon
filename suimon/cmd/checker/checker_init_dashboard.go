package checker

import (
	"github.com/bartosian/sui_helpers/suimon/cmd/checker/dashboardbuilder"
)

// InitDashboard initializes the dashboard of the "Checker" struct instance passed as a pointer receiver.
// Parameters: None.
// Returns: an error, if any occurred during dashboard initialization.
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

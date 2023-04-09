package monitor

func (c *Controller) Dynamic() error {
	return nil
}

func (c *Controller) InitDashboards() error {
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

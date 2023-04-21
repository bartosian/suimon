package monitor

import (
	"fmt"

	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/service/dashboardbuilder"
)

// Dynamic is a method of the Controller struct, responsible for initializing and rendering dashboards
// based on the configuration data.
func (c *Controller) Dynamic() error {
	// Parse the configuration data.
	if err := c.ParseConfigData(enums.MonitorTypeDynamic); err != nil {
		return err
	}

	// Initialize dashboard based on the configuration data.
	if err := c.InitDashboard(); err != nil {
		return err
	}

	// Render the dashboard and return error if any
	return c.RenderDashboards()
}

// InitDashboard initializes the enabled dashboard based on the display configuration.
// It retrieves the corresponding hosts for the dashboard and initializes the dashboard builder.
// If an error occurs during table initialization, it returns an error.
func (c *Controller) InitDashboard() error {
	selectedDashboard := c.selectedDashboard

	dashboard, err := dashboardbuilder.NewBuilder(selectedDashboard, c.gateways.cli)
	if err != nil {
		return fmt.Errorf("error creating dashboard %s: %w", selectedDashboard, err)
	}

	return dashboard.Init()
}

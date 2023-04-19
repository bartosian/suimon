package monitor

import (
	"github.com/bartosian/sui_helpers/suimon/internal/core/domain/enums"
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

	return nil
}

func (c *Controller) InitDashboard() error {

	return nil
}

package monitor

import "fmt"

// RenderDashboards displays the selected dashboard on the terminal. First, the function retrieves the name of the
// currently selected dashboard. Then, it retrieves the corresponding dynamic dashboard builder from the controller's
// `builders` map. If the builder is found, the function calls its `Render` method to render the dashboard. If the
// builder is not found, the function returns an error. Finally, the function returns nil to indicate success.
func (c *Controller) RenderDashboards() error {
	selectedDashboard := c.selectedDashboard

	builder := c.builders.dynamic[selectedDashboard]

	if err := builder.Render(); err != nil {
		return fmt.Errorf("error rendering dashboard %s: %w", selectedDashboard, err)
	}

	return nil
}

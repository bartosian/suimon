package ports

type CheckerController interface {
	ParseData() error
	InitTables() error
	RenderTables() error
	InitDashboards() error
	RenderDashboards() error
	Watch() error
}

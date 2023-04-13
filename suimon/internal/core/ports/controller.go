package ports

type RootController interface {
	BeforeStart() bool
}

type VersionController interface {
	PrintVersion()
}

type MonitorController interface {
	Monitor() error
	Static() error
	Dynamic() error
}

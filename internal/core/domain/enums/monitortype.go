package enums

type MonitorType string

const (
	MonitorTypeStatic  MonitorType = "STATIC"
	MonitorTypeDynamic MonitorType = "DYNAMIC"
)

func (e MonitorType) ToString() string {
	return string(e)
}

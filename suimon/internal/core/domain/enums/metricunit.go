package enums

type MetricUnit string

const (
	MetricUnitPercentage MetricUnit = "%"
	MetricUnitB          MetricUnit = "B"
	MetricUnitGB         MetricUnit = "GB"
	MetricUnitTB         MetricUnit = "TB"
)

func (e MetricUnit) ToString() string {
	return string(e)
}

package chart

import (
	testutil "test/e2e/util/test"
)

type Chart interface {
	Meta() Meta
	RequiredChartValues() map[string]string
}

type Meta struct {
	name string
}

func (m Meta) ChartPath() string {
	return testutil.NewPathRelativeToRootDir("test/charts/" + m.name)
}

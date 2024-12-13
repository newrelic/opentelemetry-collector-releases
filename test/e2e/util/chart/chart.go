package chart

import (
	"fmt"
	"os"
	"path"
	"strings"
)

type Chart interface {
	Meta() Meta
	RequiredChartValues() map[string]string
}

type Meta struct {
	name string
}

func (m Meta) ChartPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	testDirPath := path.Clean(fmt.Sprintf("%s/../..", pwd))
	if !strings.HasSuffix(testDirPath, "test") {
		panic(fmt.Sprintf("Unexpected directory structure: %s should be the test dir containing the charts directory (pwd: %s)", testDirPath, pwd))
	}
	return path.Join(testDirPath, "charts", m.name)
}

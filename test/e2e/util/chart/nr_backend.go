package chart

import (
	"fmt"
	envutil "test/e2e/util/env"
)

type NrBackendChart struct {
	CollectorHostname string
}

func NewNrBackendChart(testId string) NrBackendChart {
	environmentName := "local"
	if envutil.IsContinuousIntegration() {
		environmentName = "ci"
	}
	collectorHostname := fmt.Sprintf("nr-otel-collector-%s-%s", environmentName, testId)

	return NrBackendChart{
		CollectorHostname: collectorHostname,
	}
}

func (m *NrBackendChart) AsChart() Chart {
	var chart Chart = m
	return chart
}

func (m *NrBackendChart) Meta() Meta {
	return Meta{
		name: "nr_backend",
	}
}

func (m *NrBackendChart) RequiredChartValues() map[string]string {
	return map[string]string{
		"image.tag":            envutil.GetImageTag(),
		"secrets.nrBackendUrl": envutil.GetNrBackendUrl(),
		"secrets.nrIngestKey":  envutil.GetNrIngestKey(),
		"collector.hostname":   m.CollectorHostname,
	}
}

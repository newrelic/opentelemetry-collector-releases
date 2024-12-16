package chart

import envutil "test/e2e/util/env"

type NrBackendChart struct {
}

var NrBackend = NrBackendChart{}

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
	}
}

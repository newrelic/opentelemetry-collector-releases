package chart

import utilenv "test/e2e/util/env"

type MockedBackendChart struct {
}

var MockedBackend = MockedBackendChart{}

func (m *MockedBackendChart) AsChart() Chart {
	var chart Chart = m
	return chart
}

func (m *MockedBackendChart) Meta() Meta {
	return Meta{
		name: "mocked_backend",
	}
}

func (m *MockedBackendChart) RequiredChartValues() map[string]string {
	return map[string]string{
		"image.tag": utilenv.GetImageTag(),
	}
}

func (m *MockedBackendChart) GetCollectorSelector() string {
	//TODO implement me
	panic("implement me")
}

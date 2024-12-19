package spec

import "test/e2e/util/assert"

type TestCase struct {
	Metric     assert.Metric
	Assertions []assert.Assertion
}

// TODO: eventually this should be generated based on a 'test specification' file
var testCases = []TestCase{
	{
		Metric: assert.Metric{
			Name:        "system.cpu.utilization",
			WhereClause: "WHERE state='user'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Metric: assert.Metric{
			Name:        "system.disk.io",
			WhereClause: "WHERE direction='read'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Metric: assert.Metric{
			Name:        "system.disk.io",
			WhereClause: "WHERE direction='write'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
}

func GetOnHostTestCases() []TestCase {
	return testCases
}

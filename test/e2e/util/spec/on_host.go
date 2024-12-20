package spec

import "test/e2e/util/assert"

type TestCase struct {
	Name       string
	Metric     assert.Metric
	Assertions []assert.Assertion
}

// TODO: eventually this should be generated based on a 'test specification' file
var testCases = []TestCase{
	{
		Name: "host receiver cpu.utilization",
		Metric: assert.Metric{
			Name:        "system.cpu.utilization",
			WhereClause: "WHERE state='user'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver disk.io read",
		Metric: assert.Metric{
			Name:        "system.disk.io",
			WhereClause: "WHERE direction='read'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">=", Threshold: 0},
		}},
	{
		Name: "host receiver disk.io read",
		Metric: assert.Metric{
			Name:        "system.disk.io",
			WhereClause: "WHERE direction='write'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver disk.io_time",
		Metric: assert.Metric{
			Name: "system.disk.io_time",
		},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver disk.operation_time read",
		Metric: assert.Metric{
			Name:        "system.disk.operation_time",
			WhereClause: "WHERE direction='read'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">=", Threshold: 0},
		}},
	{
		Name: "host receiver disk.operation_time read",
		Metric: assert.Metric{
			Name:        "system.disk.operation_time",
			WhereClause: "WHERE direction='write'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver disk.operations read",
		Metric: assert.Metric{
			Name:        "system.disk.operations",
			WhereClause: "WHERE direction='read'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">=", Threshold: 0},
		}},
	{
		Name: "host receiver disk.operations write",
		Metric: assert.Metric{
			Name:        "system.disk.operations",
			WhereClause: "WHERE direction='write'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver system.load 1m",
		Metric: assert.Metric{
			Name: "system.cpu.load_average.1m",
		},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver system.load 5m",
		Metric: assert.Metric{
			Name: "system.cpu.load_average.5m",
		},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver system.load 15m",
		Metric: assert.Metric{
			Name: "system.cpu.load_average.15m",
		},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver memory.usage cached",
		Metric: assert.Metric{
			Name:        "system.memory.usage",
			WhereClause: "WHERE state='cached'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver memory.usage free",
		Metric: assert.Metric{
			Name:        "system.memory.usage",
			WhereClause: "WHERE state='free'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver memory.usage slab_reclaimable",
		Metric: assert.Metric{
			Name:        "system.memory.usage",
			WhereClause: "WHERE state='slab_reclaimable'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver memory.usage buffered",
		Metric: assert.Metric{
			Name:        "system.memory.usage",
			WhereClause: "WHERE state='buffered'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver memory.usage used",
		Metric: assert.Metric{
			Name:        "system.memory.usage",
			WhereClause: "WHERE state='used'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver memory.utilization free",
		Metric: assert.Metric{
			Name:        "system.memory.utilization",
			WhereClause: "WHERE state='free'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver memory.utilization used",
		Metric: assert.Metric{
			Name:        "system.memory.utilization",
			WhereClause: "WHERE state='used'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver system.paging.operations out",
		Metric: assert.Metric{
			Name:        "system.paging.operations",
			WhereClause: "WHERE direction='page_out'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">=", Threshold: 0},
		}},
	{
		Name: "host receiver system.paging.operations in",
		Metric: assert.Metric{
			Name:        "system.paging.operations",
			WhereClause: "WHERE direction='page_in'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">=", Threshold: 0},
		}},
	{
		Name: "host receiver paging.usage used",
		Metric: assert.Metric{
			Name:        "system.paging.usage",
			WhereClause: "WHERE state='used'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver paging.usage free",
		Metric: assert.Metric{
			Name:        "system.paging.usage",
			WhereClause: "WHERE state='free'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver inodes.usage free",
		Metric: assert.Metric{
			Name:        "system.filesystem.inodes.usage",
			WhereClause: "WHERE state='free'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver inodes.usage used",
		Metric: assert.Metric{
			Name:        "system.filesystem.inodes.usage",
			WhereClause: "WHERE state='used'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver filesystem.usage used",
		Metric: assert.Metric{
			Name:        "system.filesystem.usage",
			WhereClause: "WHERE state='used'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver filesystem.usage free",
		Metric: assert.Metric{
			Name:        "system.filesystem.usage",
			WhereClause: "WHERE state='free'"},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver filesystem.utilization",
		Metric: assert.Metric{
			Name: "system.filesystem.utilization",
		},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver network dropped receive",
		Metric: assert.Metric{
			Name:        "system.network.dropped",
			WhereClause: "WHERE direction='receive'",
		},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">=", Threshold: 0},
		}},
	{
		Name: "host receiver network dropped transmit",
		Metric: assert.Metric{
			Name:        "system.network.dropped",
			WhereClause: "WHERE direction='transmit'",
		},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">=", Threshold: 0},
		}},
	{
		Name: "host receiver network errors receive",
		Metric: assert.Metric{
			Name:        "system.network.errors",
			WhereClause: "WHERE direction='receive'",
		},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">=", Threshold: 0},
		}},
	{
		Name: "host receiver network errors transmit",
		Metric: assert.Metric{
			Name:        "system.network.errors",
			WhereClause: "WHERE direction='transmit'",
		},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">=", Threshold: 0},
		}},
	{
		Name: "host receiver network io receive",
		Metric: assert.Metric{
			Name:        "system.network.io",
			WhereClause: "WHERE direction='receive'",
		},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver network io transmit",
		Metric: assert.Metric{
			Name:        "system.network.io",
			WhereClause: "WHERE direction='transmit'",
		},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver network packets receive",
		Metric: assert.Metric{
			Name:        "system.network.packets",
			WhereClause: "WHERE direction='receive'",
		},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
	{
		Name: "host receiver network packets transmit",
		Metric: assert.Metric{
			Name:        "system.network.packets",
			WhereClause: "WHERE direction='transmit'",
		},
		Assertions: []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		}},
}

func GetOnHostTestCases() []TestCase {
	return testCases
}

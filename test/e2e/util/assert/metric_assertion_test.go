package assert

import (
	"fmt"
	"strings"
	"testing"
)

func TestAsQueryWithSingleAssertion(t *testing.T) {
	assertionFactory := NewMetricAssertionFactory(
		fmt.Sprintf("WHERE host.name = 'nr-otel-collector-foobar'"),
		"5 minutes ago",
	)
	singleAssertion := assertionFactory.NewMetricAssertion("system.cpu.utilization", []Assertion{
		{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
	})
	actual := singleAssertion.AsQuery()
	assertEqual(actual, `
SELECT max(system.cpu.utilization)
FROM Metric
WHERE metricName = 'system.cpu.utilization'
WHERE host.name = 'nr-otel-collector-foobar'
SINCE 5 minutes ago UNTIL now
`, t)
}

func TestAsQueryWithMultipleAssertions(t *testing.T) {
	assertionFactory := NewMetricAssertionFactory(
		fmt.Sprintf("WHERE host.name = 'nr-otel-collector-foobar'"),
		"5 minutes ago",
	)
	singleAssertion := assertionFactory.NewMetricAssertion("system.cpu.utilization", []Assertion{
		{AggregationFunction: "max", ComparisonOperator: "<", Threshold: 0},
		{AggregationFunction: "min", ComparisonOperator: ">", Threshold: 0},
		{AggregationFunction: "average", ComparisonOperator: ">", Threshold: 0},
	})
	actual := singleAssertion.AsQuery()
	assertEqual(actual, `
SELECT max(system.cpu.utilization),min(system.cpu.utilization),average(system.cpu.utilization)
FROM Metric
WHERE metricName = 'system.cpu.utilization'
WHERE host.name = 'nr-otel-collector-foobar'
SINCE 5 minutes ago UNTIL now
`, t)
}

func assertEqual(actual string, expected string, t *testing.T) {
	actualTrimmed := strings.TrimSpace(actual)
	expectedTrimmed := strings.TrimSpace(expected)
	if actualTrimmed != expectedTrimmed {
		t.Fatalf("\nExpected:\n[%s]\nbut received:\n[%s]\n", expectedTrimmed, actualTrimmed)
	}
}

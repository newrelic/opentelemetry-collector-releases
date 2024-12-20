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
	singleAssertion := assertionFactory.NewMetricAssertion(
		Metric{Name: "system.cpu.utilization", WhereClause: "WHERE state='user'"}, []Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		})
	actual := singleAssertion.AsQuery()
	assertEqual(actual, `
SELECT max(^system.cpu.utilization^)
FROM Metric
WHERE state='user'
WHERE host.name = 'nr-otel-collector-foobar'
SINCE 5 minutes ago UNTIL now
`, t)
}

func TestAsQueryWithMultipleAssertions(t *testing.T) {
	assertionFactory := NewMetricAssertionFactory(
		fmt.Sprintf("WHERE host.name = 'nr-otel-collector-foobar'"),
		"5 minutes ago",
	)
	singleAssertion := assertionFactory.NewMetricAssertion(Metric{Name: "system.cpu.utilization", WhereClause: "WHERE state='user'"}, []Assertion{
		{AggregationFunction: "max", ComparisonOperator: "<", Threshold: 0},
		{AggregationFunction: "min", ComparisonOperator: ">", Threshold: 0},
		{AggregationFunction: "average", ComparisonOperator: ">", Threshold: 0},
	})
	actual := singleAssertion.AsQuery()
	assertEqual(actual, `
SELECT max(^system.cpu.utilization^),min(^system.cpu.utilization^),average(^system.cpu.utilization^)
FROM Metric
WHERE state='user'
WHERE host.name = 'nr-otel-collector-foobar'
SINCE 5 minutes ago UNTIL now
`, t)
}

func assertEqual(actual string, expected string, t *testing.T) {
	actualTrimmed := strings.TrimSpace(actual)
	// no way to escape backticks, so we use '^' as a placeholder
	expectedTrimmed := strings.Replace(strings.TrimSpace(expected), "^", "`", -1)
	if actualTrimmed != expectedTrimmed {
		t.Fatalf("\nExpected:\n[%s]\nbut received:\n[%s]\n", expectedTrimmed, actualTrimmed)
	}
}

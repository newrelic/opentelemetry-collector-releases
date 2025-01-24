package hostmetrics

import (
	"fmt"
	"test/e2e/util/assert"
	"test/e2e/util/nr"
	"test/e2e/util/spec"
	testutil "test/e2e/util/test"
	"testing"
	"time"
)

type systemUnderTest struct {
	hostNamePattern string
	excludedMetrics []string
}

var ec2Ubuntu22 = systemUnderTest{
	hostNamePattern: testutil.NewNrQueryHostNamePattern("nightly", testutil.Wildcard, "ec2_ubuntu22_04"),
	// TODO: NR-362121
	excludedMetrics: []string{"system.paging.usage"},
}
var ec2Ubuntu24 = systemUnderTest{
	hostNamePattern: testutil.NewNrQueryHostNamePattern("nightly", testutil.Wildcard, "ec2_ubuntu24_04"),
	// TODO: NR-362121
	excludedMetrics: []string{"system.paging.usage"},
}
var k8sNode = systemUnderTest{
	hostNamePattern: testutil.NewNrQueryHostNamePattern("nightly", testutil.Wildcard, "k8s_node"),
	excludedMetrics: []string{"system.paging.usage"},
}

func TestNightlyHostMetrics(t *testing.T) {
	testutil.TagAsNightlyTest(t)

	// space out requests to not run into 25 concurrent request limit
	requestsPerSecond := 4.0
	requestSpacing := time.Duration((1/requestsPerSecond)*1000) * time.Millisecond
	client := nr.NewClient()

	for _, sut := range []systemUnderTest{ec2Ubuntu22, ec2Ubuntu24, k8sNode} {
		for i, testCase := range spec.GetOnHostTestCasesWithout(sut.excludedMetrics) {
			t.Run(fmt.Sprintf("%s/%d/%s", sut.hostNamePattern, i, testCase.Name), func(t *testing.T) {
				t.Parallel()
				assertionFactory := assert.NewNrMetricAssertionFactory(
					fmt.Sprintf("WHERE host.name like '%s'", sut.hostNamePattern),
					"2 hour ago",
				)
				assertion := assertionFactory.NewNrMetricAssertion(testCase.Metric, testCase.Assertions)
				// space out requests to avoid rate limiting
				time.Sleep(time.Duration(i) * requestSpacing)
				assertion.Execute(t, client)
			})
		}
	}
}

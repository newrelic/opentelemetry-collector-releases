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

var ec2HostNamePattern = testutil.NewNrQueryHostNamePattern("nightly", testutil.Wildcard, "ec2_ubuntu24_04")
var k8sNodeHostNamePattern = testutil.NewNrQueryHostNamePattern("nightly", testutil.Wildcard, "k8s_node")

func TestNightlyBuildCollectors(t *testing.T) {
	testutil.TagAsNightlyTest(t)

	requestsPerSecond := 4.0
	requestSpacing := time.Duration((1/requestsPerSecond)*1000) * time.Millisecond
	for _, hostNamePattern := range []string{ec2HostNamePattern, k8sNodeHostNamePattern} {

		for i, testCase := range spec.GetOnHostTestCases() {
			t.Run(fmt.Sprintf(testCase.Name), func(t *testing.T) {
				t.Parallel()
				assertionFactory := assert.NewNrMetricAssertionFactory(
					fmt.Sprintf("WHERE host.name like '%s'", hostNamePattern),
					"5 minutes ago",
				)
				client := nr.NewClient()
				assertion := assertionFactory.NewNrMetricAssertion(testCase.Metric, testCase.Assertions)
				// space out requests to avoid rate limiting
				time.Sleep(time.Duration(i) * requestSpacing)
				assertion.Execute(t, client)
			})
		}
	}
}

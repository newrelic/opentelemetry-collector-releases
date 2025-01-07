package hostmetrics

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/k8s"
	corev1 "k8s.io/api/core/v1"
	"log"
	"test/e2e/util/assert"
	"test/e2e/util/chart"
	envutil "test/e2e/util/env"
	helmutil "test/e2e/util/helm"
	k8sutil "test/e2e/util/k8s"
	"test/e2e/util/nr"
	"test/e2e/util/spec"
	testutil "test/e2e/util/test"
	"testing"
	"time"
)

const (
	TestNamespace = "nr-hostmetrics"
)

var (
	kubectlOptions            *k8s.KubectlOptions
	testChart                 chart.NrBackendChart
	testId                    string
	collectorReportedHostname string
)

type testEnv struct {
	teardown     func(tb testing.TB)
	collectorPod corev1.Pod
}

func setupTest(tb testing.TB) testEnv {
	collectorPod := k8sutil.WaitForCollectorReady(tb, kubectlOptions)

	return testEnv{collectorPod: collectorPod, teardown: func(tb testing.TB) {
		log.Println("teardown test")
	}}

}

func TestMain(m *testing.M) {
	testId = testutil.NewTestId()
	kubectlOptions = k8sutil.NewKubectlOptions(TestNamespace)
	environmentName := "local"
	if envutil.IsContinuousIntegration() {
		environmentName = "ci"
	}
	collectorReportedHostname = fmt.Sprintf("nr-otel-collector-%s-%s", environmentName, testId)
	testChart = chart.NewNrBackendChart(collectorReportedHostname)
	m.Run()
}

func TestStartupBehavior(t *testing.T) {
	testutil.TagAsSlowTest(t)

	t.Logf("host.name used for test: %s", collectorReportedHostname)
	cleanup := helmutil.ApplyChart(t, kubectlOptions, testChart.AsChart(), "hostmetrics-startup", testId)
	t.Cleanup(cleanup)
	te := setupTest(t)
	t.Cleanup(func() {
		te.teardown(t)
	})
	// wait for at least one default metric harvest cycle (60s) and some buffer to allow NR ingest to process data
	time.Sleep(70 * time.Second)
	// space out requests to not run into 25 concurrent request limit
	requestsPerSecond := 4.0
	requestSpacing := time.Duration((1/requestsPerSecond)*1000) * time.Millisecond

	for i, testCase := range spec.GetOnHostTestCases() {
		t.Run(fmt.Sprintf(testCase.Name), func(t *testing.T) {
			t.Parallel()
			assertionFactory := assert.NewNrMetricAssertionFactory(
				fmt.Sprintf("WHERE host.name = '%s'", collectorReportedHostname),
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

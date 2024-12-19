package hostmetrics

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/k8s"
	corev1 "k8s.io/api/core/v1"
	"log"
	"test/e2e/util/assert"
	"test/e2e/util/chart"
	helmutil "test/e2e/util/helm"
	k8sutil "test/e2e/util/k8s"
	"test/e2e/util/nr"
	testutil "test/e2e/util/test"
	"testing"
	"time"
)

const (
	TestNamespace = "nr-hostmetrics"
)

var (
	kubectlOptions *k8s.KubectlOptions
	testChart      chart.NrBackendChart
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
	kubectlOptions = k8sutil.NewKubectlOptions(TestNamespace)
	testChart = chart.NrBackend
	m.Run()
}

func TestStartupBehavior(t *testing.T) {
	testutil.TagAsSlowTest(t)

	cleanup := helmutil.ApplyChart(t, kubectlOptions, testChart.AsChart(), "hostmetrics-startup")
	defer cleanup()

	t.Run("NRQL validation", func(t *testing.T) {
		te := setupTest(t)
		defer te.teardown(t)
		time.Sleep(120 * time.Second)

		client := nr.NewClient()
		assertionFactory := assert.NewMetricAssertionFactory(
			fmt.Sprintf("WHERE host.name = '%s'", te.collectorPod.Name),
			"5 minutes ago",
		)
		assertionCpuUtil := assertionFactory.NewMetricAssertion("system.cpu.utilization", []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		})
		assertionDiskIo := assertionFactory.NewMetricAssertion("system.disk.io", []assert.Assertion{
			{AggregationFunction: "max", ComparisonOperator: ">", Threshold: 0},
		})
		for _, assertion := range []assert.MetricAssertion{assertionCpuUtil, assertionDiskIo} {
			assertion.Execute(t, client)
		}
	})
}

package hostmetrics

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/newrelic/newrelic-client-go/v2/pkg/nrdb"
	corev1 "k8s.io/api/core/v1"
	"log"
	"test/e2e/util/chart"
	envutil "test/e2e/util/env"
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
		time.Sleep(5 * time.Second)

		client := nr.NewClient()
		query := nrdb.NRQL(fmt.Sprintf(`
SELECT count(*)
FROM Metric
WHERE host.name = '%s'
WHERE metricName = 'system.cpu.utilization'
SINCE 5 minutes ago
`, te.collectorPod.Name))
		result, err := client.Nrdb.Query(envutil.GetNrAccountId(), query)
		if err != nil {
			t.Fatal(err)
		}
		if len(result.Results) == 1 && result.Results[0]["count"].(float64) > 0 {
			t.Logf("count(*) of system.cpu.utilization: %d", result.Results[0]["count"].(int))
		}
	})
}

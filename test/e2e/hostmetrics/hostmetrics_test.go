package hostmetrics

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	httphelper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	corev1 "k8s.io/api/core/v1"
	"log"
	"strings"
	"test/e2e/util/chart"
	helmutil "test/e2e/util/helm"
	k8sutil "test/e2e/util/k8s"
	testutil "test/e2e/util/test"
	"testing"
	"time"
)

const (
	TestNamespace = "mock-hostmetrics"
)

var (
	kubectlOptions *k8s.KubectlOptions
	testChart      chart.MockedBackendChart
	testId         string
)

// TODO: Export from mocked module
type ValidationPayload struct {
	DurationInMillis uint64 `json:"duration"`
	Transactions     uint32 `json:"transactions"`
}

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

func TestStartupBehavior(t *testing.T) {
	testutil.TagAsFastTest(t)
	kubectlOptions = k8sutil.NewKubectlOptions(TestNamespace)
	testChart = chart.MockedBackend
	testId = testutil.NewTestId()
	cleanup := helmutil.ApplyChart(t, kubectlOptions, testChart.AsChart(), "hostmetrics-startup", testId)
	defer cleanup()

	t.Run("healthcheck succeeds", func(t *testing.T) {
		te := setupTest(t)
		defer te.teardown(t)

		tunnel := k8s.NewTunnel(kubectlOptions, k8s.ResourceTypePod, te.collectorPod.Name, 13133, 13133)
		defer tunnel.Close()
		tunnel.ForwardPort(t)

		url := fmt.Sprintf("http://%s/", tunnel.Endpoint())

		httphelper.HttpGetWithRetryWithCustomValidation(t, url, &tls.Config{},
			10, 5*time.Second, func(status int, body string) bool {
				return status == 200 && strings.Contains(body, "Server available")
			})
	})

	t.Run("validation-backend logs indicate processed metrics", func(t *testing.T) {
		te := setupTest(t)
		defer te.teardown(t)

		tunnel := k8s.NewTunnel(kubectlOptions, k8s.ResourceTypeService, "validation-backend", 8080, 8080)
		defer tunnel.Close()
		tunnel.ForwardPort(t)
		url := fmt.Sprintf("http://%s/validate", tunnel.Endpoint())

		httphelper.HttpGetWithRetryWithCustomValidation(t, url, nil, 3, 3*time.Second, func(statusCode int, body string) bool {

			if statusCode != 200 {
				return false
			}

			var payload ValidationPayload
			err := json.NewDecoder(strings.NewReader(body)).Decode(&payload)

			if err != nil {
				fmt.Println(err)
				return false
			}

			return payload.Transactions >= 1
		})
	})
}

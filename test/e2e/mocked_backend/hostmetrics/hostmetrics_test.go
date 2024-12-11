package hostmetrics

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/gruntwork-io/terratest/modules/helm"
	httphelper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

const (
	TestNamespace = "hostmetrics"
)

var (
	contextName    string
	collectorRepo  string
	collectorTag   string
	kubectlOptions *k8s.KubectlOptions
	helmOptions    *helm.Options
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
	filters := metav1.ListOptions{
		LabelSelector: "app=nr-otel-collector",
	}
	k8s.WaitUntilNumPodsCreated(tb, kubectlOptions, filters, 1, 30, 10*time.Second)

	pods := k8s.ListPods(tb, kubectlOptions, filters)
	for _, pod := range pods {
		k8s.WaitUntilPodAvailable(tb, kubectlOptions, pod.Name, 30, 10*time.Second)
	}

	return testEnv{collectorPod: pods[0], teardown: func(tb testing.TB) {
		log.Println("teardown test")
	}}

}

func TestMain(m *testing.M) {
	contextName = os.Getenv("E2E_TEST__K8S_CONTEXT_NAME")
	if contextName == "" {
		panic("E2E_TEST__K8S_CONTEXT_NAME not set: provide existing kubeconfig context")
	}
	kubectlOptions = k8s.NewKubectlOptions(contextName, "", TestNamespace)

	collectorRepo = os.Getenv("E2E_TEST__IMAGE_REPO")
	if collectorRepo == "" {
		panic("E2E_TEST__IMAGE_REPO not set: provide image repo to test, e.g. newrelic/nr-otel-collector")
	}
	collectorTag = os.Getenv("E2E_TEST__IMAGE_TAG")
	if collectorTag == "" {
		panic("E2E_TEST__IMAGE_TAG not set: provide image to test which was previously loaded into local registry")
	}

	helmOptions = &helm.Options{
		KubectlOptions: kubectlOptions,
		ExtraArgs: map[string][]string{
			"install": {
				"--namespace", TestNamespace, "--create-namespace",
				"--set", fmt.Sprintf("image.tag=%s", collectorTag)},
		},
	}

	m.Run()
}

func TestStartupBehavior(t *testing.T) {
	releaseName := fmt.Sprintf(
		"hostmetrics-startup-%s", strings.ToLower(random.UniqueId()))
	defer helm.Delete(t, helmOptions, releaseName, true)
	helm.Install(t, helmOptions, "../../../charts/mocked_backend", releaseName)

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

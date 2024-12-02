package hostmetrics

import (
	"crypto/tls"
	"fmt"
	"github.com/gruntwork-io/terratest/modules/helm"
	http_helper "github.com/gruntwork-io/terratest/modules/http-helper"
	"github.com/gruntwork-io/terratest/modules/k8s"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"regexp"
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
			"upgrade": {
				"-i",
				"--namespace", TestNamespace, "--create-namespace",
				"--set", fmt.Sprintf("image.tag=%s", collectorTag)},
		},
	}
	m.Run()
}

func TestStartupBehavior(t *testing.T) {
	chartReleaseName := "harvest-hostmetrics-and-publish-to-fake-backend"
	helm.Upgrade(t, helmOptions, ".", chartReleaseName)
	defer helm.Delete(t, helmOptions, chartReleaseName, true)
	pods := k8s.ListPods(t, kubectlOptions, metav1.ListOptions{})
	for _, pod := range pods {
		k8s.WaitUntilPodAvailable(t, kubectlOptions, pod.Name, 10, 10*time.Second)
	}
	t.Run("healthcheck succeeds", func(t *testing.T) {
		http_helper.HttpGetWithRetryWithCustomValidation(t, "http://localhost:30132/", &tls.Config{},
			10, 5*time.Second, func(status int, body string) bool {
				return status == 200 && strings.Contains(body, "Server available")
			})
	})
	t.Run("validation-backend logs indicate processed metrics", func(t *testing.T) {
		var verificationPod *corev1.Pod
		for _, pod := range pods {
			if strings.HasPrefix(pod.Name, "validation-backend") {
				verificationPod = &pod
			}
		}
		var logs string
		for i := 0; i < 5; i++ {
			logs = k8s.GetPodLogs(t, kubectlOptions, verificationPod, "validation-backend")
			pattern := `Metrics\s*\{"kind":\s*"exporter", "data_type":\s*"metrics",\s*"name":\s*"debug",\s*"resource metrics":\s*\d+,\s*"metrics":\s*\d+,\s*"data points":\s*\d+\}`
			matched, err := regexp.MatchString(pattern, logs)
			if err != nil {
				t.Fatal(err)
			}
			if matched {
				return
			} else {
				time.Sleep(1 * time.Second)
				continue
			}
		}
		t.Fatalf("validation-backend logs do not indicate processed metrics:\n=\n=\n=\n%s", logs)
	})
}

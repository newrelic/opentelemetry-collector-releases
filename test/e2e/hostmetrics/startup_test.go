package hostmetrics

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
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

func TestCollectorStartup(t *testing.T) {

	chartReleaseName := "collector-startup"
	helm.Upgrade(t, helmOptions, ".", chartReleaseName)
	defer helm.Delete(t, helmOptions, chartReleaseName, true)

	pods := k8s.ListPods(t, kubectlOptions, v1.ListOptions{
		LabelSelector: "app=nr-otel-collector",
	})
	fmt.Println(pods)
	if len(pods) != 1 {
		t.Fatalf("Expected 1 pod but got %d", len(pods))
	}
	k8s.WaitUntilPodAvailable(t, kubectlOptions, pods[0].Name, 10, 10*time.Second)
	stdout, err := k8s.RunKubectlAndGetOutputE(t, kubectlOptions, "get", "pods")
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(100 * time.Minute)
	fmt.Println(stdout)
	// TODO: verify collector working via healthcheck and daemonset logs maybe?
}

package k8s

import (
	"github.com/gruntwork-io/terratest/modules/k8s"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"strings"
	envutil "test/e2e/util/env"
	"testing"
	"time"
)

func NewKubectlOptions(namespace string) *k8s.KubectlOptions {
	contextName := envutil.GetK8sContextName()
	return k8s.NewKubectlOptions(contextName, "", namespace)
}

func WaitForCollectorReady(tb testing.TB, kubectlOptions *k8s.KubectlOptions) corev1.Pod {
	filters := metav1.ListOptions{
		LabelSelector: "app=nr-otel-collector",
	}
	k8s.WaitUntilNumPodsCreated(tb, kubectlOptions, filters, 1, 30, 10*time.Second)

	pods := k8s.ListPods(tb, kubectlOptions, filters)
	for _, pod := range pods {
		k8s.WaitUntilPodAvailable(tb, kubectlOptions, pod.Name, 30, 10*time.Second)
	}
	return pods[0]
}

func GetReportedHostname(tb testing.TB, kubectlOptions *k8s.KubectlOptions, collectorPodName string) string {
	// CI runs in VMs which appears to force hostname to be the VM hostname independent of k8s configs
	if envutil.IsContinuousIntegration() {
		hostname, err := os.Hostname()
		if err != nil {
			tb.Fatal("Couldn't read hostname in CI environment", err)
		}
		tb.Logf("using hostname of test runner (CI environment): %s", hostname)
		return hostname
	}
	collectorNodeHostname, err := k8s.RunKubectlAndGetOutputE(tb, kubectlOptions, "exec", collectorPodName,
		"-c", "nr-otel-collector", "--", "hostname")
	if err != nil {
		tb.Fatal("Couldn't determine hostname of collector pod", err)
	}
	collectorNodeHostname = strings.TrimSpace(collectorNodeHostname)
	tb.Logf("using hostname of collector node: %s", collectorNodeHostname)
	return collectorNodeHostname
}

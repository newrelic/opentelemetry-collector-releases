package k8s

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/k8s"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	envutil "test/e2e/util/env"
	"testing"
	"time"
)

func NewKubectlOptions(namespacePrefix string) *k8s.KubectlOptions {
	namespace := newTestNamespace(namespacePrefix, envutil.GetDistro())
	contextName := envutil.GetK8sContextName()
	return k8s.NewKubectlOptions(contextName, "", namespace)
}

func newTestNamespace(namespacePrefix string, distro string) string {
	return fmt.Sprintf("%s-%s", namespacePrefix, distro)
}

func WaitForCollectorReady(tb testing.TB, kubectlOptions *k8s.KubectlOptions) corev1.Pod {
	filters := metav1.ListOptions{
		LabelSelector: "app=nrdot-collector",
	}
	k8s.WaitUntilNumPodsCreated(tb, kubectlOptions, filters, 1, 30, 10*time.Second)

	pods := k8s.ListPods(tb, kubectlOptions, filters)
	for _, pod := range pods {
		k8s.WaitUntilPodAvailable(tb, kubectlOptions, pod.Name, 30, 10*time.Second)
	}
	return pods[0]
}

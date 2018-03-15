// +build k8srequired

package draining

import (
	"fmt"
	"testing"

	"github.com/giantswarm/e2e-harness/pkg/harness"
	"github.com/giantswarm/microerror"
	"github.com/go-resty/resty"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/giantswarm/aws-operator/integration/template"
)

const (
	e2eAppValuesFilePath = "/tmp/e2e-app-values.yaml"
)

func Test_Integration_Draining(t *testing.T) {
	var err error

	// Creating guest cluster kube config.
	tmpFileName, currentKubeConfig, err := SetupGuestKubeConfig()
	if err != nil {
		t.Error("expected", nil, "got", err)
	}
	defer func() {
		err = TeardownGuestKubeConfig(tmpFileName, currentKubeConfig)
		if err != nil {
			t.Error("expected", nil, "got", err)
		}
	}()

	// Creating guest cluster Kubernetes client.
	var guestK8sClient kubernetes.Interface
	{
		c, err := clientcmd.BuildConfigFromFlags("", harness.DefaultKubeConfig)
		if err != nil {
			t.Error("expected", nil, "got", err)
		}
		guestK8sClient, err = kubernetes.NewForConfig(c)
		if err != nil {
			t.Error("expected", nil, "got", err)
		}
	}

	// Define a custom wait function to ensure the e2e-app is running before we
	// start our integration test.
	customWaitFor := func() error {
		i := metav1.ListOptions{LabelSelector: "app=2e2-app"}
		pods, err := guestK8sClient.CoreV1().Pods(namespace).List(i)
		if err != nil {
			return microerror.Mask(err)
		}

		if len(pods.Items) != 2 {
			return microerror.Maskf(podNotRunningError)
		}

		for _, p := range pods.Items {
			if pod.Status.Phase != v1.PodRunning {
				return microerror.Maskf(podNotRunningError)
			}
		}

		return nil
	}

	// Start the 2e2-app and wait for it to be up.
	err = f.InstallGuestChart(e2eAppValuesFilePath, template.E2EAppChartValues, customWaitFor)
	if err != nil {
		t.Error("expected", nil, "got", err)
	}

	fmt.Printf("All two pods of the 2e2-app are running.\n")

	resp, err := resty.R().Get("http://e2e-app:8000/")

	fmt.Printf("\nError: %v", err)
	fmt.Printf("\nResponse Status Code: %v", resp.StatusCode())
	fmt.Printf("\nResponse Status: %v", resp.Status())
	fmt.Printf("\nResponse Body: %v", resp)
	fmt.Printf("\nResponse Time: %v", resp.Time())
	fmt.Printf("\nResponse Recevied At: %v", resp.ReceivedAt())

	// TODO run async requests continuously
	// TODO scale down
	// TODO wait for scaled down
	// TODO check if requests were interrupted
	// TODO fail test if requests were interrupted
}

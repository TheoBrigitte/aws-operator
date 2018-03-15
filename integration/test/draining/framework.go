// +build k8srequired

package draining

import (
	"encoding/base64"
	"io/ioutil"
	"os"

	"github.com/giantswarm/e2e-harness/pkg/framework"
	"github.com/giantswarm/microerror"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func InstallGuestChart(name, values string, customWaitFor func() error) error {
	tmpfile, err := ioutil.TempFile("", name+"-values")
	if err != nil {
		return microerror.Mask(err)
	}
	defer os.Remove(tmpfile.Name())
	_, err = tmpfile.Write([]byte(os.ExpandEnv(values)))
	if err != nil {
		return microerror.Mask(err)
	}

	err = framework.HelmCmd("registry install quay.io/giantswarm/" + name + "-chart:stable -- -n " + name + " --values " + tmpfile.Name())
	if err != nil {
		return microerror.Mask(err)
	}

	err = waitFor(customWaitFor)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func SetupGuestKubeConfig() (string, string, error) {
	secretName := os.ExpandEnv("${CLUSTER_NAME}-api")
	secret, err := f.cs.CoreV1().Secrets("default").Get(secretName, metav1.GetOptions{})
	if err != nil {
		return "", "", microerror.Mask(err)
	}

	var config = `apiVersion: v1
kind: Config
clusters:
- cluster:
    certificate-authority-data: ` + mustDecodeBase64(secret.Data["ca"]) + `
    server: ` + os.ExpandEnv("https://api.${CLUSTER_NAME}.${COMMON_DOMAIN_GUEST}") + `
  name: giantswarm-guest
contexts:
- context:
    cluster: giantswarm-guest
    user: giantswarm-guest
  name: giantswarm-guest
current-context: giantswarm-guest
users:
- name: giantswarm-guest
  user:
    client-certificate-data: ` + mustDecodeBase64(secret.Data["crt"]) + `
    client-key-data: ` + mustDecodeBase64(secret.Data["key"])

	tmpfile, err := ioutil.TempFile("", "guest-kube-config")
	if err != nil {
		return "", "", microerror.Mask(err)
	}
	_, err = tmpfile.Write([]byte([]byte(config)))
	if err != nil {
		return "", "", microerror.Mask(err)
	}

	tmpFileName := tmpfile.Name()
	currentKubeConfig := os.Getenv("KUBECONFIG")

	os.Setenv("KUBECONFIG", tmpFileName)

	return tmpFileName, currentKubeConfig, nil
}

func TeardownGuestKubeConfig(tmpFileName, oldKubeConfig string) error {
	os.Setenv("KUBECONFIG", oldKubeConfig)

	err := os.Remove(tmpFileName)
	if err != nil {
		return microerror.Mask(err)
	}

	return nil
}

func mustDecodeBase64(value []byte) string {
	decoded, err := base64.StdEncoding.DecodeString(string(value))
	if err != nil {
		panic(err)
	}

	return string(decoded)
}

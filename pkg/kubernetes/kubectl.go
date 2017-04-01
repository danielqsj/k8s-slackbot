package kubernetes

import (
	"bytes"
	"fmt"
	"strings"

	"k8s.io/kubernetes/pkg/kubectl/cmd"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

// RunKubectl run kubectl command and return result
func RunKubectl(kubeconfig string, args []string) (string, error) {
	var b bytes.Buffer
	cmdOut := &b
	cmdErr := &b
	cmd := cmd.NewKubectlCommand(cmdutil.NewFactory(nil), strings.NewReader(""), cmdOut, cmdErr)
	if len(kubeconfig) > 0 {
		args = AppendKubeconfig(kubeconfig, args)
	}
	cmd.SetArgs(args)
	err := cmd.Execute()
	return fmt.Sprintf("%s\n", b.String()), err
}

// AppendKubeconfig append param kubeconfig
func AppendKubeconfig(kubeconfig string, args []string) []string {
	newKubeconfig := "--kubeconfig=" + kubeconfig
	return append(args, newKubeconfig)
}

// ConnectMaster connect to kubernetes master
func ConnectMaster(kubeconfig string) (string, error) {
	args := []string{"version"}
	return RunKubectl(kubeconfig, args)
}

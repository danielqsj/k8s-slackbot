package kubernetes

import (
	"bytes"
	"fmt"
	"strings"

	"k8s.io/kubernetes/pkg/kubectl/cmd"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

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

func AppendKubeconfig(kubeconfig string, args []string) []string {
	newKubeconfig := "--kubeconfig=" + kubeconfig
	return append(args, newKubeconfig)
}

func TestConnection(kubeconfig string) (string, error) {
	args := []string{"version"}
	return RunKubectl(kubeconfig, args)
}

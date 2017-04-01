package kubernetes

import (
	"testing"
	"reflect"
)

func TestAppendKubeconfig(t *testing.T) {
	testCases := []struct {
		kubeconfig      string
		args            []string
		expected        []string
	}{
		{"/etc/kubernetes/kubeconfig", []string{"--debug=true"}, []string{"--debug=true", "--kubeconfig=/etc/kubernetes/kubeconfig"}},
	}

	for _, testCase := range testCases {
		result := AppendKubeconfig(testCase.kubeconfig, testCase.args)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("expected %v but return %v", testCase.expected, result)
		}
	}
}
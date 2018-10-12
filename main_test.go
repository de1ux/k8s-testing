package main

import (
	"github.com/de1ux/k8s-testing/operations"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestPrintPods(t *testing.T) {

	cs := fake.NewSimpleClientset(&corev1.Pod{
		ObjectMeta: v1.ObjectMeta{
			Labels: map[string]string{
				"app": "nginx",
			},
		},
	})

	c := &operations.Client{
		Clientset: cs,
	}
	if err := operations.PrintPods(c, "nginx"); err != nil {
		t.Fatal(err)
	}
}

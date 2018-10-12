package operations

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"log"
)

type Client struct {
	Clientset kubernetes.Interface
}

func PrintPods(c *Client, app string) error {
	pods, err := c.Clientset.CoreV1().Pods("default").List(metav1.ListOptions{LabelSelector: "app=" + app})
	if err != nil {
		return err
	}

	for _, i := range pods.Items {
		log.Print(i)
	}
	return nil
}

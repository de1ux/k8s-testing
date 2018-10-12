package main

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path/filepath"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", filepath.Join(homeDir(), ".kube", "config"))
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	_, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
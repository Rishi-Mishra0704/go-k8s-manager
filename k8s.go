package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func deployToK8s() {
	kubeconfig := filepath.Join(homeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// Define the pod
	pod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "hello-user-pod",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:  "hello-user-container",
					Image: "rishimishra0704/go-k8s-manager:latest",
					Command: []string{
						"./go-k8s-manager",
						"--action=serve",
					},
					Ports: []v1.ContainerPort{
						{
							ContainerPort: 8080,
						},
					},
				},
			},
		},
	}

	// Create the pod
	podClient := clientset.CoreV1().Pods("default")
	result, err := podClient.Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Created pod %q.\n", result.GetObjectMeta().GetName())
}

// helper function to get home directory
func homeDir() string {
	if h := flag.Lookup("home"); h != nil {
		return h.Value.String()
	}
	return "/"
}

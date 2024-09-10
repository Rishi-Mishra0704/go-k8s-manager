package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	"github.com/prometheus/client_golang/prometheus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	deploymentStatus = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "k8s_deployment_status",
			Help: "Status of Kubernetes deployments",
		},
		[]string{"name"},
	)
)

func init() {
	prometheus.MustRegister(deploymentStatus)
}

func MonitorDeployments(namespace string) {
	kubeconfig := filepath.Join(homeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating Kubernetes client: %v", err)
	}

	deploymentClient := clientset.AppsV1().Deployments(namespace)
	deployments, err := deploymentClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Error listing deployments: %v", err)
	}

	for _, deployment := range deployments.Items {
		name := deployment.Name
		replicas := deployment.Status.Replicas
		availableReplicas := deployment.Status.AvailableReplicas

		if replicas != availableReplicas {
			deploymentStatus.WithLabelValues(name).Set(0) // Unhealthy
		} else {
			deploymentStatus.WithLabelValues(name).Set(1) // Healthy
		}

		fmt.Printf("Deployment %s status updated.\n", name)
	}
}

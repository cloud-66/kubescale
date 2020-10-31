package main

import (
	"context"
	"log"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Fatal(err)
		}
	}

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	s, err := client.AppsV1().
		Deployments("default").
		GetScale(context.TODO(), "nginx", metav1.GetOptions{})
	if err != nil {
		log.Fatal(err)
	}

	sc := *s
	sc.Spec.Replicas = 0

	us, err := client.AppsV1().
		Deployments("default").
		UpdateScale(context.TODO(),
			"nginx", &sc, metav1.UpdateOptions{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(*us)
}

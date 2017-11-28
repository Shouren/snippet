/*
Copyright 2016 The Kubernetes Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	_ "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
	var kubeconfig *string
	if home := homeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		services, err := clientset.CoreV1().Services("").List(metav1.ListOptions{FieldSelector: "metadata.namespace=demo"})
		// services, err := clientset.CoreV1().Services("").List(metav1.ListOptions{LabelSelector: "kubernetes.io/cluster-service=true"})
		// services, err := clientset.CoreV1().Services("").List(metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d services in the cluster\n", len(services.Items))

		for _, svc := range services.Items {
			fmt.Printf("Namespace: %s, Services: %s, Label: %s\n", svc.Namespace, svc.Name, svc.Labels)
			for _, port := range svc.Spec.Ports {
				if port.Protocol != "TCP" {
					continue
				}
				fmt.Printf("  Name: %v, Proto: %v , Port: %v, TargetPort: %v, NodePort: %v \n", port.Name, port.Protocol, port.Port, port.TargetPort, port.NodePort)
				fmt.Printf("  URL: http://%v:%v\n", svc.Name, port.Port)
			}
		}
		fmt.Println(os.Getenv("HOME"))
		fmt.Println(*kubeconfig)

		break
		time.Sleep(10 * time.Second)
	}
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

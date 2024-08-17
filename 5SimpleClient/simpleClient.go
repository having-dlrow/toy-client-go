package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"k8s.io/client-go/rest"
	"path/filepath"
	"time"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

var (
	masterURL string
)

func main() {
	clientSet := getk8sClient()
	for {
		getAllPod(clientSet)
		getDefaultNamespacePod(clientSet)
		time.Sleep(10 * time.Second)
	}

}

func getDefaultNamespacePod(clientSet *kubernetes.Clientset) {
	// Examples for error handling:
	// - Use helper functions like e.g. errors.IsNotFound()
	// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	namespace := "default"
	pod := "k8s-"
	_, err := clientSet.CoreV1().Pods(namespace).Get(context.TODO(), pod, metav1.GetOptions{})
	if errors.IsNotFound(err) {
		fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
		fmt.Printf("Error getting pod %s in namespace %s: %v\n",
			pod, namespace, statusError.ErrStatus.Message)
	} else if err != nil {
		panic(err.Error())
	} else {
		fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	}
}

func getAllPod(clientSet *kubernetes.Clientset) {
	// 전체 조회
	pods, err := clientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

	for _, pod := range pods.Items {
		fmt.Printf("Namespace: %s, Pod Name: %s \n", pod.Namespace, pod.Name)
		//fmt.Printf("Spec: %v \n", pod.Spec)
		//itemPod, _ := json.Marshal(pod)
		//fmt.Printf("items: %s", jsonPrettyPrint(string(itemPod)))
	}
}

func getk8sClient() *kubernetes.Clientset {
	config, _ := setConfig()
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientSet
}

func setConfig() (*rest.Config, error) {
	// Try to create in-cluster config
	// Try to create in-cluster config
	if config, err := rest.InClusterConfig(); err == nil {
		return config, nil
	}

	var kubeConfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeConfig file")
	} else {
		kubeConfig = flag.String("kubeConfig", "", "absolute path to the kubeConfig file")
	}
	flag.Parse()

	// use the current context in kubeConfig
	config, err := clientcmd.BuildConfigFromFlags(masterURL, *kubeConfig)
	if err != nil {
		panic(err.Error())
	}
	return config, err
}

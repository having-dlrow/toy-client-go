package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"k8s.io/client-go/util/retry"
	"os"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

func main() {

	// kubeconfig file flag
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// Build the config from the kubeconfig file
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	// Create the dynamic client
	client, err := dynamic.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	// Define GVR (GroupVersionResource) for pods
	deploymentRes := schema.GroupVersionResource{
		Group:    "apps",
		Version:  "v1",
		Resource: "deployments",
	}

	// List pods in the default namespace
	deploymentName := "demo-deployment"
	deploymentNamespace := apiv1.NamespaceDefault
	deployment := &unstructured.Unstructured{
		Object: map[string]interface{}{
			"apiVersion": "apps/v1",
			"kind":       "Deployment",
			"metadata": map[string]interface{}{
				"name": deploymentName,
			},
			"spec": map[string]interface{}{
				"replicas": 2,
				"selector": map[string]interface{}{
					"matchLabels": map[string]interface{}{
						"app": "demo",
					},
				},
				"template": map[string]interface{}{
					"metadata": map[string]interface{}{
						"labels": map[string]interface{}{
							"app": "demo",
						},
					},

					"spec": map[string]interface{}{
						"containers": []map[string]interface{}{
							{
								"name":  "web",
								"image": "nginx:1.12",
								"ports": []map[string]interface{}{
									{
										"name":          "http",
										"protocol":      "TCP",
										"containerPort": 80,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Create Deployment
	existingDeployment, err := client.Resource(deploymentRes).Namespace(deploymentNamespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		fmt.Println("Creating deployment...")
		result, err := client.Resource(deploymentRes).Namespace(apiv1.NamespaceDefault).Create(context.TODO(), deployment, metav1.CreateOptions{})
		if err != nil {
			panic(err)
		}
		fmt.Printf("Created deployment %q.\n", result.GetName())
	} else {
		prompt()
		fmt.Println("Updating deployment...")

		retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {

			//result, err := client.Resource(deploymentRes).Namespace(apiv1.NamespaceDefault).Get(context.TODO(), deploymentName, metav1.GetOptions{})
			//if err != nil {
			//	panic(fmt.Errorf("failed to get latest version of Deployment : %v", err))
			//}

			// update replicas to 1
			if err := unstructured.SetNestedField(existingDeployment.Object, int64(1), "spec", "replicas"); err != nil {
				panic(fmt.Errorf("failed to set replica value: %v", err))
			}

			_, updateErr := client.Resource(deploymentRes).Namespace(apiv1.NamespaceDefault).Update(context.TODO(), existingDeployment, metav1.UpdateOptions{})
			return updateErr
		})
		if retryErr != nil {
			panic(fmt.Errorf("update failed: %v", retryErr))
		}
		fmt.Println("Updated deployment...")
	}
}

func prompt() {
	fmt.Printf("-> Press Return key to continue.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		break
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println()
}

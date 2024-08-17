package main

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"k8s.io/client-go/rest"
	"reflect"
	"testing"
)

func TestGetDefaultNamespacePod(t *testing.T) {
	//clientSet := getk8sClient()
	clientSet := fake.NewSimpleClientset()

	// Create a pod in the default namespace
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "example-pod",
			Namespace: "default",
		},
	}

	_, err := clientSet.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("error creating pod: %v", err)
	}

}

func TestGetAllPod(t *testing.T) {
	//clientSet := getk8sClient()
	clientSet := fake.NewSimpleClientset()

	// Create a pod in the default namespace
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "example-pod",
			Namespace: "default",
		},
	}

	_, err := clientSet.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		t.Fatalf("error creating pod: %v", err)
	}

}

func Test_setConfig(t *testing.T) {
	tests := []struct {
		name    string
		want    *rest.Config
		wantErr bool
	}{
		{name: "Test Case1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := setConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("setConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}

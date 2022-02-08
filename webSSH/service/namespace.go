package service

import (
	"context"
	"gin-client-go.com/webSSH/client"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func GetNamespaces() ([]v1.Namespace, error) {
	ctx := context.Background()
	clientSet, err := client.GetK8SClientSet()
	if err != nil {
		klog.Fatal(err)
		return nil, err
	}
	namespaceList, err := clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Fatal(err)
		return nil, err
	}
	return namespaceList.Items, nil

}

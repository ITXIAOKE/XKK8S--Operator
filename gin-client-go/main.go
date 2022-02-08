package main

import (
	"context"
	"flag"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
	"path/filepath"
)

//打印kubernetes中的namespace
func main() {
	fmt.Println("hello k8s")
	var kubeConfig *string
	ctx := context.Background()
	if home := homedir.HomeDir(); home != "" {
		kubeConfig = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "数据正确k8s")
	} else {
		kubeConfig = flag.String("kubeConfig", "", "没有数据")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		klog.Fatal(err)
		return
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
		return
	}
	//调用namespace
	namespaceList, err := clientSet.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	if err != nil {
		klog.Fatal(err)
		return
	}
	namespaces := namespaceList.Items
	for _, namespace := range namespaces {
		fmt.Println("==>"+namespace.Name+"-->"+string(namespace.Status.Phase))
	}
	//调用pod
	podList, err := clientSet.CoreV1().Pods(namespaces[0].Name).List(ctx, metav1.ListOptions{})
	if err!=nil{
		klog.Fatal(err)
		return
	}
	pods:= podList.Items
	for _, pod := range pods {
		fmt.Println(pod.Name)
	}
	//调用其他
	//clientSet.CoreV1().Nodes()

}

package main

import (
	"fmt"
	"github.com/ielepro/k8s-client-common/common"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func main() {
	var (
		clientSet *kubernetes.Clientset
		podsList  *coreV1.PodList
		err       error
	)

	// 初始化k8s客户端
	if clientSet, err = common.InitClient(); err != nil {
		fmt.Println(err)
		return
	}

	// 获取default命名空间下的所有POD
	if podsList, err = clientSet.CoreV1().Pods("default").List(metaV1.ListOptions{}); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*podsList)
	return
}
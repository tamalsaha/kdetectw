package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"kmodules.xyz/client-go/dynamic"
	"log"
	"path/filepath"
)

func main() {
	masterURL := ""
	kubeconfigPath := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfigPath)
	if err != nil {
		log.Fatalln(err)
	}

	gvr := schema.GroupVersionResource{
		Group:    "",
		Version:  "v1",
		Resource: "pods",
	}

	obj, g2, err := dynamic.DetectWorkload(config, gvr, "kube-system", "coredns-86c58d9df4-8tfzp")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(g2.String())
	fmt.Println(obj)
}

package main

import (
	"fmt"
	"os"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

type HelloWorld struct {
	v1.TypeMeta `json:",inline"`
	Metadata    v1.ObjectMeta `json:"metadata"`
}

func (h *HelloWorld) DeepCopyObject() runtime.Object {
	return &HelloWorld{
		TypeMeta: h.TypeMeta,
		Metadata: h.Metadata,
	}
}
func main() {
	config, err := getConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	hwListWatcher := cache.NewListWatchFromClient(
		clientset.RESTClient(),
		"helloworlds",
		v1.NamespaceDefault,
		fields.Everything(),
	)

	_, controller := cache.NewInformer(
		hwListWatcher,
		&HelloWorld{},
		time.Minute*10,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				fmt.Printf("New HelloWorld Added: %s\n", obj.(*HelloWorld).Metadata.Name)
			},
		},
	)

	stop := make(chan struct{})
	defer close(stop)
	go controller.Run(stop)

	select {} // run forever
}

func getConfig() (*rest.Config, error) {
	if kubeconfig := os.Getenv("KUBECONFIG"); kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}

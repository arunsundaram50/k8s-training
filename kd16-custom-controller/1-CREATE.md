# Define a simple yet practical CRD and its associated Go-based controller.

### Use Case: `AutoSecret`

Automatically generating a secret containing random alphanumeric strings for applications that require credentials (like database passwords, API keys, etc.).

### CRD: `AutoSecret`

**Structure:**
```yaml
apiVersion: custom.k8s.io/v1
kind: AutoSecret
metadata:
  name: my-db-secret
spec:
  length: 16
```

- `length`: Length of the random string to generate.

**Controller Behavior:**

When an `AutoSecret` resource is created:
1. The controller generates a random alphanumeric string of the specified length.
2. It creates a Kubernetes `Secret` resource with the generated string.
3. It updates the `AutoSecret` status to indicate the `Secret` has been created.

### Simple Go-based Controller:

Here's a very simplified version of this controller. Note that it omits error handling and optimizations for brevity:

```go
package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/rest"
)

type AutoSecret struct {
	v1.TypeMeta   `json:",inline"`
	Metadata      v1.ObjectMeta `json:"metadata"`
	Spec          AutoSecretSpec `json:"spec"`
}

type AutoSecretSpec struct {
	Length int `json:"length"`
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
		clientset.CoreV1().RESTClient(),
		"autosecrets",
		v1.NamespaceAll,
		cache.ResourceEventHandlerFuncs{},
	)

	_, controller := cache.NewInformer(
		hwListWatcher,
		&AutoSecret{},
		time.Minute*10,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				// Generate a random alphanumeric string
				randomData := make([]byte, obj.(*AutoSecret).Spec.Length)
				rand.Read(randomData)
				encodedSecret := base64.StdEncoding.EncodeToString(randomData)

				// Create a Kubernetes Secret with the generated data
				clientset.CoreV1().Secrets(obj.(*AutoSecret).Metadata.Namespace).Create(&v1.Secret{
					ObjectMeta: v1.ObjectMeta{
						Name: obj.(*AutoSecret).Metadata.Name,
					},
					Data: map[string][]byte{
						"key": []byte(encodedSecret),
					},
				})

				fmt.Printf("AutoSecret %s processed and Secret created\n", obj.(*AutoSecret).Metadata.Name)
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
```

This simple controller listens for the creation of `AutoSecret` custom resources, generates a random alphanumeric string based on the specified length, and creates a Kubernetes `Secret` containing this generated string. It's a basic example, but in real-world scenarios, additional features like error handling, secret rotations, and more would be beneficial.
## Create a standalone pod --virtually, the `docker run` equivalent

### Build & publish the image
```bash
docker image build -t docker.io/arunsundaramco70/myhello .
docker push docker.io/arunsundaramco70/myhello
```

### Run a pod with that image
```bash
kubectl run demo --image=arunsundaramco70/myhello --port 9999 --labels app=demo
```

The command will run a new Kubernetes pod with a specified image and port number, along with a label. Here's a breakdown of what each part does:

1. `kubectl run demo`: This will start a new pod named "demo."

2. `--image=arunsundaramco70/myhello`: The pod will be created using the Docker image found at the specified location, `arunsundaramco70/myhello`.

3. `--port 9999`: This will expose port 9999 on the container, which can later be used to forward traffic to the pod using a service or other means.

4. `--labels app=demo`: This attaches a label to the pod with the key "app" and the value "demo." Labels are key-value pairs that are used to select and manage groups of resources within Kubernetes.

Keep in mind that as of Kubernetes 1.18, the `kubectl run` command is recommended for running a single container in a pod, and more complex workloads should be managed with `kubectl create` or `kubectl apply` using YAML or JSON manifests. Also, the `kubectl run` command alone won't make the pod accessible from outside the cluster. You would usually create a service to expose the pod externally or use port forwarding if you want to interact with it from outside the Kubernetes cluster.

### Get info
```bash
kubectl get pods --selector app=demo
kubectl describe pod demo
```

### Visiting the site
```bash
kubectl port-forward demo 9999:8888
```

Go to <http://localhost:9999>

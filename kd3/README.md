# k3. Running a simple application

Instead of jumping right into k8s, we will deploy the application starting in a simple way and incrementally we will do a k8s deploy. This will help us in many ways
- allow us to test the application before making it part of a k8s deploy
- when issue occur, we would know the problem is likely to be the incremental steps
- help us understand what we gain with the increased complexity and what k8s is offerning us

Also, we will summarize what we learn in each levels in the following table:
#### What we learned as an Operator:
| Issue | Host | Venv | Docker | K8S (imp)| K8S (dec) |
|-|-|-|-|-|-|
|Disowning Library Installation |x|x|√|
|If using Python, user will have to use virtual environment to avoid library conflict with other applications running in the same host/machine|x|√|√|
|Runs in an environment seperate from the host|x|√ §1|√|
|Oversight to bring service back up (if it crashes)|x|x|x|
|There is oversight if it behaves badly (like take 100% CPU)|x|x|√|
|There is way to scale it up|x|x|x|
|Even if I run multiple instances to cope with the demand, scaling down is still manual|x|x|x|
|Updating the application causes downtime|x|x|x|

§1: Create a `virtual environment` so that the Python libraries installed don't mess with the host's library versions
```bash
python3 -m venv venv_v1
source venv_v1/bin/activate
pip3 install -r requirements.txt
python3 ./main.py
```

### 1) Runnning the microservice in the simplest possible way
To run a simple Python microservice defined in `main.py` that serves one simple endpoint `/hello`.
```bash
python3 ./main.py
```

 
### 2) Running it as a Docker container
- In order to run it as a container, we first have to define a `Dockerfile` which describes (declaratively lists) how the image has to be built (typically by developers)
```bash
docker image build -t arunsundaramco70/kd1 .
docker image push arunsundaramco70/kd1
```
- The devops folks can use/run this image, like so:
```bash
docker container run arunsundaramco70/kd1
```

#### Gotchas
- The command `docker container run arunsundaramco70/kd1` runs a container it a different network as compared to the host machine's network.
- In order to communicate with the container, we will have to port-forward, like so:
```bash
docker container run -p8001:8001 arunsundaramco70/kd1 
docker container run -p8002:8001 arunsundaramco70/kd1 
```
- To limit resource usage, one can specifiy the limits. For instance, to limit CPU, one can do like so:
```bash
docker container run -p8001:8001 --cpu-shares=10 arunsundaramco70/kd1 
```

### 3) Running the application in K8S (imperative)
```bash
kubectl run kd1 --image=arunsundaramco70/kd1 --port 8001 --labels app=kd1
kubectl port-forward pod/kd1 8001
```

### 3a) Running multiple replicas (as a deployment) and exposing the deployment
```bash
kubectl create deployment kd3-deployment --image=arunsundaramco70/kd3 --replicas=3
kubectl expose deployment kd3-deployment --name=kd3-service --port=8001 --type=LoadBalancer
```

### 4) Running the application in K8S (declarative) --i.e. using a deployment object declared in an YAML file, say `deployment.yaml`
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-deployment
spec:
  replicas: 3 # POD replicas
  selector:
    matchLabels:
      app: my-pod
  template: # POD template
    metadata:
      labels:
        app: my-pod # POD label
    spec: # Container specifications
      containers:
      - name: my-container
```

To apply the above deployment:
```bash
kubectl apply -f deployment.yaml
```

Here are some commands to manage the deployment:
```bash
kubectl get deployment my-deployment
kubectl get pods
kubectl describe deployment my-deployment
kubectl delete deployment my-deployment
# or
kubectl delete -f deployment.yaml
```

## 5) Fronting the application with a service
Though we created a set of pods, we are not able to reach it. This is where Service comes in.
Apply the service
```bash
kubectl apply -f service.yaml
```

The specification is as follows:
- `port`: This is the port that will be exposed by the service, i.e., the port that other services in the cluster will use to communicate with this service.
- `targetPort`: This is the port on the pod where the application is running and listening for incoming requests. It's the port that the service will forward requests to.
- `nodePort`: This is the port on each node where the service will be exposed. External traffic coming into the cluster will come in through this port.

Here, `port: 80` means that within the Kubernetes cluster, other services or applications would reach our `myhello-service` on port 80. This port could be any valid port number (not necessarily matching the `targetPort`) as per our application architecture and communication plan within the Kubernetes cluster.


Find the Cluster-IP like so:
```bash
kubectl get service xyz-service
```

Here is a sample output:
```txt
NAME          TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
xyz-service   LoadBalancer   10.99.206.163   localhost     8000:30146/TCP   17d
```

For this example, you can reach the application like so:
```
http://localhost:8000/hello
```

# [Services in detail](https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types)

## Types of services
### ClusterIP
Exposes the Service on a cluster-internal IP. Choosing this value makes the Service only reachable from within the cluster. This is the default that is used if you don't explicitly specify a type for a Service. You can expose the Service to the public internet using an Ingress or a Gateway.
### NodePort
Exposes the Service on each Node's IP at a static port (the NodePort). To make the node port available, Kubernetes sets up a cluster IP address, the same as if you had requested a Service of type: ClusterIP.
### LoadBalancer
Exposes the Service externally using an external load balancer. Kubernetes does not directly offer a load balancing component; you must provide one, or you can integrate your Kubernetes cluster with a cloud provider.
### ExternalName
Maps the Service to the contents of the externalName field (for example, to the hostname api.foo.bar.example). The mapping configures your cluster's DNS server to return a CNAME record with that external hostname value. No proxying of any kind is set up.

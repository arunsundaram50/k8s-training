# Running a simple application

We will see how to run a simple Python application `main.py` that serves one simple endpoint `/hello`.

Instead of jumping right into k8s, we will deploy the application starting in a simple way and incrementally we will do a k8s deploy. This will help us in many ways
- allow us to test the application before making it part of a k8s deploy
- when issue occur, we would know the problem is likely to be the incremental steps
- help us understand what we gain with the increased complexity and what k8s is offerning us

## 1) Runnning it as a Python script
```bash
python3 ./main.py
```
### Main issues (from Operators perspective):
| Issue | Host | Venv | Docker | K8S (imp)| K8S (dec) |
|-|-|-|-|-|-|
|Disowning Library Installation |x|x|√|
|If using Python, user will have to use virtual environment to avoid library conflict with other applications running in the same host/machine|x|√|√|
|||Creating §1||
|Runs in an environment seperate from the host|x|√|√|
|Oversight to bring service back up (if it crashes)|x|x|x|
|There is oversight if it behaves badly (like take 100% CPU)|x|x|√|
|There is way to scale it up|x|x|x|
|Even if I run multiple instances to cope with the demand, scaling down is still manual|x|x|x|
|Updating the application causes downtime|x|x|x|

§1: 
```bash
python3 -m venv venv_v1
source venv_v1/bin/activate
python3 ./main.py
```
 
## 2) Run it as a Docker container
- In order to run it as a container, we first have to define a `Dockerfile` which describes (declaratively lists) how the image has to be built (typically by developers)
```bash
docker image build -t arunsundaramco70/kd1 .
docker image push arunsundaramco70/kd1
```
- The devops folks can use/run this image, like so:
```bash
docker container run arunsundaramco70/kd1
```
Gotchas
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

## 3) Running application in K8S (imperative)
```bash
kubectl run kd1 --image=arunsundaramco70/kd1 --port 8001 --labels app=kd1
kubectl port-forward pod/kd1 8001
```

## 4) Running the application using a object declared in an YAML file
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

```bash
kubectl apply -f deployment.yaml
kubectl get deployment my-deployment
kubectl get pods
kubectl describe deployment my-deployment
kubectl delete deployment my-deployment
# or
kubectl delete -f deployment.yaml
```

## 5) Fronting the applicaiton with a service
Though we created a set of pods, we are not able to reach it. This is where Service comes in.
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

# k2. A Tour of Deploying Application Using K8S

We will deploy the application using kd1/deployment.yaml.
This will create an unreachable set of pods.
To be able to reach the pod's containers, we will "front" the pods with a service.


# Command Line Deployment and Service Creation
### make sure the image is published to the hub
```
cd hello
docker image build -t docker.io/arunsundaramco70/hello .
docker push docker.io/arunsundaramco70/hello
```

### delete the standalone pod and create a deployment
#### delete, if deployment exists:  kubectl delete deployment demo 
```
kubectl delete pod hello-pod
kubectl create deployment hello-deployment --image=arunsundaramco70/hello --port 8001
kubectl describe deployment hello-deployment
```

### List services. Notice no load balancer is running, and there is no way to reach 8001
(delete, if services aleady exists: kubectl delete services hello-service)
```
kubectl get services
```


### Expose service without LB to 8001
```
kubectl expose deployment hello-deployment --type=NodePort --name=hello-service --port=80 --target-port=8001
```
Note:
- this doesn't expose the service to an IP as the service type is `NodePort`. §1
- change the type to `--type=LoadBalancer` and see how you are able to hit the application using `localhost`


#### §1: Find the Cluster-IP & use the port and hit it from the browser 
```
kubectl get service hello-service
```
For example, if you see
```
NAME              TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
kubernetes        ClusterIP   10.96.0.1       <none>        443/TCP        86m
hello-service   NodePort    10.104.191.51   <none>        80:30658/TCP   5m12s
```
we would hit <http://localhost:30658>

## Enabling inter-pod communication
Examine and use the `deployment.yaml` that allows for the `upper-pod` to talk to `hello-pod`

# kd2: A Tour of Deploting Application Using K8S

We will deploy the application using kd1/deployment.yaml.
This will create an unreachable set of pods.
To be able to reach the pod's containers, we will "front" the pods with a service.


# Command Line Deployment and Service Creation
## Create a deployment

### make sure the image is in hub
```
docker image build -t docker.io/arunsundaramco70/myhello .
docker push docker.io/arunsundaramco70/myhello
```

### delete the standalone pod and create a deployment
#### delete, if deployment exists:  kubectl delete deployment demo 
```
kubectl delete pod demo
kubectl create deployment demo --image=arunsundaramco70/myhello --port 8888
kubectl describe deployment demo
```

### List services, see no load balancer is running, and there is no way to reach 8888
#### delete, if services aleady exists: kubectl delete services demo
```
kubectl get services
```


### Expose service without LB to 8888
```
kubectl expose deployment demo --type=NodePort --name=myhello-service --port=80 --target-port=8888
```

### Find the Cluster-IP & use the port and hit it from the browser 
```
kubectl get service myhello-service
```
For example, if you see
```
NAME              TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
kubernetes        ClusterIP   10.96.0.1       <none>        443/TCP        86m
myhello-service   NodePort    10.104.191.51   <none>        80:30658/TCP   5m12s
```
we would hit <http://localhost:30658>

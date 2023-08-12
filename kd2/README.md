# k2. [Services](https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types)

## Types of services
### ClusterIP
Exposes the Service on a cluster-internal IP. Choosing this value makes the Service only reachable from within the cluster. This is the default that is used if you don't explicitly specify a type for a Service. You can expose the Service to the public internet using an Ingress or a Gateway.
### NodePort
Exposes the Service on each Node's IP at a static port (the NodePort). To make the node port available, Kubernetes sets up a cluster IP address, the same as if you had requested a Service of type: ClusterIP.
### LoadBalancer
Exposes the Service externally using an external load balancer. Kubernetes does not directly offer a load balancing component; you must provide one, or you can integrate your Kubernetes cluster with a cloud provider.
### ExternalName
Maps the Service to the contents of the externalName field (for example, to the hostname api.foo.bar.example). The mapping configures your cluster's DNS server to return a CNAME record with that external hostname value. No proxying of any kind is set up.


## Create a service
- Create a service to expose the deployment, with another file named `service.yaml`
- Deploy the service using kubectl
```
kubectl apply -f service.yaml
```

Now, our webapp should be accessible from a browser at localhost:30000. The reason we're able to access it on port 30000 is due to the service type being "NodePort", which exposes the service on a static port on each node. In a production setting, we would typically use a different service type, such as LoadBalancer or Ingress, to expose our application.

Note: The `port` field in the service manifest determines the port number on which the service itself is exposed internally within the Kubernetes cluster. It is essentially the port that other services or applications within the cluster will use to communicate with this service.

The specification is as follows:

- `port`: This is the port that will be exposed by the service, i.e., the port that other services in the cluster will use to communicate with this service.
- `targetPort`: This is the port on the pod where the application is running and listening for incoming requests. It's the port that the service will forward requests to.
- `nodePort`: This is the port on each node where the service will be exposed. External traffic coming into the cluster will come in through this port.

Here, `port: 80` means that within the Kubernetes cluster, other services or applications would reach our `myhello-service` on port 80. This port could be any valid port number (not necessarily matching the `targetPort`) as per our application architecture and communication plan within the Kubernetes cluster.

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

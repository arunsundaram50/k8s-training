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

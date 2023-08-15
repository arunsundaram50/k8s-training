# K8s Command Line format or K8s basic architecture

Imperative
```
kubectl <OPERATION> <OBJECT> <OBJECT_ID> <OPTIONS>
```
OPEARATIONS: get, apply, delete, describe, add, patch, edit
OBJECT: Pod, Node, Cluster, Service, Ingress, Deployment, LB, ...

Declarive
Create a `pod.yaml`
Like Dockerfile, we will describe the containers
```
kubectl apply -f pod.yaml
```

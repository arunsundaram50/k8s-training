# K8s Command Line format

### Imperative commands are of the following format
```
kubectl <OPERATION> <OBJECT> <OBJECT_ID> <OPTIONS>
```
- OPEARATIONS: get, apply, delete, describe, add, patch, edit
- OBJECT: Pod, Node, Cluster, Service, Ingress, Deployment, LB, ...
OPERATIONS loosely resemble HTTP methods

### Declarive commands 
Create a YAML file, for instance: `pod.yaml`
```
kubectl apply -f pod.yaml
```
Once we apply this file, we can make changes and then apply again. 
K8s will evaluate the changes, and apply only the changes. K8s may decide to delete and create resources based on the kind of change made to the YAML file.

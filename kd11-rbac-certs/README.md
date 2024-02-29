
K8S | SUBJECT   |  RESOURCE (NOUN)   |  VERB |
|-|-|-|-|
English     |              Nanda  | hunted  | a bird |
Grammar classification |   noun   |   verb    |    subject |
K8s example | ServiceAccount | Pod | delete |
K8s example | User Account | Node | create |


- SUBJECT: Role, User, Service account
- RESOURCE: Pod, Node, ConfigMap, Deployment, Service, Ingress, etc.
- VERB: get, delete, create, apply, etc.


### Setup a local user
#### Certificate Generation part
```bash
# localize
mkdir cert && cd cert
# Generate private key
openssl genrsa -out arunsundaram.key 2048
# Generate CSR
openssl req -new -key arunsundaram.key -out arunsundaram.csr -subj "/CN=arunsundaram/O=docker-training"
# Generate a Cert
openssl x509 -req -in arunsundaram.csr -CA ~/.minikube/ca.crt -CAkey ~/.minikube/ca.key -CAcreateserial -out arunsundaram.crt -days 364
```
#### K8s certification authentication part
##### Let k8s know about the user
```bash
kubectl config set-credentials arunsundaram --client-certificate=arunsundaram.crt --client-key=arunsundaram.key
```

##### Using the cert/arun
```bash
# Create a context. $CLUSTER_NAME would be docker-desktop or minikube
kubectl config set-context arunsundaram-context --cluster=$CLUSTER_NAME --user=arunsundaram
# To assume/use a context
kubectl config use-context arunsundaram-context
# To see the current context
kubectl config current-context
# To switch back to default context ($CONTEXT_NAME can be docker-desktop, minikube, etc.)
kk config set-context $CONTEXT_NAME
# To list all contexts
kk config get-contexts
# To delete
kk config delete-context arunsundaram-context
```

#### To create a role
```bash
kk create role pod-list-role -n kube-system --resource=pods --verb="get,list"
```
#### Bind the role to a subject (ex: user arunsundaram)
```bash
kk create rolebinding arunsundaram-pod-list-role-binding -n kube-system --role=pod-list-role --user=arunsundaram
```

#### Try these commands
```bash
kk get pods --namespace=kube-system # should work
kk describe pod/etcd-minikube --namespace=kube-system # should work
kk logs pod/etcd-minikube --namespace=kube-system # should NOT work
kk events pod/etcd-minikube --namespace=kube-system # should NOT work
kk delete pod/etcd-minikube --namespace=kube-system --dry-run=server # should NOT work, but ensure the role is not modified to allow delete
```

### Setup a serviceaccount
```bash
k create serviceaccount pod-manager
```

Create a role that can do everything with a pod
```bash
k create role pod-manager-role --verb="get,delete,create" --resource="pod"
```

And bind pod-manager-role to pod-manager, create a role-binding
```bash
k create rolebinding pmb --role=pod-manager-role --serviceaccount=default:pod-manager
```


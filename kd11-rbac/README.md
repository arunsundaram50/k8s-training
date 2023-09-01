
K8S | SUBJECT   |  RESOURCE (NOUN)   |  VERB |
|-|-|-|-|
English     |              Nanda  | hunted  | a bird |
Grammar classification |   noun   |   verb    |    subject |
K8s example | ServiceAccount | Pod | delete |
K8s example | User Account | Node | create |


- SUBJECT: Role, User, Service account
- RESOURCE: Pod, Node, ConfigMap, Deployment, Service, Ingress, etc.
- VERB: Describe, delete, create, apply, etc.


### Setup a local user
```bash
# localize
mkdir cert && cd cert
# Generate private key
openssl genrsa -out arunsundaram.key 2048
# Generate CSR
openssl req -new -key arunsundaram.key -out arunsundaram.csr -subj "/CN=arunsundaram/O=docker-training"
# Generate a Cert
openssl ×509 -req -in arunsundaram.csr -CA ~/.minikube/ca.crt -CAkey ~/.minikube/ca.key -CAcreateserial -out arunsundaram.crt -days 364
# 
kubectl config set-credentials arunsundaram --client-certificate=arunsundaram.crt --client-key=arunsundaram.key
# 
kubectl config set-context arunsundaram-context --cluster=MY-CLUSTER --user=arunsundaram
#
kubectl config use-context arunsundaram-context
#
kubectl config current-context
#
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


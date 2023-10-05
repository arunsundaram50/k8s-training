# k9. Ingress vs Service

## Service
Service name becomes a DNS Name (inside of K8S Network), and it gives a stable hostname to access a set of pods (based on the matchingLabels)
- Pod to pod communication
- Pod to Ingress 
- Can setup k8s-level security

## Ingress
Multiple hostnames and paths can refer to a stable DNS name (external to the k8s network)
```
http://www.my-server.com/hello -> hello-service
http://hello.my-server.com/ -> hello-service

http://www.my-server.com/upper -> upper-service
http://upper.my-server.com/ -> upper-service
```

It's roughly equivalent to AWS Route 55, Ngnix and Apache Web Server Routing


## Installing Ingress controller in an iMac
- brew install helm
- helm version
- kubectl config use-context docker-desktop
- helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
- kubectl create namespace ingress-controller
- helm install my-ingress ingress-nginx/ingress-nginx --namespace ingress-controller

### To check the status
- k get pods -n ingress-controller

### To delete
- helm uninstall my-ingress -n ingress-controller

### To "activate" it again
- helm install my-ingress ingress-nginx/ingress-nginx --namespace ingress-controller

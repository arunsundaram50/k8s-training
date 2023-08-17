# k1: Transitioning from Docker to K8S

## Prerequisite --have the image ready and test
In the world of docker, we build the image like so
```bash
docker image build -t arunsundaramco70/hello .
docker image push arunsundaramco70/hello
```

and run a built image like so
```bash
docker container run -p 8001:8001 --rm arunsundaramco70/hello
```

## To run this as a k8s pod & expose it on 8001
```bash
kubectl run hello-pod --image arunsundaramco70/hello --image-pull-policy=Always
kubectl expose pod hello-pod --port=8001 --name=hello-service --type=LoadBalancer
# http://localhost:8001/hello/peter
```

# Standalone Pods
## Create a standalone pod --virtually, the `docker run` equivalent
```
docker image build -t docker.io/arunsundaramco70/myhello .
docker push docker.io/arunsundaramco70/myhello

kubectl run demo --image=arunsundaramco70/myhello --port 9999 --labels app=demo
kubectl get pods --selector app=demo
kubectl describe pod demo

kubectl port-forward demo 9999:8888
```
Go to <http://localhost:9999>

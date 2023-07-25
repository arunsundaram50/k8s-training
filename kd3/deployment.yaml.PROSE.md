## Create a deployment
- Create `deployment.yaml` file and define a Kubernetes deployment named "myhello-deployment" that runs one replica of your Docker image, exposing port 8888 in the container.
- Deploy it to Kubernetes using kubectl:
```
kubectl apply -f deployment.yaml
```


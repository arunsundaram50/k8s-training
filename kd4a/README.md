# Config Maps

## Create and apply the ConfigMap
```bash
kubectl apply -f my-config-map.yaml 
```

## Create the pods (or deployments) that read the config map
```bash
kubectl apply -f my-pod.yaml
```

In the above case, there is only one pod.
ConfigMaps values are "sent" to the pods at the time the pods are created. Any updates to the ConfigMap will not not be seen by the pods that are already running. In order for the pods to see the updated values, you will have to delete them and apply them again. To have zero-downtime, create a deployment with more than one pod replica, and perform a rolling update.
```bash
kubectl -f my-deployment.yaml
```

Make changes to `my-config-map.yaml` and apply it
```bash
kubectl -f my-config-map.yaml
```

And perform a rolling update, like so:
```bash
kubectl rollout restart deployment my-deployment
```
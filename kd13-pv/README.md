```bash
kubectl apply -f pv.yaml
kubectl apply -f pvc.yaml
kubectl apply -f pod.yaml
```

verify the mount

```bash
kubectl exec -it mypod -- /bin/sh
# Inside the pod shell, list the contents of the /data directory
ls /data
```

PV [host dir/SAN/etc] -> PVC -> POD/Volume [PVC] -> Container VolumeMount [Volume] -> /data

pod1, pod2, pod3...
  ^
  |
Node 
  ^
  |
Kubernetes [Docker container]
  ^
  |
Docker Desktop 
  ^
  |
iMac 

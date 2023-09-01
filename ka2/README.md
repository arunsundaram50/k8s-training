## Installing Minikube

### Prep
```bash
sudo apt update && sudo apt upgrade -y
```

### get file
```bash
wget https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
```

### move it to executable location
```bash
sudo mv minikube-linux-amd64 /usr/local/bin/minikube
sudo chmod +x /usr/local/bin/minikube
```

### Check Minikube Installation
```bash
minikube version
```

### Start Minikube
```bash
minikube start --driver=docker
# or
sudo minikube start --driver=docker
```

### Check Kubernetes Cluster
```bash
kubectl get nodes
```
## Installing Minikube

### Prep the system
```bash
sudo apt update && sudo apt upgrade -y
```

### Get the executable file for Linux...
```bash
wget https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
```

#### ... and move it to executable location
```bash
sudo mv minikube-linux-amd64 /usr/local/bin/minikube
sudo chmod +x /usr/local/bin/minikube
```

### Check Minikube Installation
```bash
minikube version
```

## Start Minikube
### Ensure containerd is okay...
```bash
# Step 1: Validate containerd Status
sudo systemctl status containerd
# Step 2: Verify containerd Socket
ls /run/containerd/containerd.sock 
```

#### ...and start minikube control plane
```bash
# Step 3: Start Minikube with containerd
minikube start --container-runtime=containerd
# Step 4: Validate the Configuration
kubectl describe node minikube
```

### Stop
```bash
minikube stop
```


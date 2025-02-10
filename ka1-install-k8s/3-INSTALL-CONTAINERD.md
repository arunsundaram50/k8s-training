## How to install `containerd` on Ubuntu:

1. **Update and Install Dependencies**:

   First, update the package information and install required dependencies:
   
   ```bash
   sudo apt update && sudo apt upgrade -y
   sudo apt install -y curl apt-transport-https
   ```

2. **Set up the Docker Stable Repository**:

   `containerd` is bundled within Docker products, so you can use Docker's repositories to get the latest version:
   
   ```bash
   curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
   sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
   ```

3. **Install containerd**:

   Now, update the package database again and install `containerd`:
   
   ```bash
   sudo apt update
   sudo apt install -y containerd.io
   ```

4. **Start and Enable containerd Service**:

   Once the installation is complete, start the `containerd` service and enable it to start on boot:
   
   ```bash
   sudo systemctl start containerd
   sudo systemctl enable containerd
   ```

5. **Optional Configuration**:

   If you want to change the default configuration for `containerd`, you can do so by editing its configuration file, which is typically located at `/etc/containerd/config.toml`. After making changes, remember to restart the service:

   ```bash
   sudo systemctl restart containerd
   ```

6. **Verify Installation**:

   You can check if `containerd` is running with:

   ```bash
   sudo systemctl status containerd
   ```

That's it! You now have `containerd` installed and running on your Ubuntu system. If you're using this with Kubernetes, remember to configure Kubernetes to use `containerd` as its container runtime.

## Inform K8S we have containerd
If you want `kubeadm init` to use `containerd` as the container runtime for your Kubernetes cluster, you need to specify that during initialization. Here's how you can do it:

1. **Configure containerd to work with Kubernetes**:
   
   By default, `containerd` uses the `cgroupfs` cgroup driver, but Kubernetes recommends using the `systemd` cgroup driver when using `systemd` as the init system. 

   - Ensure you have a `containerd` configuration file. If you don't have one, you can generate a default one with:
     ```bash
     sudo containerd config default > /etc/containerd/config.toml
     ```
     
   - Edit the configuration file `/etc/containerd/config.toml` and ensure the following exists (or add it if missing) under the `[plugins."io.containerd.grpc.v1.cri".containerd.runtimes.runc.options]` section:
     ```
     SystemdCgroup = true
     ```

   - Restart `containerd` after making these changes:
     ```bash
     sudo systemctl restart containerd
     ```

2. **Specify containerd as the Container Runtime in kubeadm**:

   Use the `--cri-socket` flag to specify `containerd`'s socket path:
   
   ```bash
   sudo kubeadm init --pod-network-cidr=10.244.0.0/16 --cri-socket /run/containerd/containerd.sock
   ```

3. **Verify**:

   After initializing your cluster, you can verify that `containerd` is being used by checking the node status:

   ```bash
   kubectl describe node <your-node-name>
   ```

   Under the `Container Runtime Version`, you should see `containerd://<version>`.

That's it! By specifying the `--cri-socket` flag, you've informed `kubeadm` to use `containerd` as the container runtime.
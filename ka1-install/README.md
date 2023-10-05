# Kubernetes Administration

## How to setup a K8S control plane
Setting up a Kubernetes control plane (master) involves installing and configuring several components, including `kube-apiserver`, `etcd`, `kube-scheduler`, `kube-controller-manager`, and others. But don't worry, you don't have to manage each of these individually. Tools like `kubeadm` simplify this process a lot.

Here's a basic step-by-step guide for setting up a Kubernetes master node using `kubeadm`. This guide assumes you're using a machine with Ubuntu, but the process will be similar on other Linux distributions.

1. **Set Up the Machine:**

   Just like adding a node to a cluster, first you need a machine with a compatible operating system. This could be a physical machine, a virtual machine in a local hypervisor, or an instance on a cloud provider.


2 **Install containerd**
  See the [./CONTAINERD-INSTALL.md](continerd-install.md) documentation

3. **Install Kubernetes:**

   Install the Kubernetes components on your machine:

   ```bash
   curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
   echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
   sudo apt-get update
   sudo apt-get install -y kubelet kubeadm kubectl
   sudo apt-mark hold kubelet kubeadm kubectl
   ```

   This installs:

   - `kubelet`, which is the most fundamental building block of a Kubernetes cluster.
   - `kubeadm`, which you can use to bootstrap a Kubernetes cluster.
   - `kubectl`, which you can use to interact with your cluster.

4. **Initialize the Master Node:**

   Now, you can initialize the master node. Use `kubeadm init` to do this:

```bash
sudo kubeadm init --pod-network-cidr=10.244.0.0/16
```

   The `--pod-network-cidr` option sets the range of IP addresses for the pod network. The exact range may depend on the network plugin you're using.

   This command will output a `kubeadm join` command at the end. Make sure to save this command, as you'll need it to join worker nodes to your cluster.

5. **Configure `kubectl`:**

   After initializing the master node, you'll need to configure `kubectl` to interact with your Kubernetes cluster:

   ```bash
   mkdir -p $HOME/.kube
   sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
   sudo chown $(id -u):$(id -g) $HOME/.kube/config
   ```

6. **Install a Pod Network:**

   Finally, you need to install a Pod Network. Kubernetes supports a range of options, including Calico, Flannel, and others. This guide will use Flannel:

   ```bash
   kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
   ```

   Note: The network must be installed before any applications. Also, `kubectl apply -f` may be used to install applications and controllers.

7. **Check the Status of Your Cluster:**

   You can use `kubectl` to check the status of your cluster and see that your master node is running correctly:

   ```bash
   kubectl get nodes
   ```

   If everything is set up correctly, you should see your master node with a status of 'Ready'.

Remember to adjust these instructions based on your specific environment and setup. For example, you may need to open certain ports on your firewall, depending on your security settings. Always check the official Kubernetes documentation for the most accurate and up-to-date information.

## How to add a new node to a cluster?
To add a node to the Kubernetes cluster, you need to configure the new machine to join the cluster. This process varies depending on how your cluster is set up, but if you're using kubeadm, which is a common way to set up Kubernetes clusters, here's the general process:

Install `kubeadm`, `kubelet` and `docker` on the new node. Kubelet is the most fundamental component of Kubernetes and must be installed on any machine that's part of the cluster. Docker is used to run the containers that your applications consist of. The versions of these components should match the ones used in your cluster.
On your master node, run kubeadm token create --print-join-command. This will output a kubeadm join command with the token and the ip of the master node to be used for joining the new node to the cluster.
Copy that command and run it on the new node. This will connect the new node to the cluster.
You can then confirm the new node has joined the cluster by going back to the master node and running kubectl get nodes, which will show a list of all the nodes in the cluster.
Please note, the process may vary depending on the environment, network settings, and how your Kubernetes cluster is set up. Please refer to your specific Kubernetes installation and setup documentation for precise details. The Kubernetes website provides a comprehensive guide for managing a cluster, including adding and removing nodes.

### Here are the detailed steps
To add a new node to a Kubernetes (K8s) cluster, the general process involves setting up a new machine, installing Kubernetes on it, and then joining it to the existing cluster.

Here's a step-by-step guide, assuming you're using `kubeadm` to manage your cluster. This will work for most general-purpose setups, but some cloud providers might have their own specific methods to add nodes.

1. **Set Up the New Machine:**

   First, you need a new machine with a compatible operating system. This could be a physical machine, a virtual machine in a local hypervisor, or an instance on a cloud provider. The process for this will vary greatly depending on where and how you're setting up the machine.

2. **Install Docker:**

   Once your new machine is set up, install Docker. Here's an example of how you might do this on a machine running Ubuntu:
   ```bash
   sudo apt-get update
   sudo apt-get install apt-transport-https ca-certificates curl software-properties-common
   curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
   sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
   sudo apt-get update
   sudo apt-get install docker-ce
   ```
   
3. **Install Kubernetes:**

   You'll also need to install Kubernetes on the new machine. Here's how you might do this:
   ```bash
   curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
   echo "deb http://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
   sudo apt-get update
   sudo apt-get install -y kubelet kubeadm kubectl
   sudo apt-mark hold kubelet kubeadm kubectl
   ```
   
4. **Join the Cluster:**

   With Docker and Kubernetes installed, you can now join the new node to the cluster. To do this, you'll need to run `kubeadm join` on the new machine, along with the appropriate arguments to connect to your existing cluster. The exact command to do this will have been output by `kubeadm init` when you first set up the cluster. It will look something like this:
   ```bash
   kubeadm join --token <token> <master-ip>:<master-port> --discovery-token-ca-cert-hash sha256:<hash>
   ```
   If you don't have the token and the hash anymore, you can get them from your master node:

   - To get the token, you can use `kubeadm token list` on the master node.
   - To get the hash, you can use the following command on the master node: `openssl x509 -pubkey -in /etc/kubernetes/pki/ca.crt | openssl rsa -pubin -outform der 2>/dev/null | openssl dgst -sha256 -hex | sed 's/^.* //'`.

5. **Verify the Node:**

   Back on the master, you can now verify that your new node has been added to the cluster. Use `kubectl get nodes`, and you should see your new machine listed as a node, although it might take a minute or two to reach the 'Ready' status.

Remember to adjust these steps based on your specific setup. For example, you may need to adjust firewall settings or other security configurations, depending on your environment.

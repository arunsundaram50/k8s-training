# DaemonSet for CNI/calico

A Container Network Interface (CNI) plugin is responsible for defining how the different containers within a cluster will communicate with each other. CNI plugins enable various networking features such as overlay networks, network policies, subnet management, etc.

Choosing an appropriate CNI plugin is essential to provide the required network functionality and performance to suit a specific use case. There are several different CNI plugins available such as Flannel, Calico, Weave, etc., each offering its own unique set of features and capabilities.

Here's how you might interact with a CNI plugin using Kubernetes commands:

### Installing a CNI Plugin

Typically, the installation of a CNI plugin is done using a YAML file that describes the plugin's components and configuration. Here's an example command to install the Calico CNI plugin:

```bash
kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml
```

### Initializing a Cluster with kubeadm

When creating a cluster using `kubeadm`, you'll specify a networking addon. The `--pod-network-cidr` option is often used with a particular CNI plugin to specify the range of IP addresses for the pod network.

For example, you might use the following command to create a cluster using Flannel:

```bash
kubeadm init --pod-network-cidr=10.244.0.0/16
```

And then you can apply the Flannel manifest with:

```bash
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/master/Documentation/kube-flannel.yml
```

### Checking the CNI Configuration

You might want to examine the CNI configuration files or the status of the CNI plugin daemon sets. Here are a few commands that can be useful:

- List the CNI plugin configuration files:

  ```bash
  ls /etc/cni/net.d/
  ```

- Check the status of the CNI plugin's daemon sets, for example, Calico:

  ```bash
  kubectl get daemonset -n kube-system -l k8s-app=calico-node
  ```

These are general examples and might vary depending on the specific CNI plugin in use. Always refer to the official documentation for the selected CNI plugin for detailed instructions and considerations.
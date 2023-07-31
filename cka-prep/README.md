# Docker and Kubernetes CKA Sessions

## Sessions for [CKA](https://training.linuxfoundation.org/certification/certified-kubernetes-administrator-cka/)

### Storage 10%
- [Understand storage classes, persistent volumes](#./1.1.md)
- [Understand volume mode, access modes and reclaim policies for volumes](#/show/%24docker-training%2Fcka-prep%2F1.2.md)
- Understand persistent volume claims primitive
- Know how to configure applications with persistent storage

### Troubleshooting 30%
- [Evaluate cluster and node logging](#/show/%24docker-training%2Fcka-prep%2F2.1.md)
- Understand how to monitor applications
- Manage container stdout & stderr logs
- Troubleshoot application failure
- Troubleshoot cluster component failure
- Troubleshoot networking

### Workloads & Scheduling 15%
- Understand deployments and how to perform rolling update and rollbacks
- Use ConfigMaps and Secrets to configure applications
- Know how to scale applications
- [Understand the primitives used to create robust, self-healing, application deployments](#/show/%24docker-training%2Fcka-prep%2F3.4.md)
- Understand how resource limits can affect Pod scheduling
- Awareness of manifest management and common templating tools

### Cluster Architecture, Installation & Configuration 25%
- Manage role based access control (RBAC)
- Use Kubeadm to install a basic cluster
- Manage a highly-available Kubernetes cluster
- Provision underlying infrastructure to deploy a Kubernetes cluster
- Perform a version upgrade on a Kubernetes cluster using Kubeadm
- Implement etcd backup and restore

#### Services & Networking 20%
- Understand host networking configuration on the cluster nodes
- Understand connectivity between Pods
- Understand ClusterIP, NodePort, LoadBalancer service types and endpoints
- Know how to use Ingress controllers and Ingress resources
- Know how to configure and use CoreDNS
- Choose an appropriate container network interface plugin


# Topics needing multi-node cluster(s)

The following topics in our list can benefit from practicing in a multi-node cluster:

### Cluster Architecture, Installation & Configuration 25%
- **Manage a highly-available Kubernetes cluster:** By definition, a highly-available Kubernetes cluster requires multiple nodes. You should understand how to set up and manage a multi-node cluster to ensure availability.

### Services & Networking 20%
- **Understand host networking configuration on the cluster nodes:** This topic requires understanding how networking is configured across different nodes in a cluster.
- **Understand connectivity between Pods:** While you can understand the basics of Pod connectivity within a single node, seeing how Pods communicate across nodes deepens this understanding.
- **Understand ClusterIP, NodePort, LoadBalancer service types and endpoints:** While ClusterIP and NodePort services can be tested with a single node, the LoadBalancer service type is typically used in a multi-node, production environment. To fully understand and practice this service type, a multi-node cluster is more representative of a real-world scenario.

### Troubleshooting 30%
- **Troubleshoot cluster component failure:** Having a multi-node cluster will allow you to simulate node failures and practice resolving them.

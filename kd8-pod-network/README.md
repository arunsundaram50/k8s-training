# k8. DaemonSets

A DaemonSet is a Kubernetes resource that ensures that some or all nodes in the cluster run a copy of a specific Pod. When a new node is added to the cluster, the DaemonSet automatically adds the required Pod to the new node. Conversely, if a node is removed from the cluster, those Pods are garbage collected.

DaemonSets are particularly useful for deploying system-level applications that must be present on every node in the cluster. Some common use cases for DaemonSets include:

- **Logging Agents:** If you want to ensure that logs are collected from all the nodes in your cluster, you might run a logging agent as a DaemonSet.
- **Monitoring Agents:** Similarly, you might use a DaemonSet to deploy a monitoring agent on every node in your cluster.
- **Network Plugins:** Many network plugins, including CNI plugins, are deployed as DaemonSets to ensure that networking functionality is consistently available on every node in the cluster.
- **Storage:** DaemonSets can also be used to manage storage, ensuring that each node has the necessary storage drivers or other related components.

Here's an example command to get information about DaemonSets in the `kube-system` namespace, where system components are usually run:

```bash
kubectl get daemonsets -n kube-system
```

And here's an example of how you might describe a specific DaemonSet, such as the `kube-proxy`:

```bash
kubectl describe daemonset kube-proxy -n kube-system
```

DaemonSets are a powerful way to ensure that certain functionality is consistently available across your cluster, regardless of how nodes are added or removed. They are especially valuable for system-level functionality that needs to be close to the underlying host system.
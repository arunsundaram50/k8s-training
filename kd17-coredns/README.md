# CoreDNS

## Basic anatomy
CoreDNS uses a `Corefile` for configuration. A typical file looks like so:
```
.:53 {
    errors
    health
    kubernetes cluster.local in-addr.arpa ip6.arpa {
       pods insecure
       fallthrough in-addr.arpa ip6.arpa
    }
    forward . /etc/resolv.conf
    cache 30
}
```
**In the above configuration**
The server listens on port 53.
- It serves the cluster.local domain, which is common for Kubernetes clusters.
- Forwards non-local domains based on the /etc/resolv.conf file.
- Caches responses for 30 seconds.


Once CoreDNS is running within your Kubernetes cluster, any pods with DNS policy set to "Default" will use CoreDNS for DNS queries. This means you can access services by their service name, e.g., myservice.mynamespace.svc.cluster.local.

Outside of Kubernetes, you can use CoreDNS for a variety of purposes, such as DNS load balancing, forwarding, etc., depending on the plugins you enable in the Corefile.

**Plugins**
One of the strong features of CoreDNS is its extensibility through plugins. Each stanza in the Corefile effectively enables and configures a plugin. There's a rich ecosystem of plugins, allowing CoreDNS to be used in diverse ways beyond just service discovery in Kubernetes.

To further leverage CoreDNS, you'll want to review available plugins and see which ones align with your needs. Each plugin comes with its configuration options and directives, so you'll need to refer to the documentation for each when modifying your Corefile.


To check if CoreDNS is running:
```shell
kubectl get pods -n kube-system -l k8s-app=kube-dns
```

Get the current config map:
```shell
kubectl get cm coredns -n kube-system -o yaml
```

Edit the configuration:
```shell
kubectl edit cm coredns -n kube-system
```

restart the CoreDNS pods for changes to take effect:
```shell
kubectl rollout restart deployment coredns -n kube-system
```

## How to add custom stub domains
A common customization made to CoreDNS in Kubernetes environments is to add custom stub domains. This allows you to forward DNS queries for specific domains to specific nameservers, while other queries are forwarded to the default nameserver. This is useful when you have services outside the Kubernetes cluster with their own DNS, but you want in-cluster services to be able to resolve them.

For instance, let's say you have an internal company DNS server at `10.150.0.10` that serves DNS records for the domain `company.internal`. You want your applications inside the Kubernetes cluster to resolve names like `service.company.internal` using that nameserver.

Here's how you can achieve this by customizing the CoreDNS `Corefile`:

### 1. Modify the CoreDNS ConfigMap:

First, get the current configuration:

```bash
kubectl get cm coredns -n kube-system -o yaml > coredns-cm.yaml
```

Open `coredns-cm.yaml` in an editor, and modify the `Corefile` section. 

### 2. Add the Stub Domain Configuration:

```text
company.internal:53 {
    errors
    cache 30
    forward . 10.150.0.10
}
```

This configuration ensures that any DNS query ending in `company.internal` will be forwarded to the DNS server at `10.150.0.10`.

### 3. Apply the Configuration:

After modifying `coredns-cm.yaml`, apply the changes:

```bash
kubectl apply -f coredns-cm.yaml
```

### 4. Restart CoreDNS:

To ensure the changes are picked up, restart the CoreDNS pods:

```bash
kubectl rollout restart deployment coredns -n kube-system
```

Now, any service or pod within the cluster trying to resolve a `.company.internal` domain will query the `10.150.0.10` DNS server.

This kind of stub domain customization is quite useful in enterprise environments with internal domains or in situations where you're integrating Kubernetes into existing infrastructure.

## Stub Domain and Subdomains
The terms "stub domain" and "subdomain" refer to different concepts in the realm of DNS and domain naming:

1. **Stub Domain**:
   - A stub domain in the DNS context refers to a setup where DNS queries for a specific domain (and its subdomains) are forwarded to a particular nameserver, rather than resolving them using the default nameservers.
   - The "stub" aspect indicates that the domain's authoritative server is somewhere else, and the current system (like CoreDNS in our context) should just act as a "stub", forwarding requests for that domain to the designated nameserver.
   - It doesn't host the full zone information but knows where to forward queries for that zone.

2. **Subdomain**:
   - A subdomain is a domain that is a part of a larger domain in the DNS hierarchy. For example, in `blog.example.com`, `blog` is a subdomain of `example.com`.
   - Subdomains are used to organize and partition the namespace of the larger domain.

When configuring DNS settings like in CoreDNS, the term "stub domain" is used to denote domains for which the DNS service will forward queries to a designated nameserver. This doesn't necessarily have anything to do with whether or not the domain is technically a "subdomain" of another domain.

In essence, any domain (whether it's a top-level domain, second-level domain, or deeper subdomain) can be treated as a stub domain if we decide to forward queries for it to a specific nameserver.

## Blackholing: blocking or redirecting certian domains
Another common and useful customization for CoreDNS is **blocking or redirecting certain domain names**, sometimes known as "DNS blackholing." This can be useful for security reasons, to prevent access to malicious websites, or to redirect traffic within your infrastructure.

### **DNS Blackholing with CoreDNS:**

1. **Setup a Blocked Zone in the Corefile**:

    If you want to block access to certain domain names, you can set up a zone for that domain and then serve a "null" IP or a custom IP. Let's say you want to block access to `malicious.com`:

    ```text
    malicious.com {
        respond 0.0.0.0
    }
    ```

    This configuration will cause CoreDNS to respond with the IP address `0.0.0.0` for any DNS query to `malicious.com`, effectively blocking access to it. You can also use the `127.0.0.1` IP address if you want queries to loop back to the local machine.

2. **Customize the Response for Multiple Domains**:

    If you have a list of domains to block or redirect, you can do so in one stanza:

    ```text
    malicious.com malicious.net badsite.org {
        respond 0.0.0.0
    }
    ```

3. **Redirecting Domains**:

    If you want to redirect traffic for a specific domain to another IP (e.g., redirecting internal requests to an internal server), you can use the `respond` directive:

    ```text
    internal-service.mycompany.com {
        respond 10.0.0.5
    }
    ```

    In this case, any DNS queries for `internal-service.mycompany.com` will be resolved to the IP `10.0.0.5`.

4. **Apply the Configuration**:

    After modifying the CoreDNS `Corefile` to include these customizations, save the changes and apply them by updating the config map and restarting the CoreDNS pods:

    ```bash
    kubectl apply -f coredns-cm.yaml
    kubectl rollout restart deployment coredns -n kube-system
    ```

This kind of customization is straightforward but can be highly effective, especially in controlled environments where you want to ensure certain domains are inaccessible or when you want to direct traffic in specific ways based on the requested domain name.
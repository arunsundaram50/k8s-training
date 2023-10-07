To use the `AutoSecret` controller in a container, you'd follow these steps:

### 1. Containerize the Controller:

First, you'd need to create a `Dockerfile` for the controller.

**Dockerfile**:

```Dockerfile
FROM golang:1.17 AS builder

WORKDIR /workspace
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o autosecret-controller .

FROM alpine:3.14

COPY --from=builder /workspace/autosecret-controller /autosecret-controller

ENTRYPOINT ["/autosecret-controller"]
```

Build the Docker image:

```bash
docker build -t autosecret-controller:latest .
```

### 2. Push the Docker Image:

Push the image to a container registry (Docker Hub, Google Container Registry, etc.):

```bash
docker tag autosecret-controller:latest yourusername/autosecret-controller:latest
docker push yourusername/autosecret-controller:latest
```

### 3. Kubernetes Deployment:

You'll need to deploy the controller inside the Kubernetes cluster. It will need permissions to watch `AutoSecret` CRDs and create `Secret` resources.

**RBAC Permissions**:

First, set up the necessary RBAC permissions.

`rbac.yaml`:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: autosecret-controller
rules:
- apiGroups:
  - custom.k8s.io
  resources:
  - autosecrets
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - secrets
  verbs:
  - get
  - list
  - create

---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: autosecret-controller
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: autosecret-controller
subjects:
- kind: ServiceAccount
  name: autosecret-controller
  namespace: default
```

Apply the permissions:

```bash
kubectl apply -f rbac.yaml
```

**Deployment**:

Next, deploy the controller.

`deployment.yaml`:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: autosecret-controller
spec:
  replicas: 1
  selector:
    matchLabels:
      app: autosecret-controller
  template:
    metadata:
      labels:
        app: autosecret-controller
    spec:
      serviceAccountName: autosecret-controller
      containers:
      - name: autosecret-controller
        image: yourusername/autosecret-controller:latest
```

Apply the deployment:

```bash
kubectl apply -f deployment.yaml
```

### 4. Monitor and Use:

With the controller running in your cluster, you can now create `AutoSecret` resources, and the controller will automatically create associated `Secret` resources with random values.

**Monitor the controller**:

```bash
kubectl logs -l app=autosecret-controller -f
```

Now, whenever you create an `AutoSecret` custom resource, the controller should detect it, generate a random secret, and create a corresponding Kubernetes `Secret` resource. The controller logs will provide insights into its operations.
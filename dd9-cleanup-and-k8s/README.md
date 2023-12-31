# 9. The need for k8s...
... based on our experience with `docker-compose`

Docker provides restart policies which can be used to control whether our containers start automatically when they exit, or when Docker restarts. These are:
- no: This is the default policy. No action will be taken when a container stops.
- always: This will always restart the container regardless of the exit status. The container will also be started on daemon startup, regardless of the current state.
- on-failure: This will restart the container if the container exits with a non-zero exit status. We can optionally specify a maximum number of times Docker will try to restart the container.
- unless-stopped: This will always restart the container unless it was manually stopped.

```bash
docker run -d --restart always my_service
```
In the above example, my_service will always restart if it crashes or if the docker daemon restarts.

This can be specified in the docker componse YAML file
```yaml
version: '3'
services:
  my_service:
    image: my_service
    restart: always
```

Note: `restart: always` also makes sure the container starts after a host machine restart as well.


### Kubernetes
If we're running our containers on a Kubernetes platform, the Kubernetes platform ensures our pods (which contain our Docker containers) are always running. If a pod crashes, Kubernetes automatically starts a new pod to replace it.
```
In the example above, the host computer's `./my-data` is visible as `/data` by the container. The files written to (or read by) the container under the directory `/data` outlives the life-span of the container.


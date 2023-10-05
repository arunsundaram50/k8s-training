# From testing to deployment

## Docker part: Create images and test them seperately
  ```bash
  docker image build -t arunsundaramco70/hello .
  docker container run --rm -p 8001:8001 arunsundaramco70/hello 
  docker image build -t arunsundaramco70/upper .
  docker container run --rm -p 8002:8001 arunsundaramco70/upper
  ```
Now, <http://localhost:8001/hello/peter> should work.

Since the upper container (the mircoservice inside the container) has a dependency on hello microservice we can't test upper <http://localhost:8002/upper/peter> by itself even though both are running.
This is because we haven't enabled container<->container communication. This is where docker-compose.yaml comes in.
  ```bash
  docker-compose up
  ```
Now you can hit http://localhost:8002/upper/peter


## Deploying using K8S
```bash
k apply -f pod.yaml
k apply -f service.yaml
```

http://localhost:30089/upper/peter

Notice how the service fails if `hello` service wasn't running.

This is where we use `deployment.yaml` instead of managing `pod.yaml` and `service.yaml` seperately as `deployment.yaml` contains the all the pods and services.

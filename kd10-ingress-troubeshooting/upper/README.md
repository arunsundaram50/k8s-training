# Building and testing `upper`
```bash
docker image build -t arunsundaramco70/upper .
docker image push arunsundaramco70/upper
k apply -f pod.yaml
k apply -f service.yaml
```

http://localhost:30089/upper/peter

Notice how the service fails if `hello` service wasn't running.
This is where we use `deployment.yaml` instead of managing `pod.yaml` and `service.yaml` seperately as `deployment.yaml` contains the all the pods and services.

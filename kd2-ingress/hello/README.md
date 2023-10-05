# Building and testing `hello`
```bash
docker image build -t arunsundaramco70/hello .
docker image push arunsundaramco70/hello 
k apply -f pod.yaml
k apply -f service.yaml
```
http://localhost:30088/hello/peter

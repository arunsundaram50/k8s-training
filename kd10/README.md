# Using Ingress

```
http://hello.my-server.com/ -> hello-service
http://upper.my-server.com/ -> upper-service
```
To troubleshoot the service
```
kubectl port-forward svc/upper-service 8080:8089
```

```
curl -v --resolve upper.my-server.com:<port>:<ingress-controller-ip> http://upper.my-server.com/2eeqewqw
```
Using /etc/hosts we resolve upper.my-server.com to localhost.
So a request to http://hello.my-server.com:8088/hello/asddas goes to localhost:80.
The Ingress Controller listens to localhost:8088 and localhost:8089

```bash
Connection http://upper.my-server.com => localhost 
GET /upper/asddas
http://hello.my-server.com:8088/hello/asddas
http://upper.my-server.com:8089/upper/asddas
```
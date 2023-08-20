# k10. Using Ingress

Using `ingress.yaml` the we want to route the URLs to the respective services
```
http://hello.my-server.com/ -> hello-service
http://upper.my-server.com/ -> upper-service
```
Edit /etc/hosts to resolve hello.my-server.com and upper.my-server.com to localhost
So a request to http://hello.my-server.com:8088/hello/asddas goes to localhost:80.

## Troubleshooting
If this didn't work, one can troubleshoot it in steps.
First make sure the service is accessible by temporarily port forwarding
```
kubectl port-forward svc/upper-service 8080:8089
```

The Ingress Controller seems to listen to localhost:8088 and localhost:8089

But
```
http://hello.my-server.com/hello/asddas
http://upper.my-server.com/upper/asddas
```
should have worked.

But, only 
```
http://hello.my-server.com:8088/hello/asddas
http://upper.my-server.com:8089/upper/asddas
```
did.

See the ingress-controller logs:
```
kubectl logs -n ingress-controller -l app.kubernetes.io/name=ingress-nginx -f
```

The last line indicated an error: 
```
W0820 13:07:20.108016       7 controller.go:328] ignoring ingress my-ingress in default based on annotation : ingress does not contain a valid IngressClass
```

Upon searching the net for this error, it seems `ingress.yaml` needs a valid ingress.class, so added it
```
 kubernetes.io/ingress.class: "nginx"
```
And now, the ports work without a port#

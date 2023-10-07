Setup the controller
```
go mod init helloworldcontroller
go work use .
go get k8s.io/client-go
go run main.go
```
Create the CRD helloworld resource
```
kubectl apply -f helloworld.yaml
```


Internally
- Service:
  - K8S service 
    - LoadBalancer
    - NodePort
  - Microservice (the RestfulAPIs we created using Python)
    - upper
    - hello
  - Ingress

What we want:
  - http://my-server.com/upper
  to call our microservice upper-service

Locally, we can test upper-service like so:
  - http://localhost:8888 
  and that is because we have a LB-service inside of K8s 


### Query Request
GET http://my-server.com/employee/10
```json
{
  "id": 10,
  "name": "xyz"
}
```

### Create Request
POST http://my-server.com/employee
```json
{
  "name": "abc"
}
```
Response
200 OK
```json
{
  "id": 11,
  "name": "abc"
}
```

### Update Request
PUT http://my-server.com/employee/11
```json
{
  "name": "abcABC"
}
```
Response
200 OK
```json
{
  "id": 11,
  "name": "abcABC"
}
```

#### Delete Request
DELETE http://my-server.com/employee/11


Ontology Examples
```
POST http://amazon.com/cart/items/xyz12321
GET http://amazon.com/cart/items


http://123.231.13.34:93434/cart-service 
  => http://amazon.com/cart/items
  => http://cart.amazon.com/items
```


http://my-server.com/upper/hello -> upper-service:8888 [SERVICE]-> 10.1.0.163:8001 [POD]
                               INGRESS                       SERVICE

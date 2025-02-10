# 3. Ephemeral state

### Build the microservice image
```
docker image build -t arunsundaramco70/dd3 .
```

### Run the microservice image
```
docker container run  -p 8001:8001 arunsundaramco70/dd3
```

### Access the microservice from your browser
```
http://localhost:8001/hello2
```

### Locate the container ID
```
docker container list
```

### Attach
```
docker exec -it <CONTAINER_ID> bash
```

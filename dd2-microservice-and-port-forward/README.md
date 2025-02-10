# 2. Building a Microservice 

### Build the microservice image
`docker image build -t arunsundaramco70/dd2 .`

### Run the microservice image
Try #1 (does not expose port)
```
docker container run arunsundaramco70/dd2
```
Try #2
```
docker container run -p 8001:8001 arunsundaramco70/dd2
```

### Access the microservice from your browser
`http://localhost:8001/hello`


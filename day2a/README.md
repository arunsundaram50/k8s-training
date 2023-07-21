# A "hello, world" Microservice with 'ephameral' request count

### Build the microservice image
`docker image build -t arunsundaramco70/day2a .`

### Run the microservice image
`docker container run  -p 8001:8001 arunsundaramco70/day2a`

### Access the microservice from your browser
`http://localhost:8001/hello2`

### Attach
`docker exec -it <CONTAINER_ID> bash`

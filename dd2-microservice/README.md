# 2. Building a Microservice 

### Build the microservice image
`docker image build -t arunsundaramco70/day2 .`

### Run the microservice image
test
`docker container run arunsundaramco70/day2 `
versus
`docker container run -p 8001:8001 arunsundaramco70/day2`

### Access the microservice from your browser
`http://localhost:8001/hello`


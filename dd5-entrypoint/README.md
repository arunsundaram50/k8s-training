# 5. Docker command summary

## Other Dockerfile primitives/commands
- ARG
- [.] CMD: array preferred
- [.] ENV
- [.] ENTRYPOINT
- HEALTHCHECK
- LABEL
- ONBUILD INSTRUCTION
- STOPSIGNAL signal
- USER daemon
- VOLUME [ "/data" ]
- WORKDIR /the/workdir/path


## Image commands: docker image ....
- [.] build
- [.] ls
- inspect
- export
- [.] search
- [.] spush, pull
- [.] rmi


## Container commands: docker container ...
- [.] run ([.]-d, [.]-it, [.]--rm, [.]-p, [.]-e)
- [.] stop
- [.] ps
- attach, exec
- [.] logs

## System commands: docker system
- [.] prune
- [.] custom script using awk to cleanup


## Run a mongo server and client
Run the server
```
docker run -p 27017:27017 -d -e MONGO_INITDB_ROOT_USERNAME=admin -e MONGO_INITDB_ROOT_PASSWORD=passowrd --name mongodb --network mongo-network mongo
```

Run the client
```
docker run -d \
 -p 8081:8081
 -e ME_CONFIG_MONGODB_ADMINUSERNAME=admin \
 -e ME_CONFIG_MONGODB_ADMINPASSWORD=password \
 -e ME_CONFIG_MONGODB_SERVER=mongodb \
 --name mongo-express \
 --network mongo-network \
mongo-express
```

## When the OS prepares a executable program for running, it sets up:
Inputs: environment, arguments, stdin
Outputs: stdout, stderr


# 8. Docker Volumes

## Mounting bind volumes
```
  app1:
    build: ./app1
    volumes:
      ./my-data:/data
```
In the example above, the host machine's ./my-data folder is 'seen' by the container as /data. Any file the container writes in the /data directory outlives the life of the container.

## Mounting docker managed volumes
```
version: '3'
services:
  web:
    image: nginx:latest
    volumes:
      - my-volume:/usr/share/nginx/html

volumes:
  my-volume:
```

In the above example:

A named volume called `my-volume` is defined under the volumes key at the bottom of the file.
The web service (using the nginx image) mounts this volume at the path `/usr/share/nginx/html`.
This setup means that the contents of `/usr/share/nginx/html` inside the nginx container will be stored in the my-volume volume on the host. Any other containers that mount this same volume will have access to the same data.

You can bring up the services defined in this file by running docker-compose up in the directory where the docker-compose.yml file is located. You can stop and remove the services (without removing the volume) by running docker-compose down.

If you want to remove the volume as well when bringing down the services, you can `run docker-compose down -v`. Note that this will delete all the data in the volume.

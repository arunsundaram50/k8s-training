# 7. Docker Compose

## Why we need `docker_compose`?
1. Individual images have to be built
2. Running the images have to be properly coordinated
3. Monitoring the containers is a huge task
4. Containers are not able to `easily` talk to each other


## Common docker-compose commands
```
docker-compose build
docker-compose up
docker-compose up -d

docker-compose down
docker-compose stop
docker-compose pause

docker-compose logs
docker-compose logs -f
docker-compose up -v 
```

Typically up also does a build if did not find the container images.

The `pause`, `stop`, and `down` commands in Docker Compose have different effects on containers, and here's an overview of what each one does:

1. **`pause`**:
   - The `docker-compose pause` command pauses all the running containers defined in the `docker-compose.yml` file.
   - A paused container's process execution is literally "paused," freezing its state.
   - Network traffic is also suspended.
   - The container remains running but doesn't consume CPU resources.
   - You can resume a paused container with `docker-compose unpause`.

2. **`stop`**:
   - The `docker-compose stop` command stops the running containers.
   - This stops the processes inside the containers but does not remove the containers themselves.
   - You can start the containers again with `docker-compose start`.
   - Stopping a container also means that any data that is not stored in a volume or bind mount will be lost when the container is removed later on.

3. **`down`**:
   - The `docker-compose down` command stops the running containers and also removes them.
   - It also removes any networks and volumes defined as "external: false" in your `docker-compose.yml` file.
   - You can add the `--volumes` or `-v` option to remove all the volumes defined in the `docker-compose.yml` file as well.
   - Using `down` means that you'll need to create and start the containers again with `docker-compose up`.

In summary, `pause` freezes the containers without stopping them, `stop` halts the containers without removing them, and `down` stops and removes the containers (and can also remove networks and volumes).


## Mounting volumes
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

A named volume called my-volume is defined under the volumes key at the bottom of the file.
The web service (using the nginx image) mounts this volume at the path /usr/share/nginx/html.
This setup means that the contents of /usr/share/nginx/html inside the nginx container will be stored in the my-volume volume on the host. Any other containers that mount this same volume will have access to the same data.

You can bring up the services defined in this file by running docker-compose up in the directory where the docker-compose.yml file is located. You can stop and remove the services (without removing the volume) by running docker-compose down.

If you want to remove the volume as well when bringing down the services, you can run docker-compose down -v. Note that this will delete all the data in the volume.


## Cleaning up
Docker containers, images, and other resources can accumulate over time and eat up your system's disk space. Here's how you can find and remove old containers that are no longer needed.

### Find Old Containers

You can list all containers, including the stopped ones, by running:

```bash
docker ps -a
```

If you want to list only the containers that have exited, you can use:

```bash
docker ps -a -f status=exited
```

### Remove Old Containers

Once you've identified the containers you'd like to remove, you can remove them by their container ID or name.

To remove a single container:

```bash
docker rm <CONTAINER_ID_OR_NAME>
```

You can remove all stopped containers at once using:

```bash
docker container prune
```

Or, if you want to remove all containers that are not currently running (this includes the ones with a status of `created`, `exited`, etc.):

```bash
docker rm $(docker ps -a -q)
```

### Remove Old Images, Volumes, and Networks

In addition to removing old containers, you might also want to clean up old images, volumes, and networks. You can use the following commands:

Remove unused images:

```bash
docker image prune
```

Remove all unused images, not just the dangling ones:

```bash
docker image prune -a
```

Remove unused volumes:

```bash
docker volume prune
```

Remove unused networks:

```bash
docker network prune
```

### All-In-One Cleanup Command

If you want to remove all unused data (containers, volumes, networks, and images), you can use:

```bash
docker system prune -a
```

This command will free up space by removing all unused objects. Be cautious when using it, as it will remove a lot of data that you might want to keep.

Remember that these commands may require administrative permissions, so you might need to prepend them with `sudo` if you encounter permission issues. Also, always be sure of what you are removing, especially in a production environment, to avoid unexpected issues.
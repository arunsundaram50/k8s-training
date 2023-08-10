# 7. Docker Compose

## Why we need `docker_compose`?
1. Individual images have to be built one by one manually using `docker image build ...`
2. Running the containers have to be properly coordinated, and dependencies managed
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

Note: typically `docker-compose up` also does a build if the images needed aren't available. However, it seems like the images won't be updated if there are changes to the image's source materials.

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

### Reading data off of the volume
If you've got data in a Docker volume that was created using Docker Compose, and you want to export it to your host system, here's a general guideline on how you can do this.

Suppose you have a volume mounted inside your container at `/data`. Here's how you could export the contents of that volume to a tar file on your host system:

1. **Identify the Container using the Volume**: First, you'll need to figure out the name or ID of the container that's using the volume. You can do this using the `docker ps` command or by checking your `docker-compose.yml` file.

2. **Open a Shell Inside the Container**: Execute a shell inside the container:

   ```bash
   docker exec -it <container_name_or_id> /bin/bash
   ```

   Replace `<container_name_or_id>` with the actual name or ID of the container.

3. **Create a Tar File of the Data Inside the Container**: While inside the container's shell, you can create a tar file of the `/data` directory:

   ```bash
   tar cvf /data.tar /data
   ```

4. **Copy the Tar File to the Host**: Exit the container's shell, and then use the `docker cp` command to copy the tar file to your host system:

   ```bash
   docker cp <container_name_or_id>:/data.tar /path/to/host/directory
   ```

   Replace `/path/to/host/directory` with the actual path on your host system where you want to store the tar file.

5. **Cleanup**: If you'd like, you can remove the tar file from inside the container using:

   ```bash
   docker exec -it <container_name_or_id> rm /data.tar
   ```

6. **Extract the Tar File on the Host**: Finally, you can extract the tar file on your host system using:

   ```bash
   tar xvf /path/to/host/directory/data.tar -C /path/to/extract/
   ```

That's it! The contents of the mounted volume are now on your host system.

Remember to replace the paths and container names with those that are relevant to your specific setup. Make sure to have appropriate permissions and be aware that this method could potentially overwrite existing files in the target directory on your host system, so take necessary precautions!


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
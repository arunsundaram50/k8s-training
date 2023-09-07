## Here are some Docker cleanup tasks

### Remove Stopped Containers
```bash
docker container prune
```

### Remove Unused Images
```bash
docker image prune -a
```
The -a flag will remove all images not used by *any* container. Omitting this flag will only remove dangling images (images that are not tagged and not used by any container).


### Remove Unused Volumes
```bash
docker volume prune
```

### Remove Unused Networks
```bash
docker network prune
```

### All Together (not suitable for use in Production)
One can combine all of these into one command that will remove stopped containers, all unused images, all unused volumes, and all unused networks:

```bash
docker system prune --volumes
```
The --volumes flag also removes volumes, in addition to the other items that are pruned.

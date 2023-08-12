## Passing environment variables to the container

### Define environment variable inside `Dockerfile`
Environment variables can  be defined in the `Dockerfile` like so:
```Dockerfile
# Use an official Python runtime as a parent image
FROM python:3.8-slim-buster

# Set environment variables
ENV ORACLE_USER_NAME scott
```

### Override environment variable that's inside `Dockerfile` at runtime
You can also pass env variables like so. This will override what's specified in the Dockerfile
```bash
docker container run -e ORACLE_USER_NAME=scott -e ORACLE_PASSWORD=tiger imagename
```
Given this, the code inside the container can read it like so:
```python3
password = os.environ.get('ORACLE_PASSWORD')
```

### Override environment variable that's inside `Dockerfile` at buildtime
This is strictly not just using `ENV`, but also using `ARG` instruction.

Given a Dockerfile like so:
```Dockerfile
FROM alpine
ARG MY_VARIABLE
ENV MY_VARIABLE_IN_IMAGE=$MY_VARIABLE
CMD echo $MY_VARIABLE_IN_IMAGE
```

One can issue a build command like so:
```bash
docker build --build-arg MY_VARIABLE=$MY_VARIABLE -t my_image .
```

And run the container like so:
```bash
docker run my_image  # Outputs the value of MY_VARIABLE from the host
```
The value of MY_VARIABLE in the host environment at the time you build the image will be baked into the image and will be used as the value of the MY_VARIABLE_IN_IMAGE environment variable inside the container.

This approach makes the value of the environment variable a permanent part of the image, so it will be the same every time you run a container from that image, even if you change the value of MY_VARIABLE in the host environment after building the image. If you want to set a different value for the environment variable, you would need to build a new image with a new value for the build argument.

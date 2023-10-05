
## Capturing output of `RUN` commands!
One would expect the command
```
docker build -t your-image . 2>&1 | tee build_output.txt
```
would capture the output of `RUN` instructions inside of the `Dockerfile`.

For example, to examine the OS used by the image, you might have something like `RUN uname -a` and expect to see its output each time you run `docker build`. 

However, the output of `uname -a` will likely be captured and cached after the first time running it and you will only see the word "cached" or similar and not see the output.

- One way to solve it is to use the `--no-cache` flag.

- Another crude way to see the output is to temporarily have `CMD uname -a`. 

- Using CACHEBURST (the argument name CACHEBUST is arbitrary)
```
ARG CACHEBUST=1
RUN apt-get update && apt-get install -y your-package
```
and run
```
docker build --build-arg CACHEBUST=$(date +%s) -t your-image .
```

These are not optimal as there are better ways to troubleshoot a failing build, as we will see later.

## Finding a base image to use.
- Search `https://hub.docker.com` for the image
- Run search locally like so: `docker search python`

## You can have only one `CMD`, but many `RUN` instructions.
This Docker's design philosophy. Each image is meant to have one `CMD`.

## Other tips
- Combination of `Dockerfile` instructions and `build-arg`uments are used for caching
- Changing arguments to the `CMD` instruction creates new image (SHA)

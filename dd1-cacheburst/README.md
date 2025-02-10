# Some build GOTCHAs

## Capturing output of `RUN` commands!
One would expect the command
```
docker build -t dd1 -f Dockerfile-with-no-cachebust . 2>&1 | tee build_output.txt
```
would capture the output of `RUN` instructions inside of the `Dockerfile`.

For example, to examine the OS used by the image, you might have something like `RUN uname -a` and expect to see its output each time you run `docker build`. 

However, the output of `uname -a` will likely be captured and cached after the first time running it and you will only see the word "cached" or similar and not see the output.

- One way to solve it is to use the `--no-cache` flag. But, this makes image building inefficient.

- Another crude way to see the output is to temporarily have `CMD uname -a`.  But, this takes manual (one-time exception) work.

- Use CACHEBUST (the argument name CACHEBUST is arbitrary). This is cleaner.

Just add the line towards the end of the `Dockerfile`
```
ARG CACHEBUST=1
```
and run (see the GOTCHA-2 below)
```
docker build --build-arg CACHEBUST=$(date +%s) -t dd1 . | cat 
```

### GOTCHA-2
The reason we need a `cat` or a `tee` is because the output of RUN commands inside Dockerfile are cached and often "not seen" in the terminal.
When piped out to other commands, `docker` command does not buffer the output. Another way around this is to change how progress is reported.
```
docker buildx build --progress=plain --build-arg CACHEBUST=$(date +%s) -t dd1 .
```

All these are not optimal as there are better ways to troubleshoot a failing build, as we will see later.

# Finding a base image to use.
- Search `https://hub.docker.com` for the image
- Run search locally like so: `docker search python`

## You can have only one `CMD`, but many `RUN` instructions.
This is one of Docker's design philosophy. Each image is meant to have ONLY one `CMD`.

## Other tips
- Combination of `Dockerfile` instructions and `build-arg`uments are used for caching
- Changing arguments to the `CMD` instruction creates new image (SHA)

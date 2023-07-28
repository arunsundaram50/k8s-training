# 4. Saving States

As we saw in dd2 & dd3 session, any states maintained by `main.py` in a variable or in a file within the container was ephemeral.

In order to save the state, we have to move the `count.txt` file out of the container and into the host file system, for example.

## Not using external volume
First let's run the container without an external volume mounted (which creates an instance of `count.txt` for each container)
```
docker image build -t main .
docker container run main
```
Kill the container and run it again
```
docker container run main
```
only to see another `count.txt` was created with the count starting from 0 again!

## Using an external (mapped) volume
Now let's run the container with container's `/data` mapped to an external directory in the host OS. The first container would create `count.txt` in this host file system
```
docker image build -t main .
docker container run main
```
Kill the container and run it again. The second run (which creates a new container) uses the `count.txt` that was created by the first container as `count.txt` seen by both are the same "external" file.
```
docker container run main
```
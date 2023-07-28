# 3. Saving States

As we saw in dd2 & dd3 session, any states maintained by `main.py` in a variable or in a file within the container was ephemeral.

In order to save the state, we have to move the `count.txt` file out of the container and into the host file system, for example.

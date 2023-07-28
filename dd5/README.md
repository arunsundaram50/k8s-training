# 5. Docker command summary

## Image commands
- build
- ls
- inspect
- export
- search
- push, pull
- rmi

## Container commands
- run (-d, -it, --rm, -p)
- stop
- ps
- attach, exec
- logs
- rm

## System commands
- prune
- custom script using sed, awk to cleanup

## Other Dockerfile commands
- ARG
- CMD: array preferred
- ENV
- ENTRYPOINT
- HEALTHCHECK
- LABEL
- ONBUILD INSTRUCTION
- STOPSIGNAL signal
- USER daemon
- VOLUME [ "/data" ]
- WORKDIR /the/workdir/path

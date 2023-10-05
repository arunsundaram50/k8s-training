
# Python Docker Images
## Running as a script
### All happens at runtime
Python script (example: hello.py) [1KB] -->  Python Interpreter [200MB] --> Output


## Running as a compiled binary (.exe, .bin, etc.)
### Happens at compile time (meaning does not have to happen inside the Docker container)
Python script (example: hello.py) [1KB] --> Python compiler for Linux [~1GB] --> hello.bin [.5GB]
### Happens at runtime, meaning all the components it needs should be inside the Docker image
```
[alpine] hello.bin [.5GB] -> Output
```
A large part of Python will be inside of hello.bin


# Go Docker Images
Go script (example: hello.go) [1KB] -> Go Compiler -> hello.bin [5MB]
[alpine] hello.bin [5MB] -> Output

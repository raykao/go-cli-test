## Dev Environment

I do not want to install go on my local system.  Instead I prefer running inside docker and mounting the project src directory into it as a volume.  I can still write within VS Code and use the docker interactive terminal to run commands as needed.

```
# Build docker dev container with included dockerfile
# while this is a single stage dockerfile at the moment I still prefer to target the specific docker multistage build defined in the docker file (dev)

docker build --target dev . -t go

docker run -it -v ${PWD}:/src go sh
```
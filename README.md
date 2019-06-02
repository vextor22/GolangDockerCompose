# GolangDockerCompose
Writing a dockerfile for golang development isn't straightforward. To maintain quick compiles, create a multi-stage Docker image. Build can cache 'go get' and ship the executable to an Ubuntu image.

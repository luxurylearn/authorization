# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from golang base image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/yuranye/auth-service

# Copy everything from the current directory to the WORKDIR in container
COPY ../../cmd/auth .

# Download all dependencies
# https://stackoverflow.com/questions/28031603/what-do-three-dots-mean-in-go-command-line-invocations
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["auth-service"]



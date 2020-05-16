# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest

# Add Maintainer Info
LABEL maintainer="<hongli@redhat.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

COPY ./bin/server .

# Expose port 8080 to the outside world
EXPOSE 50080

# Command to run the executable
CMD ["./server"]

# step1: build binary
FROM golang:alpine

# Add Maintainer Info
LABEL maintainer="<hongli@redhat.com>"

WORKDIR /app

COPY ./bin/server .

# Expose port 50080
EXPOSE 50080

# Run the executable
ENTRYPOINT ["./server"]

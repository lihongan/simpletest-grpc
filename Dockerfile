# step1: build binary
FROM golang:alpine

# Add Maintainer Info
LABEL maintainer="<hongli@redhat.com>"

WORKDIR /app

COPY ./cert ./cert
COPY ./bin/secure-server ./bin/secure-client ./bin/

# Expose port 50443
EXPOSE 50443

# Run the executable
ENTRYPOINT ["./bin/secure-server"]

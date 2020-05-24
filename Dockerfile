# step1: build binary
FROM golang:alpine

# Add Maintainer Info
LABEL maintainer="<hongli@redhat.com>"

WORKDIR /app

COPY ./bin/secure-server ./bin/secure-client ./

# Expose port 50443
EXPOSE 50443

# Run the executable
ENTRYPOINT ["./secure-server"]

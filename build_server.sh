GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -v -ldflags '-extldflags "-static"' -o bin/server ./server

### Build todo app.
FROM golang:1.11.0-stretch

# Go build 
COPY . /go/src/todo
WORKDIR /go/src/todo/cmd/todo
RUN go build


## Run our API server
FROM debian:stretch-slim

# Install ca-certificates.
RUN apt-get update && apt-get install -y ca-certificates && update-ca-certificates

# Executables
COPY --from=0 /go/src/todo/cmd/todo/todo /usr/bin/todo

COPY --from=0 /go/src/todo/configs/ssl/key /usr/ssl/key
COPY --from=0 /go/src/todo/configs/ssl/cert /usr/ssl/cert

# Run
EXPOSE 80 443 8090
ENTRYPOINT ["todo"]
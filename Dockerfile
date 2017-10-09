# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.9


ADD . /go/src/gawkbox-assignment

WORKDIR /go/src/gawkbox-assignment

RUN go build

# Document that the service listens on port 8080.
EXPOSE 8080
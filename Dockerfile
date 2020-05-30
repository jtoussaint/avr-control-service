# Start from an Alpine image with the latest version of Go installed
FROM golang:latest as builder

# Install Git for go get
RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git && \
    apt-get install ca-certificates -y

# Set ENV
ENV GOPATH /go/
ENV GO_WORKDIR $GOPATH/src/github.com/jdtoussa/avr-svc

# Set WORKDIR to go source code directory
WORKDIR $GO_WORKDIR

# Add files to image
ADD . $GO_WORKDIR

# Fetch Golang Dependency and Build Binary
RUN go get
RUN go install
RUN go test -v

# Build the service into a single binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .

# Start with a small image
FROM scratch

# Copy the SSL certs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary to the /app directory
COPY --from=builder /go/src/github.com/jdtoussa/avr-svc/main /app/
WORKDIR /app

EXPOSE 8080
CMD ["./main"]

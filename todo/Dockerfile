# Dockerfile
FROM golang:latest AS builder
ENV GO111MODULE=on

# Download modules
WORKDIR $GOPATH/src/github.com/thanhtuan260593/todo
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /todo .

# Run
FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /todo ./
EXPOSE 5000
ENTRYPOINT ["./todo"]
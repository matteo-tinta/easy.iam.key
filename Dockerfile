FROM ynap/certs as ssl
# intermediate image, certs are already present in /tmp/certs
# https://git.yoox.net/users/tintam/repos/docker.zscaler/browse

# Stage 1: Build the Go binary statically
FROM golang:1.24.5-bookworm AS builder

# SSL Certificates
COPY --from=ssl /tmp/certs/ /usr/local/share/ca-certificates/
RUN update-ca-certificates

# Install Git if needed by `go get`
RUN apt update && apt install git

WORKDIR /app

COPY ./bin/app.tar/ ./
RUN tar -xf app.tar --strip-components=1 -C ./

# Download dependencies
RUN go mod download

# Build statically-linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -o generate ./cmd/generate
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -extldflags=-static" -o generate.amd64 ./cmd/generate
RUN GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w -extldflags=-static" -o generate.exe ./cmd/generate

RUN CGO_ENABLED=0 GOOS=linux go build -o decode ./cmd/decode
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -extldflags=-static" -o decode.amd64 ./cmd/decode
RUN GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w -extldflags=-static" -o decode.exe ./cmd/decode

# Stage 2: Copy to scratch (final minimal image)
FROM scratch

WORKDIR /root/

COPY --from=builder /app/generate .
COPY --from=builder /app/decode .

CMD ["./generate"]

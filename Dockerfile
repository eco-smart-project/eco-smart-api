# Stage 1: Build the Go binary
FROM golang:latest AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# Specify the target architecture for the Go build
RUN GOOS=linux GOARCH=amd64 go build -o app/ ./...

# Final Stage
FROM alpine:latest

COPY --from=builder /build/app /app

# Ensure the binary is executable
RUN chmod +x /app

EXPOSE 8080

CMD ["/app"]

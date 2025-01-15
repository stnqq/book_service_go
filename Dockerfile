# Stage 1: Build
FROM golang:1.23.4 as builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main ./cmd/main.go

# Stage 2: Run
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]

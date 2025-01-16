FROM golang:1.23.4 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main ./cmd/main.go


FROM alpine:latest
RUN apk add --no-cache libc6-compat
WORKDIR /root/
COPY --from=builder /app/main .
RUN chmod +x main
EXPOSE 8080
CMD ["./main"]
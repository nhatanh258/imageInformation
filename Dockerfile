# Sử dụng Golang để build ứng dụng
FROM golang:1.24 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o myapp


# Sử dụng Alpine Linux để chạy ứng dụng (nhẹ và nhanh)
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/myapp .
COPY --from=builder /app/config.json .
COPY --from=builder /app/database.db .

# Mở cổng 8080
EXPOSE 8080

# Chạy ứng dụng
CMD ["./myapp"]

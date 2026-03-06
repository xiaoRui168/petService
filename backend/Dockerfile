# Stage 1: Build
FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# CGO_ENABLED=0 编译脱离环境依赖的静态程序
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o pet_server main.go

# Stage 2: Run
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/pet_server .
EXPOSE 8080
CMD ["./pet_server"]
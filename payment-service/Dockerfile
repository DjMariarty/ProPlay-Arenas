FROM golang:1.25-alpine AS builder

WORKDIR /app
RUN apk add --no-cache git ca-certificates

# Настройка GOPROXY с fallback для обхода проблем с сетью
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=sum.golang.org
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Сборка бинарника из директории с main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/app

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/app ./
EXPOSE 8084
CMD ["./app"]

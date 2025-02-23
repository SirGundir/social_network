# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app

# Устанавливаем GOPROXY
ENV GOPROXY=https://goproxy.io,direct

# Устанавливаем зависимости для сборки с CGO
RUN apk add --no-cache gcc musl-dev

# Копируем файлы модулей
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем приложение с включенным CGO
RUN CGO_ENABLED=1 GOOS=linux go build -o auth-service .

# Тестовая стадия
FROM golang:1.23-alpine AS tester
WORKDIR /app
COPY --from=builder /app /app
RUN go test -v ./... > test-results.txt 2>&1 || true

# Runtime stage
FROM alpine:latest
WORKDIR /app

# Устанавливаем зависимости для SQLite
RUN apk add --no-cache sqlite

# Копируем бинарный файл и конфигурации
COPY --from=builder /app/auth-service .
COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./auth-service"]
# Stage 1: Сборка приложения и мигратора
FROM golang:1.23-alpine AS builder

# Устанавливаем зависимости
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник основного приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/sso/main.go

# Собираем бинарник мигратора
RUN CGO_ENABLED=0 GOOS=linux go build -o migrator ./cmd/migrator/main.go

# Stage 2: Финальный образ
FROM alpine:latest

# Устанавливаем необходимые зависимости (если есть)
RUN apk add --no-cache bash

# Создаем рабочую директорию
WORKDIR /app

# Копируем бинарники из первого этапа
COPY --from=builder /app/app .
COPY --from=builder /app/migrator .
COPY --from=builder /app/migrations ./migrations

# Копируем конфигурационные файлы
COPY config ./config

# Запускаем мигратор, а затем основное приложение
CMD ["sh", "-c", "./migrator && ./app"]
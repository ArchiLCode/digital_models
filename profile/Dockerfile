# Stage 1: Сборка приложения и мигратора
FROM golang:1.23-alpine AS builder

# Устанавливаем зависимости
WORKDIR /profile
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник основного приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o profile ./cmd/profile/main.go

# Собираем бинарник мигратора
RUN CGO_ENABLED=0 GOOS=linux go build -o migrator ./cmd/migrator/main.go

# Stage 2: Финальный образ
FROM alpine:latest

# Устанавливаем необходимые зависимости (если есть)
RUN apk add --no-cache bash

# Создаем рабочую директорию
WORKDIR /profile

# Копируем бинарники из первого этапа
COPY --from=builder /profile/profile .
COPY --from=builder /profile/migrator .
COPY --from=builder /profile/migrations ./migrations

# Копируем конфигурационные файлы
COPY config ./config

# Запускаем мигратор, а затем основное приложение
CMD ["sh", "-c", "./migrator && ./profile"]
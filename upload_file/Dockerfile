FROM golang:1.23-alpine AS builder

# Установка необходимых пакетов
RUN apk add --no-cache git

# Установка рабочей директории
WORKDIR /app

# Копирование go.mod и go.sum (если есть)
COPY go.mod go.sum ./

# Загрузка зависимостей
RUN go mod download

# Копирование исходного кода
COPY . .

# Компиляция приложения
RUN CGO_ENABLED=0 GOOS=linux go build -o photo-upload-service ./cmd/

# Многоэтапная сборка для минимального образа
FROM alpine:3.18

# Добавление необходимых пакетов
RUN apk --no-cache add ca-certificates tzdata

# Установка рабочей директории
WORKDIR /app

# Копирование скомпилированного приложения из предыдущего этапа
COPY --from=builder /app/photo-upload-service .

# Создание директории для загрузки файлов
RUN mkdir -p /app/uploads && chmod 755 /app/uploads

# Экспорт порта
EXPOSE 8081

# Запуск приложения
CMD ["./photo-upload-service"]
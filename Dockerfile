# Этап 1: Сборка приложения (используем официальный образ Go)
FROM golang:1.24-alpine AS builder

# Устанавливаем зависимости для сборки (если нужны)
RUN apk add --no-cache git ca-certificates

# Рабочая директория
WORKDIR /backend

# Копируем файлы зависимостей
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем весь исходный код
COPY . .

# Собираем приложение (CGO отключен для статической сборки)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /backend/server ./cmd/app

# Этап 2: Финальный образ (используем минимальный alpine)
FROM alpine:3.19

# Устанавливаем tzdata для работы с временными зонами (опционально)
RUN apk add --no-cache tzdata

# Копируем бинарник из этапа builder
COPY --from=builder /backend/server /usr/local/bin/server

# Копируем статику/конфиги (если есть)
# COPY --from=builder /app/static /static

# Порт, который слушает приложение
EXPOSE 3000

# Команда запуска
CMD ["/usr/local/bin/server"]
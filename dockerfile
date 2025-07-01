FROM golang:1.24 AS builder

WORKDIR /app

# 1. Копируем только go.mod и go.sum — шаг для кеширования зависимостей
COPY go.mod go.sum ./

# 2. Загружаем зависимости (быстро, с кешем)
RUN go mod download

# 3. Копируем остальной исходный код
COPY . .

# 4. Сборка
RUN go build -o main main.go

# 5. Финальный минимальный образ (по желанию)
FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]

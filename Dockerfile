# Указываем базовый образ
FROM golang:1.21-alpine

# Устанавливаем зависимости
RUN apk add --no-cache git

# Устанавливаем рабочую директорию
WORKDIR /cmd

# Копируем модули и устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Компилируем приложение
RUN go build -o main ./cmd/main.go

# Определяем команду для запуска
CMD ["CONFIG_PATH=./config/local.yaml go run ./cmd/main.go"]

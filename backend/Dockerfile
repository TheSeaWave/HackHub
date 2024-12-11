# Используем базовый образ Go
FROM golang:1.22

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем файлы go.mod и go.sum для управления зависимостями
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем все файлы проекта в контейнер
COPY . .
# конфиг путь
COPY config/local.yaml /app/config/local.yaml

# Указываем полный путь к главному файлу
RUN go build -o main ./cmd/main.go

# Указываем порт, который слушает приложение
EXPOSE 8082

# Команда для запуска приложения
CMD ["./main"]

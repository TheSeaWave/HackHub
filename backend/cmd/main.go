package main

import (
	"HackHub/internal/config"
	"HackHub/internal/db"
	"HackHub/internal/lib/logger/sl"
	"HackHub/internal/server"
	"database/sql"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() { // CONFIG_PATH=./config/local.yaml go run ./cmd/main.go
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)

	log.Info("starting hackhub", slog.String("env", cfg.Env))
	log.Debug("debug message are enabled")

	// запуск сервера
	srv := server.NewServer()

	if err := srv.Run(":8082"); err != nil {
		log.Error("Ошибка запуска сервера", sl.Err(err))
	}

	// Подключаемся к базе данных
	dsn := "postgres://" + cfg.Database.User + ":" + cfg.Database.Password +
		"@" + cfg.Database.Host + ":" + cfg.Database.Port +
		"/" + cfg.Database.DBName + "?sslmode=" + cfg.Database.SSLMode

	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Error("Failed to connect to the database", sl.Err(err))
	}
	defer conn.Close()

	// Инициализируем базу данных
	db.InitializeDatabase(conn, "./internal/db/init.sql")

	log.Error("Service is running on", sl.Err(err))

	// TODO: Main development tasks

	// 1. API Endpoints
	// - Implement /surveys (POST, GET): Creation and retrieval of surveys.
	// - Implement /teams (POST, GET): Team management.
	// - Add middleware for JWT-based authentication and role-based access control.

	// 2. Authentication & Authorization
	// - Configure JWT token issuance and validation.
	// - Create middleware for validating requests with tokens.
	// - Add password hashing for secure storage.

	// 3. Static Content Serving
	// - Set up Go server to serve frontend assets from /static directory.
	// - Implement routing to distinguish between /api requests and static files.

	// 4. Database Integration
	// - Set up PostgreSQL connection pool using GORM or native database/sql.
	// - Migrate schema for users, surveys, and teams tables.
	// - Seed initial test data for development.

	// 5. Logging
	// - Integrate slog for structured logging.
	// - Add logs for request lifecycle (start, end, errors).
	// - Include logging for database queries and business logic.

	// 6. Testing
	// - Write unit tests for each API endpoint.
	// - Add integration tests for user flows (e.g., survey creation, team joining).
	// - Verify JWT handling with tests for valid and invalid tokens.

	// 7. Notifications
	// - Implement notification API and backend processing logic.
	// - Define data model for storing notifications in the database.

	// 8. Deployment Preparation
	// - Create Dockerfile for Go server and PostgreSQL.
	// - Test containerized deployment in a local environment.
	// - Document setup steps for production deployment.

}

func setupLogger(env string) *slog.Logger { // здесь прописывается в каком формате будет логирование на разных инваинментрах и с какого уровня показывать логи
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

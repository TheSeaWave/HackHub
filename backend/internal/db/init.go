package db

import (
	"context"
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib" // Подключаем драйвер pgx
)

func InitializeDatabase(db *sql.DB, sqlFilePath string) {
	// Загружаем SQL-скрипт из файла
	sqlContent, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		log.Fatalf("Failed to read SQL file: %v", err)
	}

	// Выполняем SQL-скрипт
	ctx := context.Background()
	_, err = db.ExecContext(ctx, string(sqlContent))
	if err != nil {
		log.Fatalf("Failed to execute SQL script: %v", err)
	}

	log.Println("Database initialized successfully!")
}

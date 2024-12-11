package database // sudo -u postgres psql
//SELECT * FROM teams;

import (
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL драйвер
)

var DB *sqlx.DB

// InitDB инициализирует подключение к базе данных
func InitDB(dataSourceName string) {
	var err error
	// Измените dataSourceName на 'database' для подключения к сервису базы данных
	DB, err = sqlx.Connect("postgres", "postgres://postgres:postgres@database:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	log.Println("Успешное подключение к базе данных")
}

// CloseDB закрывает соединение с базой данных
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Printf("Ошибка при закрытии базы данных: %v", err)
	}
}

// --- Методы для пользователей ---
type User struct {
	ID           int       `db:"id"`
	Name         string    `db:"name"`
	Email        string    `db:"email"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `db:"created_at"`
}

// CreateUser добавляет нового пользователя в базу данных
func CreateUser(user *User) error {
	query := `INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3) RETURNING id, created_at`
	err := DB.QueryRow(query, user.Name, user.Email, user.PasswordHash).Scan(&user.ID, &user.CreatedAt)
	return err
}

// GetUserByEmail ищет пользователя по email
func GetUserByEmail(email string) (*User, error) {
	var user User
	err := DB.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// --- Методы для команд ---
type Team struct {
	ID          int       `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	CaptainID   int       `db:"captain_id"`
	Avatar      string    `db:"avatar"`
	CreatedAt   time.Time `db:"created_at"`
}

// CreateTeam добавляет новую команду в базу данных
func CreateTeam(team *Team) error {
	query := `INSERT INTO teams (name, description, captain_id, avatar) VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	err := DB.QueryRow(query, team.Name, team.Description, team.CaptainID, team.Avatar).Scan(&team.ID, &team.CreatedAt)
	return err
}

// GetTeamByID ищет команду по ID
func GetTeamByID(id int) (*Team, error) {
	var team Team
	err := DB.Get(&team, "SELECT * FROM teams WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &team, nil
}

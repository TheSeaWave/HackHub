package server

import (
	database "HackHub/internal/repository"
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

// Структура команды
type Team struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CaptainID   int       `json:"captain_id" binding:"required"`
	Avatar      string    `json:"avatar" binding:"required,url"`
	Members     []Member  `json:"members"`
	CreatedAt   time.Time `json:"created_at"`
}

type Member struct {
	FullName string `json:"full_name" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Telegram string `json:"telegram" binding:"required"`
}

// var teams []Team
// var teamID = 1

// Структуры и переменные:
type User struct {
	ID           int       `json:"id"`
	Name         string    `json:"name" binding:"required"`
	Email        string    `json:"email" binding:"required,email"`
	Password     string    `json:"password" binding:"required"`
	PasswordHash string    `json:"-"` // Для хранения хэша пароля
	CreatedAt    time.Time `json:"created_at"`
}

// var users []User

// var userID = 1

var jwtSecret = []byte("supersecretkey") // Секрет для подписи JWT

// Survey представляет анкету участника.
type Survey struct {
	ID            int       `json:"id"`
	FullName      string    `json:"full_name" binding:"required"`
	Group         string    `json:"group" binding:"required"`
	Telegram      string    `json:"telegram" binding:"required"`
	Role          string    `json:"role" binding:"required"`
	Stack         []string  `json:"stack" binding:"required"`
	About         string    `json:"about" binding:"required"`
	Achievements  []string  `json:"achievements" binding:"required"`
	Status        string    `json:"status" binding:"required"`
	PortfolioLink string    `json:"portfolio_link" binding:"required,url"`
	Avatar        string    `json:"avatar" binding:"required,url"`
	Teams         []string  `json:"teams" binding:"required"`
	LastUpdated   time.Time `json:"last_updated"`
	Rating        int       `json:"Rating"`
	Experience    int       `json:"Experience"`
	Like          int       `json:"Like"`
}

// var (
// 	surveys []Survey
// 	nextID  = 1
// )

// Server представляет серверное приложение.
type Server struct {
	router *gin.Engine
}

// NewServer создает новый сервер.
func NewServer() *Server {
	r := gin.Default()

	server := &Server{
		router: r,
	}

	server.routes()

	// Определяем абсолютный путь к папке build
	wd, err := os.Getwd() // получаем рабочую директорию
	if err != nil {
		panic("Не удалось получить рабочую директорию: " + err.Error())
	}
	buildPath := filepath.Join(wd, "build")

	// Подключаем статические файлы
	r.Use(static.Serve("/", static.LocalFile(buildPath, true)))

	// Fallback для SPA
	r.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(buildPath, "index.html"))
	})
	return server
}

// routes определяет маршруты для сервера.
func (s *Server) routes() {
	s.router.POST("/surveys", s.createSurvey)
	s.router.GET("/surveys", s.getSurveys)
	s.router.PUT("/surveys/:id", s.updateSurvey)
	s.router.DELETE("/surveys/:id", s.deleteSurvey)
	s.authRoutes()
	s.teamRoutes()
	s.router.GET("/health", s.healthCheck)
}

// Run запускает сервер на указанном адресе.
func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}

// healthCheck возвращает статус 200 OK для проверки работоспособности сервера
func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

// createSurvey обрабатывает запрос на создание анкеты.
func (s *Server) createSurvey(c *gin.Context) {
	var newSurvey Survey
	if err := c.ShouldBindJSON(&newSurvey); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newSurvey.LastUpdated = time.Now().UTC()

	query := `
        INSERT INTO surveys (full_name, "group", telegram, role, stack, about, achievements, status, portfolio_link, avatar, teams, last_updated, rating, experience, "like")
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id;
    `
	var newID int
	err := database.DB.QueryRow(query, newSurvey.FullName, newSurvey.Group, newSurvey.Telegram, newSurvey.Role,
		pq.Array(newSurvey.Stack), newSurvey.About, pq.Array(newSurvey.Achievements), newSurvey.Status,
		newSurvey.PortfolioLink, newSurvey.Avatar, pq.Array(newSurvey.Teams), newSurvey.LastUpdated,
		newSurvey.Rating, newSurvey.Experience, newSurvey.Like).Scan(&newID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания анкеты: " + err.Error()})
		return
	}

	newSurvey.ID = newID

	c.JSON(http.StatusCreated, gin.H{
		"id":      newSurvey.ID,
		"message": "Анкета успешно создана",
	})
}

// getSurveys обрабатывает запрос на получение списка анкет.
func (s *Server) getSurveys(c *gin.Context) {
	query := `
        SELECT id, full_name, "group", telegram, role, stack, about, achievements, status, portfolio_link, avatar, teams, last_updated, rating, experience, "like"
        FROM surveys
    `
	rows, err := database.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения анкет: " + err.Error()})
		return
	}
	defer rows.Close()

	var surveys []Survey
	for rows.Next() {
		var survey Survey
		err := rows.Scan(&survey.ID, &survey.FullName, &survey.Group, &survey.Telegram, &survey.Role,
			pq.Array(&survey.Stack), &survey.About, pq.Array(&survey.Achievements), &survey.Status,
			&survey.PortfolioLink, &survey.Avatar, pq.Array(&survey.Teams), &survey.LastUpdated,
			&survey.Rating, &survey.Experience, &survey.Like)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обработки анкет: " + err.Error()})
			return
		}
		surveys = append(surveys, survey)
	}

	c.JSON(http.StatusOK, surveys)
}

// updateSurvey обрабатывает запрос на обновление анкеты.
func (s *Server) updateSurvey(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID анкеты"})
		return
	}

	var updatedSurvey Survey
	if err := c.ShouldBindJSON(&updatedSurvey); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSurvey.LastUpdated = time.Now().UTC()

	query := `
        UPDATE surveys
        SET full_name = $1, "group" = $2, telegram = $3, role = $4, stack = $5, about = $6,
            achievements = $7, status = $8, portfolio_link = $9, avatar = $10, teams = $11,
            last_updated = $12, rating = $13, experience = $14, "like" = $15
        WHERE id = $16;
    `
	_, err = database.DB.Exec(query, updatedSurvey.FullName, updatedSurvey.Group, updatedSurvey.Telegram,
		updatedSurvey.Role, pq.Array(updatedSurvey.Stack), updatedSurvey.About, pq.Array(updatedSurvey.Achievements),
		updatedSurvey.Status, updatedSurvey.PortfolioLink, updatedSurvey.Avatar, pq.Array(updatedSurvey.Teams),
		updatedSurvey.LastUpdated, updatedSurvey.Rating, updatedSurvey.Experience, updatedSurvey.Like, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления анкеты: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Анкета успешно обновлена"})
}

// deleteSurvey обрабатывает запрос на удаление анкеты.
func (s *Server) deleteSurvey(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID анкеты"})
		return
	}

	query := `DELETE FROM surveys WHERE id = $1;`
	_, err = database.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка удаления анкеты: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Анкета успешно удалена"})
}

// filterSurveys фильтрует анкеты по указанным критериям.
// func (s *Server) filterSurveys(role, stack, status string) []Survey {
// 	filtered := []Survey{}

// 	for _, survey := range surveys {
// 		if role != "" && survey.Role != role {
// 			continue
// 		}
// 		if stack != "" && !contains(survey.Stack, stack) {
// 			continue
// 		}
// 		if status != "" && survey.Status != status {
// 			continue
// 		}

// 		filtered = append(filtered, survey)
// 	}

// 	return filtered
// }

func (s *Server) authRoutes() {
	s.router.POST("/auth/register", s.registerUser)
	s.router.POST("/auth/login", s.loginUser)
	s.router.GET("/auth/me", s.getCurrentUser, s.authMiddleware)
}

func (s *Server) registerUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные: " + err.Error()})
		return
	}

	// Проверяем, есть ли пользователь с таким email
	var existingID int
	err := database.DB.QueryRow(`SELECT id FROM users WHERE email = $1;`, newUser.Email).Scan(&existingID)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка проверки email: " + err.Error()})
		return
	}
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email уже зарегистрирован"})
		return
	}

	// Хэшируем пароль
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при хэшировании пароля"})
		return
	}
	newUser.PasswordHash = string(passwordHash)

	// Сохраняем пользователя в базу данных
	query := `
        INSERT INTO users (name, email, password_hash, created_at)
        VALUES ($1, $2, $3, NOW()) RETURNING id;
    `
	var newID int
	err = database.DB.QueryRow(query, newUser.Name, newUser.Email, newUser.PasswordHash).Scan(&newID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при добавлении пользователя: " + err.Error()})
		return
	}
	newUser.ID = newID

	c.JSON(http.StatusCreated, gin.H{
		"message": "Пользователь успешно зарегистрирован",
		"user":    newUser,
	})
}

func (s *Server) loginUser(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные: " + err.Error()})
		return
	}

	// Ищем пользователя в базе данных
	var user User
	query := `SELECT id, name, email, password_hash FROM users WHERE email = $1;`
	err := database.DB.QueryRow(query, credentials.Email).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при авторизации: " + err.Error()})
		}
		return
	}

	// Проверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		return
	}

	// Генерируем JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":    time.Now().Add(24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при генерации токена: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
func (s *Server) getCurrentUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не аутентифицирован"})
		return
	}

	var user User
	query := `SELECT id, name, email, created_at FROM users WHERE id = $1;`
	err := database.DB.QueryRow(query, userID).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении пользователя: " + err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
func (s *Server) authMiddleware(c *gin.Context) {
	// Получаем заголовок Authorization
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Требуется аутентификация"})
		c.Abort()
		return
	}

	// Извлекаем токен
	tokenString := authHeader[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неверный метод подписи")
		}
		return jwtSecret, nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен: " + err.Error()})
		c.Abort()
		return
	}

	// Извлекаем данные из токена
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["userID"].(float64)) // Токен содержит ID пользователя
		c.Set("userID", userID)
		c.Next() // Продолжаем выполнение запроса
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
		c.Abort()
	}
}

// 123
func (s *Server) teamRoutes() {
	s.router.POST("/teams", s.createTeam)
	s.router.GET("/teams/:id", s.getTeam)
}

func (s *Server) createTeam(c *gin.Context) {
	var newTeam Team
	if err := c.ShouldBindJSON(&newTeam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTeam.CreatedAt = time.Now().UTC()

	// Сохраняем команду в базе данных
	query := `
        INSERT INTO teams (name, description, captain_id, avatar, created_at)
        VALUES ($1, $2, $3, $4, $5) RETURNING id;
    `
	var newID int
	err := database.DB.QueryRow(query, newTeam.Name, newTeam.Description, newTeam.CaptainID, newTeam.Avatar, newTeam.CreatedAt).Scan(&newID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать команду: " + err.Error()})
		return
	}

	newTeam.ID = newID

	c.JSON(http.StatusCreated, gin.H{
		"id":      newTeam.ID,
		"message": "Команда успешно создана",
	})
}

func (s *Server) getTeam(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID команды"})
		return
	}

	var team Team
	query := `
        SELECT id, name, description, captain_id, avatar, created_at
        FROM teams
        WHERE id = $1;
    `
	err = database.DB.QueryRow(query, id).Scan(&team.ID, &team.Name, &team.Description, &team.CaptainID, &team.Avatar, &team.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Команда не найдена"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения команды: " + err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, team)
}

func calculateRating(survey Survey) int {
	rating := 0.0

	// Условие для новичков
	if survey.Experience < 1 && len(survey.Achievements) > 1 {
		rating += 5.0
	}

	// Рассчитываем рейтинг для остальных пользователей
	if survey.Experience > 0 {
		achievementsCount := float64(len(survey.Achievements))
		stackCount := float64(len(survey.Stack))
		experience := float64(survey.Experience)
		likeCount := float64(survey.Like)

		// Основная формула
		r := (0.5 * stackCount) +
			(4.0 * (achievementsCount / experience)) +
			(0.25 * likeCount)

		// Условие для начинающих пользователей с малым опытом и достижениями
		if (achievementsCount/experience) < 2 && experience < 2 {
			r *= 2.0
		}

		rating += r
	}

	// Округление результата до ближайшего целого числа
	return int(math.Round(rating))
}

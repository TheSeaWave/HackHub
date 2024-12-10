package server

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

var teams []Team
var teamID = 1

// Структуры и переменные:
type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password" binding:"required"`
	PasswordHash string `json:"-"` // Для хранения хэша пароля
}

var users []User
var userID = 1

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

var (
	surveys []Survey
	nextID  = 1
)

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

}

// Run запускает сервер на указанном адресе.
func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}

// createSurvey обрабатывает запрос на создание анкеты.
func (s *Server) createSurvey(c *gin.Context) {
	var newSurvey Survey

	if err := c.ShouldBindJSON(&newSurvey); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newSurvey.ID = nextID
	newSurvey.LastUpdated = time.Now().UTC()
	newSurvey.Rating = calculateRating(newSurvey)
	nextID++

	surveys = append(surveys, newSurvey)

	c.JSON(http.StatusCreated, gin.H{
		"id":      newSurvey.ID,
		"message": "Анкета успешно создана",
	})
}

// getSurveys обрабатывает запрос на получение списка анкет.
func (s *Server) getSurveys(c *gin.Context) {
	role := c.Query("role")
	stack := c.Query("stack")
	status := c.Query("status")

	filtered := s.filterSurveys(role, stack, status)

	c.JSON(http.StatusOK, filtered)
}

// updateSurvey обрабатывает запрос на обновление анкеты.
func (s *Server) updateSurvey(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	var updatedSurvey Survey
	if err := c.ShouldBindJSON(&updatedSurvey); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, survey := range surveys {
		if survey.ID == id {
			updatedSurvey.ID = survey.ID
			updatedSurvey.LastUpdated = time.Now().UTC()

			surveys[i] = updatedSurvey
			c.JSON(http.StatusOK, gin.H{"message": "Анкета успешно обновлена"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Анкета не найдена"})
}

// deleteSurvey обрабатывает запрос на удаление анкеты.
func (s *Server) deleteSurvey(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	for i, survey := range surveys {
		if survey.ID == id {
			surveys = append(surveys[:i], surveys[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Анкета успешно удалена"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Анкета не найдена"})
}

// filterSurveys фильтрует анкеты по указанным критериям.
func (s *Server) filterSurveys(role, stack, status string) []Survey {
	filtered := []Survey{}

	for _, survey := range surveys {
		if role != "" && survey.Role != role {
			continue
		}
		if stack != "" && !contains(survey.Stack, stack) {
			continue
		}
		if status != "" && survey.Status != status {
			continue
		}

		filtered = append(filtered, survey)
	}

	return filtered
}

// contains проверяет, содержит ли срез указанный элемент.
func contains(slice []string, item string) bool {
	for _, elem := range slice {
		if elem == item {
			return true
		}
	}
	return false
}

func (s *Server) authRoutes() {
	s.router.POST("/auth/register", s.registerUser)
	s.router.POST("/auth/login", s.loginUser)
	s.router.GET("/auth/me", s.authMiddleware, s.getCurrentUser)
}

func (s *Server) registerUser(c *gin.Context) {
	var newUser User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные: " + err.Error()})
		return
	}

	for _, user := range users {
		if user.Email == newUser.Email {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email уже зарегистрирован"})
			return
		}
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить пароль"})
		return
	}

	newUser.PasswordHash = string(passwordHash)
	newUser.Password = "" // Убираем пароль из ответа
	newUser.ID = userID
	userID++
	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{"message": "Пользователь успешно зарегистрирован", "user": newUser})
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

	for _, user := range users {
		if user.Email == credentials.Email {
			err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(credentials.Password))
			if err == nil {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"userID": user.ID,
					"exp":    time.Now().Add(time.Hour * 24).Unix(),
				})

				tokenString, err := token.SignedString(jwtSecret)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токен"})
					return
				}

				c.JSON(http.StatusOK, gin.H{"token": tokenString})
				return
			}
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
}

func (s *Server) getCurrentUser(c *gin.Context) {
	userID, _ := c.Get("userID")
	for _, user := range users {
		if user.ID == userID {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Пользователь не найден"})
}

func (s *Server) authMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Требуется аутентификация"})
		c.Abort()
		return
	}

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

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID := int(claims["userID"].(float64))
		c.Set("userID", userID)
		c.Next()
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
		c.Abort()
	}
}

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

	newTeam.ID = teamID
	newTeam.CreatedAt = time.Now().UTC()
	teamID++
	teams = append(teams, newTeam)

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

	for _, team := range teams {
		if team.ID == id {
			c.JSON(http.StatusOK, team)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Команда не найдена"})
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

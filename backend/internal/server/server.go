package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

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

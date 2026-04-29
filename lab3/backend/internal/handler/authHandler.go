package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	// TODO: добавить сервис при реализации
	// auth service.AuthServiceInterface
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

// POST /api/auth/signup
func (h *AuthHandler) SignUp(c *gin.Context) {
	slog.Info("auth handler: sign up: initiated")

	var userData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
	}

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// TODO: вызвать сервис для регистрации
	resultId := 1 // h.auth.PostOne(c.Request.Context(), &userData)

	slog.Info("auth handler: sign up: succeeded (stub)")
	
	c.JSON(http.StatusOK, gin.H{
		"addedId": resultId,
		"message": "User created successfully (stub)",
	})
}

// POST /api/auth/signin
func (h *AuthHandler) SignIn(c *gin.Context) {
	slog.Info("auth handler: sign in: initiated")

	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// TODO: вызвать сервис для аутентификации
	// user, err := h.auth.GetOneByEmailFull(c.Request.Context(), creds.Email)
	// token, err := h.auth.GenerateAuthToken(...)

	slog.Info("auth handler: sign in: succeeded (stub)")
	c.JSON(http.StatusOK, gin.H{
		"token":  "stub-jwt-token",
		"role":   "user",
		"userId": 1,
	})
}

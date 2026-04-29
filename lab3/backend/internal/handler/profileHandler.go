package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

type ProfileHandler struct {
	// TODO: добавить сервис при реализации
	// profile service.ProfileServiceInterface
}

func NewProfileHandler() *ProfileHandler {
	return &ProfileHandler{}
}


// GET /web/profile
func (h *ProfileHandler) GetProfilePage(c *gin.Context) {
	slog.Info("profile handler: get profile page: initiated")

	// Заглушка данных пользователя
	user := gin.H{
		"ID":       1,
		"Email":    "user@example.com",
		"Nickname": "User",
		"Role":     "user",
	}

	c.HTML(http.StatusOK, "profile.html", gin.H{
		"Title": "Профиль пользователя",
		"User":  user,
	})
}

// GetProfile — получение профиля (JSON API)
// GET /api/profile
func (h *ProfileHandler) GetProfile(c *gin.Context) {
	slog.Info("profile handler: get profile: initiated")

	// Заглушка данных пользователя
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": gin.H{
			"id":       1,
			"email":    "user@example.com",
			"nickname": "User",
			"role_id":  1,
		},
	})
}

// UpdateProfile — обновление профиля (JSON API)
// POST /api/profile
func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	slog.Info("profile handler: update profile: initiated")

	var profileData struct {
		Nickname string `json:"nickname"`
		Email    string `json:"email"`
	}

	if err := c.ShouldBindJSON(&profileData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// TODO: вызвать сервис для обновления профиля
	// err := h.profile.Update(c.Request.Context(), userID, &profileData)

	slog.Info("profile handler: update profile: succeeded (stub)")
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Profile updated",
		"data": gin.H{
			"nickname": profileData.Nickname,
			"email":    profileData.Email,
		},
	})
}

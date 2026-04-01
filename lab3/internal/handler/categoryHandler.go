package handler

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

type CategoryHandler struct {
	// TODO: добавить сервис при реализации
	// category service.CategoryServiceInterface
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

// GET /web/category/:id
func (h *CategoryHandler) GetCategoryPage(c *gin.Context) {
	slog.Info("category handler: get category page: initiated")

	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)

	// Заглушка
	category := gin.H{
		"ID":          id,
		"Name":        "Category " + idStr,
		"Description": "Description for category " + idStr,
	}

	products := []gin.H{
		{"ID": 1, "Name": "Product 1", "Price": 99.99},
		{"ID": 2, "Name": "Product 2", "Price": 149.99},
	}

	c.HTML(http.StatusOK, "category.html", gin.H{
		"Title":      "Категория",
		"Category":   category,
		"Products":   products,
	})
}

// GET /api/category/:id
func (h *CategoryHandler) GetCategory(c *gin.Context) {
	slog.Info("category handler: get category: initiated")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category id"})
		return
	}

	// Заглушка
	category := gin.H{
		"id":          id,
		"name":        "Category " + idStr,
		"description": "Description for category " + idStr,
		"products": []gin.H{
			{"id": 1, "name": "Product 1", "price": 99.99},
			{"id": 2, "name": "Product 2", "price": 149.99},
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   category,
	})
}

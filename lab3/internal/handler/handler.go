package handler

import (
	"github.com/gin-gonic/gin"
)

// Интерфейсы для хендлеров — позволяют легко подключить сервисы later
type AuthHandlerInterface interface {
	SignUp(c *gin.Context)
	SignIn(c *gin.Context)
}

type ProfileHandlerInterface interface {
	GetProfilePage(c *gin.Context) // HTML страница
	GetProfile(c *gin.Context)     // JSON API
	UpdateProfile(c *gin.Context)  // JSON API
}

type ProductHandlerInterface interface {
	GetProductsPage(c *gin.Context) // HTML страница
	GetProductPage(c *gin.Context)  // HTML страница
	GetProducts(c *gin.Context)     // JSON API
	CreateProduct(c *gin.Context)   // JSON API
	GetProductByID(c *gin.Context)  // JSON API
}

type CategoryHandlerInterface interface {
	GetCategoryPage(c *gin.Context) // HTML страница
	GetCategory(c *gin.Context)     // JSON API
}

type Handler struct {
	Auth     AuthHandlerInterface
	Profile  ProfileHandlerInterface
	Product  ProductHandlerInterface
	Category CategoryHandlerInterface
}

// Сейчас используются заглушки
func NewHandler() *Handler {
	return &Handler{
		Auth:     NewAuthHandler(),
		Profile:  NewProfileHandler(),
		Product:  NewProductHandler(),
		Category: NewCategoryHandler(),
	}
}

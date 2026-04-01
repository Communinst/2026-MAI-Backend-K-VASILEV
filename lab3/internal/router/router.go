package router

import (
	"os"

	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/internal/handler"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/internal/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct {
	handler *handler.Handler
}

func NewRouter(h *handler.Handler) *Router {
	return &Router{
		handler: h,
	}
}

func (r *Router) Init(trustedProxies []string) *gin.Engine {

	router := gin.Default()
	router.SetTrustedProxies(trustedProxies)
	router.Static("/static", os.Getenv("STATIC_PATH"))
	router.LoadHTMLGlob(os.Getenv("TEMPLATES_PATH"))

	web := router.Group("/web")
	{
		web.GET("/profile", r.handler.Profile.GetProfilePage)
		web.GET("/products", r.handler.Product.GetProductsPage)
		web.GET("/products/:id", r.handler.Product.GetProductPage)
		web.GET("/category/:id", r.handler.Category.GetCategoryPage)
	}

	// /api/ — JSON API
	api := router.Group("/api")
	{
		// Auth
		api.POST("/auth/signup", r.handler.Auth.SignUp)
		api.POST("/auth/signin", r.handler.Auth.SignIn)

		// Profile — только GET/POST
		api.	GET("/profile", middleware.RequireMethods("GET"), r.handler.Profile.GetProfile)
		api.POST("/profile", middleware.RequireMethods("POST"), r.handler.Profile.UpdateProfile)

		// Products
		api.GET("/products", middleware.RequireMethods("GET"), r.handler.Product.GetProducts)
		api.POST("/products", middleware.RequireMethods("POST"), r.handler.Product.CreateProduct)
		api.GET("/products/:id", middleware.RequireMethods("GET"), r.handler.Product.GetProductByID)

		// Categories
		api.GET("/category/:id", middleware.RequireMethods("GET"), r.handler.Category.GetCategory)
	}

	return router
}

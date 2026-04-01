package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	// TODO: добавить сервис при реализации
	// product service.ProductServiceInterface
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

// GET /web/products
func (h *ProductHandler) GetProductsPage(c *gin.Context) {
	slog.Info("product handler: get products page: initiated")

	// Заглушка
	products := []gin.H{
		{"ID": 1, "Name": "Product 1", "Price": 99.99, "Category": "Electronics"},
		{"ID": 2, "Name": "Product 2", "Price": 149.99, "Category": "Electronics"},
		{"ID": 3, "Name": "Product 3", "Price": 29.99, "Category": "Books"},
	}

	c.HTML(http.StatusOK, "products.html", gin.H{
		"Title":    "Список продуктов",
		"Products": products,
	})
}


// GET /api/products
func (h *ProductHandler) GetProducts(c *gin.Context) {
	slog.Info("product handler: get products: initiated")

	// Заглушка данных продуктов
	products := []gin.H{
		{"id": 1, "name": "Product 1", "price": 99.99, "category_id": 1},
		{"id": 2, "name": "Product 2", "price": 149.99, "category_id": 1},
		{"id": 3, "name": "Product 3", "price": 29.99, "category_id": 2},
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   products,
	})
}

// POST /api/products
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	slog.Info("product handler: create product: initiated")

	var productData struct {
		Name       string  `json:"name"`
		Price      float64 `json:"price"`
		CategoryID uint    `json:"category_id"`
	}

	if err := c.ShouldBindJSON(&productData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// TODO: вызвать сервис для создания продукта
	// id, err := h.product.Create(c.Request.Context(), &productData)

	slog.Info("product handler: create product: succeeded (stub)")
	c.JSON(http.StatusCreated, gin.H{
		"status":  "ok",
		"message": "Product created",
		"data": gin.H{
			"id":          4,
			"name":        productData.Name,
			"price":       productData.Price,
			"category_id": productData.CategoryID,
		},
	})
}

// GET /api/products/:id
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	slog.Info("product handler: get product by id: initiated")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
		return
	}

	// Заглушка
	product := gin.H{
		"id":          id,
		"name":        "Product " + idStr,
		"price":       99.99,
		"category_id": 1,
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   product,
	})
}


// GET /web/products/:id
func (h *ProductHandler) GetProductPage(c *gin.Context) {
	slog.Info("product handler: get product page: initiated")

	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 64)

	// Заглушка
	product := gin.H{
		"ID":          id,
		"Name":        "Product " + idStr,
		"Price":       99.99,
		"Category":    "Electronics",
		"Description": "Это описание для продукта #" + idStr + ". Здесь может быть подробная информация о товаре.",
	}

	c.HTML(http.StatusOK, "product.html", gin.H{
		"Title":   product["Name"],
		"Product": product,
	})
}

package handler

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/models"
	"github.com/Communinst/2026-MAI-Backend-K-VASILEV/lab3/backend/internal/service"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	product service.ProductServiceInterface
}

func NewProductHandler(productService service.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		product: productService,
	}
}

// GET /web/products
func (h *ProductHandler) GetProductsPage(c *gin.Context) {
	slog.Info("product handler: get products page: initiated")

	products, err := h.product.GetProducts(c.Request.Context())
	if err != nil {
		c.HTML(http.StatusInternalServerError, "products.html", gin.H{
			"Title": "Ошибка загрузки",
			"Error": "Не удалось загрузить списки",
		})
		return
	}

	c.HTML(http.StatusOK, "products.html", gin.H{
		"Title":    "Список продуктов",
		"Products": products,
	})
}

// GET /api/products
func (h *ProductHandler) GetProducts(c *gin.Context) {
	slog.Info("product handler: get products: initiated")

	products, err := h.product.GetProducts(c.Request.Context())
	if err != nil {
		slog.Error("product handler: get products: failed", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   products,
	})
}

// POST /api/products
func (h *ProductHandler) CreateProduct(c *gin.Context) {
	slog.Info("product handler: create product: initiated")

	var productData models.Product

	if err := c.ShouldBindJSON(&productData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	productData.ID = 0

	id, err := h.product.CreateProduct(c.Request.Context(), &productData)
	if err != nil {
		slog.Error("product handler: failed to create it", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	slog.Info("product handler: create product: succeeded")
	c.JSON(http.StatusCreated, gin.H{
		"status":  "ok",
		"message": "Product created",
		"data": gin.H{
			"id": id,
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

	product, err := h.product.GetProductByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   product,
	})
}

// GET /api/search
func (h *ProductHandler) SearchProducts(c *gin.Context) {
	slog.Info("product handler: search products: initiated")

	q := c.Query("q")
	if q == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing 'q' query parameter"})
		return
	}

	results, err := h.product.SearchProducts(c.Request.Context(), q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   results,
	})
}

// GET /web/products/:id
func (h *ProductHandler) GetProductPage(c *gin.Context) {
	slog.Info("product handler: get product page: initiated")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "product.html", gin.H{"Error": "Invalid id"})
		return
	}

	product, err := h.product.GetProductByID(c.Request.Context(), id)
	if err != nil {
		c.HTML(http.StatusNotFound, "product.html", gin.H{"Error": "Product not found"})
		return
	}

	c.HTML(http.StatusOK, "product.html", gin.H{
		"Title":   product.Name,
		"Product": product,
	})
}

// PUT /api/products/:id
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	slog.Info("product handler: update product: initiated")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
		return
	}

	var productData models.Product
	if err := c.ShouldBindJSON(&productData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := h.product.UpdateProduct(c.Request.Context(), id, &productData); err != nil {
		slog.Error("product handler: failed to update", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Product updated"})
}

// DELETE /api/products/:id
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	slog.Info("product handler: delete product: initiated")

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product id"})
		return
	}

	if err := h.product.DeleteProduct(c.Request.Context(), id); err != nil {
		slog.Error("product handler: failed to delete", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Product deleted"})
}

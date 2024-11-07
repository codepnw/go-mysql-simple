package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/codepnw/go-mysql-simple/internal/database"
	"github.com/gin-gonic/gin"
)

type IProductHandler interface {
	CreateProduct(c *gin.Context)
	GetProducts(c *gin.Context)
	GetProduct(c *gin.Context)
	UpdateProduct(c *gin.Context)
	DeleteProduct(c *gin.Context)
}

type productHandler struct {
	db *database.Queries
}

type productReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewProducts(db *database.Queries) IProductHandler {
	return &productHandler{db: db}
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	request := productReq{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	product, err := h.db.CreateProduct(context.Background(), database.CreateProductParams{
		Title:       request.Title,
		Description: request.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *productHandler) GetProducts(c *gin.Context) {
	products, err := h.db.ListProducts(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *productHandler) GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	product, err := h.db.GetOneProduct(context.Background(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	request := productReq{}

	err := h.db.UpdateProduct(context.Background(), database.UpdateProductParams{
		Title:       request.Title,
		Description: request.Description,
		ID: int64(id),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	product, err := h.db.GetOneProduct(context.Background(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *productHandler) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.db.DeleteProduct(context.Background(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "product deleted"})
}

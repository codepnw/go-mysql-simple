package routes

import (
	"github.com/codepnw/go-mysql-simple/internal/database/migrations"
	"github.com/codepnw/go-mysql-simple/internal/handlers"
	"github.com/gin-gonic/gin"
)

func NewRoutes(r *gin.Engine, db *migrations.Queries) {
	handler := handlers.NewProducts(db)
	g := r.Group("/products")

	g.POST("/", handler.CreateProduct)
	g.GET("/", handler.GetProducts)
	g.GET("/:id", handler.GetProduct)
	g.PATCH("/:id", handler.UpdateProduct)
	g.DELETE("/:id", handler.DeleteProduct)
}

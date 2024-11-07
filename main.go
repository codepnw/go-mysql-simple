package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/codepnw/go-mysql-simple/internal/database"
	"github.com/codepnw/go-mysql-simple/internal/handlers"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("dev.env")

	conn, err := sql.Open("mysql", os.Getenv("MYSQL_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	db := database.New(conn)

	r := gin.Default()
	g := r.Group("/products")

	handler := handlers.NewProducts(db)

	g.POST("/", handler.CreateProduct)
	g.GET("/", handler.GetProducts)
	g.GET("/:id", handler.GetProduct)
	g.PATCH("/:id", handler.UpdateProduct)
	g.DELETE("/:id", handler.DeleteProduct)

	r.Run(":8080")
}

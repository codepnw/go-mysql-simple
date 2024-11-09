package main

import (
	"log"
	"os"

	"github.com/codepnw/go-mysql-simple/internal/database"
	"github.com/codepnw/go-mysql-simple/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const envFile = "dev.env"

func main() {
	if err := godotenv.Load(envFile); err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	conn := database.ConnectDatabase(os.Getenv("MYSQL_DSN"))

	routes.NewRoutes(r, conn)

	r.Run(":" + os.Getenv("APP_PORT"))
}

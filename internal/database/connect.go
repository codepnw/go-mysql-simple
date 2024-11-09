package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/codepnw/go-mysql-simple/internal/database/migrations"
)

func ConnectDatabase(dsn string) *migrations.Queries {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	log.Println("database connected...")

	return migrations.New(conn)
}

package bootstrap

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)


func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	}
	
	func NewDB() (*sql.DB, error) {
		
		_ = godotenv.Load()

	dbURL := os.ExpandEnv("$DATABASE_USER:$DATABASE_PASSWORD@tcp($DATABASE_HOST:$DATABASE_PORT)/$DATABASE_NAME")
	
	log.Println(dbURL)
// "root:root@tcp(127.0.0.1:3336)/api_go"
	db, err:= sql.Open("mysql",dbURL)

	if err != nil {
		return nil, err
	}

	return db, nil
}

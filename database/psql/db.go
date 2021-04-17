package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func OpenDB() {
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("DB_PORT")
		user     = os.Getenv("DB_USER")
		password = os.Getenv("DB_PASS")
		dbname   = os.Getenv("DB_NAME")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("An error occured while connecting to database: ", err)
	} else {
		log.Printf("Connected to database %s\n", dbname)
	}

	// database pooling
	db.DB.SetMaxIdleConns(10)
	db.DB.SetMaxOpenConns(10)
	db.DB.SetConnMaxLifetime(time.Minute * 5)

	DB = db
}

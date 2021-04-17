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

	// Initialise a new connection pool
	db, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("An error occured while connecting to database: ", err)
	} else {
		log.Printf("Connected to database %s\n", dbname)
	}

	// Set the maximum number of concurrently idle connections to 10. Setting this
	// to less than or equal to 0 will mean that no idle connections are retained.
	db.DB.SetMaxIdleConns(10)

	// Set the maximum number of concurrently open connections (in-use + idle) to 10.
	// Setting this to less than or equal to 0 will mean there is no
	// maximum limit (which is also the default setting).
	db.DB.SetMaxOpenConns(10)

	// Set the maximum lifetime of a connection to 5 minute. Setting it to 0
	// means that there is no maximum lifetime and the connection is reused
	// forever (which is the default behavior).
	db.DB.SetConnMaxLifetime(time.Minute * 5)

	DB = db
}

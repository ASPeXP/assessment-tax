package postgres

import (
	"database/sql"
	"log"
	"os"

	// "github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type Postgres struct {
	Db *sql.DB 
}

func New() (*Postgres, error) {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
		return nil, err 
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err )
	}
	return &Postgres{Db: db}, nil 
}
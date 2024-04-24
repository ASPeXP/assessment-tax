package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Postgres struct {
	Db *sql.DB 
}

func New() (*Postgres, error) {
	databaseSource := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", os.Getenv("host"), os.Getenv("db_port"), os.Getenv("ADMIN_USERNAME"), os.Getenv("ADMIN_PASSWORD"), os.Getenv("database_name"))
	
	db, err := sql.Open("postgres", databaseSource)
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
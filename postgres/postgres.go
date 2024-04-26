package postgres

import (
	"database/sql"
	"fmt"
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
	// connStr := "postgres://"+os.Getenv("ADMIN_USERNAME")+":"+os.Getenv("ADMIN_PASSWORD")+"@"+os.Getenv("host")+":"+os.Getenv("db_port")+"/"+os.Getenv("database_name")+"?sslmode=disable" 
	// db, err := sql.Open("postgres", connStr)
	fmt.Println(os.Getenv("DATABASE_URL"))
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
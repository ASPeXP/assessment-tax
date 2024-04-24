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
	// connStr := fmt.Sprintf("host=%s port=%s user=%s "+
		// "password=%s dbname=%s sslmode=disable", os.Getenv("host"), os.Getenv("db_port"), os.Getenv("ADMIN_USERNAME"), os.Getenv("ADMIN_PASSWORD"), os.Getenv("database_name"))
	// connStr := "postgres://adminTax:admin!@localhost:5432/ktaxes?sslmode=disable"
	connStr := "postgres://"+os.Getenv("ADMIN_USERNAME")+":"+os.Getenv("ADMIN_PASSWORD")+"@"+os.Getenv("host")+":"+os.Getenv("db_port")+"/"+os.Getenv("database_name")+"?sslmode=disable" 
	fmt.Println(connStr)
	db, err := sql.Open("postgres", connStr)
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
package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)


func TestPOSTPersonalDeduction(t *testing.T) {
		amount := 10001.0
		userName := os.Getenv("ADMIN_USERNAME")
		pwd := os.Getenv("ADMIN_PASSWORD")
		host := os.Getenv("host")
		db_port := os.Getenv("db_port")
		db_name := os.Getenv("database_name")
		
		connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", userName, pwd, host, db_port, db_name)
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			// Handle the error
			panic(err)
		}
		defer db.Close()
		
		err = db.Ping()
		if err != nil {
			log.Fatal(err )
		}
		p := &Postgres{Db: db}

	exp := amount
		
	deduction, err := p.GetDeduction()
	if err != nil {
		log.Fatal(err )
	}
	if deduction == (Deduction{}){
		// return "no deduction available"
		deduction, err = p.PostDeduction(amount)
		if err != nil {
			log.Fatal(err )
		}
	}else{
		deduction, err = p.UpdateDeduction(amount)
		if err != nil {
			log.Fatal(err )
		}
	}
		
		got := deduction.Amount
	
	if exp != got {
		t.Errorf("expect %f but got %f", exp, got)
	}

	
}
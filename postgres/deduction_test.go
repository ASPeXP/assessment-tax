package postgres

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)


func TestPOSTPersonalDeduction(t *testing.T) {
		amount := 60000.0
		// userName := os.Getenv("ADMIN_USERNAME")
		// pwd := os.Getenv("ADMIN_PASSWORD")
		// host := os.Getenv("host")
		// db_port := os.Getenv("db_port")
		// db_name := os.Getenv("database_name")
		
		// connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", userName, pwd, host, db_port, db_name)
		connStr := os.Getenv("DATABASE_URL")
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
		
	deduction, err := p.GetPersonalDeduction()
	if err != nil {
		log.Fatal(err )
	}
	if deduction == (Deduction{}){
		// return "no deduction available"
		deduction, err = p.PostPersonalDeduction(amount)
		if err != nil {
			log.Fatal(err )
		}
	}else{
		deduction, err = p.UpdatePersonalDeduction(amount)
		if err != nil {
			log.Fatal(err )
		}
	}
		
		got := deduction.Amount
	
	if exp != got {
		t.Errorf("expect %f but got %f", exp, got)
	}

	
}

func TestPOSTKReceiptDeduction(t *testing.T) {
		amount := 70000.0
		// userName := os.Getenv("ADMIN_USERNAME")
		// pwd := os.Getenv("ADMIN_PASSWORD")
		// host := os.Getenv("host")
		// db_port := os.Getenv("db_port")
		// db_name := os.Getenv("database_name")
		
		// connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", userName, pwd, host, db_port, db_name)
		connStr := os.Getenv("DATABASE_URL")
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
		
	deduction, err := p.GetKReceiptDeduction()
	if err != nil {
		log.Fatal(err )
	}
	if deduction == (Deduction{}){
		// return "no deduction available"
		deduction, err = p.PostKReceiptDeduction(amount)
		if err != nil {
			log.Fatal(err )
		}
	}else{
		deduction, err = p.UpdateKReceiptDeduction(amount)
		if err != nil {
			log.Fatal(err )
		}
	}
		
		got := deduction.Amount
	
	if exp != got {
		t.Errorf("expect %f but got %f", exp, got)
	}

	
}
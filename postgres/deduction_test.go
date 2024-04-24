package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
)

// func PersonalDeduct(amount float64) string {
// 	return fmt.Sprintf(`{"personalDeduction": %.1f}`, amount )
// }
// func TestPOSTPersonal(t *testing.T) {

// 	t.Run("given personal deduct amount 10001 then can get return from database 10001", func(t *testing.T) {

// 		amount := 10001.0

// 		p, err := New()
// 		if err != nil {
// 			panic(err)
// 		}

// 		_, err = p.PostDeduction(amount)
// 		if err != nil {
// 			panic(err)
// 		}

// 		want := amount
// 		pd, err := p.GetDeduction()
// 		if err != nil {
// 			panic(err)
// 		}

// 		got := pd.Amount
// 		assert.Equal(t, want, got, "It's should return 10001.0 when I give 10001.0")

// 	})

// }
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
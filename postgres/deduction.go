package postgres

import (
	"time"
	// "github.com/aspexp/assessment-tax/postgres"
)
type Deduction struct {
	ID int64 `postgres:"id"`
	Amount float64 `postgres:"amount"`
	Created_at time.Time `postgres:"created_at"`
}
func (p *Postgres) GetDeduction() (Deduction, error ){
	
	rows, err := p.Db.Query("select * from personal_deduct where id=1")
	// dd := Deduction{}
	if err != nil {
		return Deduction{}, err 
	}
	defer rows.Close()
	var dd Deduction
	for rows.Next() {
		err = rows.Scan(&dd.ID, &dd.Amount, &dd.Created_at)
		if err != nil {
			return dd, err 
		}
	}
	
	return dd, err 
}
func (p *Postgres) PostDeduction(amount float64) (Deduction, error){

	created_time := time.Now()
	result, err := p.Db.Exec(`Insert into personal_deduct(id, amount, created_at) values(1, $1, $2)`, amount, created_time)
	var dd Deduction
	if err != nil {
		return dd, err 
	}
	rowAffected, err  := result.RowsAffected()
	if err != nil {
		return dd, err 
	}
	if rowAffected > 0 {
		dd.Amount = amount 
		dd.ID = 1
		dd.Created_at = created_time
	}
	
	return dd, nil 
}
func (p *Postgres) UpdateDeduction(amount float64) (Deduction, error){

	created_time := time.Now()
	result, err := p.Db.Exec(`Update personal_deduct set amount=$1, created_at=$2 where id=1`, amount, created_time)
	var dd Deduction
	if err != nil {
		return dd, err 
	}
	rowAffected, err  := result.RowsAffected()
	if err != nil {
		return dd, err 
	}
	if rowAffected > 0 {
		dd.Amount = amount 
		dd.ID = 1
		dd.Created_at = created_time
	}
	
	return dd, nil 
}

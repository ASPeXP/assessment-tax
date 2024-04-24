package postgres

import (
	"fmt"
	"time"
)
type Deduction struct {
	ID int `postgres:"id"`
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
	// for rows.Next() {
		var dd Deduction
		err = rows.Scan(&dd.ID, &dd.Amount, &dd.Created_at)
		if err != nil {
			return dd, err 
		}
	// }
	
	return dd, err 
}
func (p *Postgres) PostDeduction(amount float64) (Deduction, error){

	dd, err := p.GetDeduction()
	if err != nil {
		return Deduction{}, err 
	}
	fmt.Println(dd)
	return dd, nil 
	// sqlStr := fmt.Sprintf(`Insert into personal_deduct(amount, created_at) values(%.1f, %s)`, amount, time.Now())
	// result, err := p.Db.Exec(sqlStr)
	// dd := Deduction{}
	// if err != nil {
	// 	return dd, err 
	// }
	
}


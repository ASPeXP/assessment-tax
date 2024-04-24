package tax

import (
	"fmt"

	"github.com/aspexp/assessment-tax/postgres"
)


func TaxHandler(pmi PersonalTaxInfo) string {
	// return CalTaxPTI(pmi)
	return CalTaxPTITaxLevel(pmi)

}
func InsertPersonalDeduct(amount float64) string {

	p, err := postgres.New()
	if err != nil {
		panic(err)
	}

	deduction, err := p.GetDeduction()
	if err != nil {
		return "error:" + err.Error()
	}
	if deduction == (postgres.Deduction{}){
		// return "no deduction available"
		deduction, err = p.PostDeduction(amount)
		if err != nil {
			return "error:" + err.Error()
		}
	}else{
		deduction, err = p.UpdateDeduction(amount)
		if err != nil {
			return "error:" + err.Error()
		}
	}
	return fmt.Sprintf(`{"personalDeduction": %.1f }`, deduction.Amount)
}
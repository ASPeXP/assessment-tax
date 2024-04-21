package tax

import (
	"fmt"
)

type PersonalTaxInfo struct {
	Income           float64
	Wht              float64
	PersonalDeducted float64
}

func CalTax(amount float64, wht float64, personalDeducted float64) string {

	amount = (amount - personalDeducted) 
	taxVal := 0.00 
	if amount > 2000000  {
		amount = amount - 2000000
		tax := amount * 0.35
		taxVal += tax

		amount = 2000000 - 1000000
		tax = amount * 0.20
		taxVal += tax

		amount = 1000000 - 500000
		tax = amount * 0.15
		taxVal += tax

		amount = 500000-150000
		tax = amount * 0.10
		taxVal += tax 
		taxVal -= wht
		return fmt.Sprintf(`{"tax": %.2f }`, taxVal)
	}

	if amount > 1000000 && amount <= 2000000 {
		amount = amount - 1000000
		tax := amount * 0.20
		taxVal += tax

		amount = 1000000 - 500000
		tax = amount * 0.15
		taxVal += tax

		amount = 500000-150000
		tax = amount * 0.10
		taxVal += tax 
		taxVal -= wht
		return fmt.Sprintf(`{"tax": %.2f }`, taxVal)
	}
	if amount > 500000 && amount <= 1000000 {
		amount = amount - 500000
		tax := amount * 0.15
		taxVal += tax

		amount = 500000-150000
		tax = amount * 0.10
		taxVal += tax 
		taxVal -= wht
		return fmt.Sprintf(`{"tax": %.2f }`, taxVal)
	}
	if amount > 150000 && amount <= 500000 {
		amount = amount - 150000
		tax := amount * 0.10
		taxVal += tax
		taxVal -= wht
		return fmt.Sprintf(`{"tax": %.2f }`, taxVal)
	}
	


	return fmt.Sprintf(`{"tax": %.2f }`, taxVal)
}
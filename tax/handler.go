package tax

func TaxHandler(amount float64, wht float64, personalDeducted float64) string {
	return CalTax(amount, wht, personalDeducted)

}
package tax

func TaxHandler(pmi PersonalTaxInfo) string {
	// return CalTaxPTI(pmi)
	return CalTaxPTITaxLevel(pmi)

}
func InsertPersonalDeduct(amount float64) string {
	return ""
}
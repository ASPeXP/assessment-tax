package tax

import (
	"fmt"
	// "os"
	// "path/filepath"
	"strconv"
	"strings"

	"github.com/aspexp/assessment-tax/postgres"
)


func TaxHandler(pti PersonalTaxInfo) string {
	return CalTaxPTITaxLevel(pti)
}
func InsertPersonalDeduct(amount float64) string {

	p, err := postgres.New()
	if err != nil {
		panic(err)
	}

	deduction, err := p.GetPersonalDeduction()
	if err != nil {
		return "error:" + err.Error()
	}
	if amount > 100000 {
		amount = 100000
	}
	if amount <= 10000 {
		return "error: Amount must greater than 10000 THB."  
	}
	if deduction == (postgres.Deduction{}){
		// return "no deduction available"
		deduction, err = p.PostPersonalDeduction(amount)
		if err != nil {
			return "error:" + err.Error()
		}
	}else{
		deduction, err = p.UpdatePersonalDeduction(amount)
		if err != nil {
			return "error:" + err.Error()
		}
	}
	return fmt.Sprintf(`{"personalDeduction": %.1f }`, deduction.Amount)
}
func InsertKReceiptDeduct(amount float64) string {

	p, err := postgres.New()
	if err != nil {
		panic(err)
	}

	deduction, err := p.GetKReceiptDeduction()
	if err != nil {
		return "error:" + err.Error()
	}
	if amount > 100000 {
		amount = 100000
	}
	if amount <= 0 {
		return "error: Amount must greater than 0 THB."  
	}
	if deduction == (postgres.Deduction{}){
		// return "no deduction available"
		deduction, err = p.PostKReceiptDeduction(amount)
		if err != nil {
			return "error:" + err.Error()
		}
	}else{
		deduction, err = p.UpdateKReceiptDeduction(amount)
		if err != nil {
			return "error:" + err.Error()
		}
	}
	return fmt.Sprintf(`{"kReceipt": %.1f }`, deduction.Amount)
}

func GetTaxCSV(filePath string ) string {
	text := ReadCSV(filePath)
	if text == nil {
		return "provided csv is empty."
		
	}
	var bill string 
	for i, each_ln := range text {
		if i == 0 {
			continue
		}

		if !strings.Contains(each_ln, ","){
			return "provided csv not in correct format."
		}

		line_data := strings.Split(each_ln, ",")

		income, err := strconv.ParseFloat(line_data[0], 64)
		if err != nil {
			return "error:" + err.Error()
		}
		wht, err := strconv.ParseFloat(line_data[1], 64)
		if err != nil {
			return "error:" + err.Error()
		}
		donation, err := strconv.ParseFloat(line_data[2], 64)
		if err != nil {
			return "error:" + err.Error()
		}
		var pti = PersonalTaxInfo{
			Income:           income,
			Wht:              wht,
			PersonalDeducted: 60000.0,
			Donation:         donation,
		}
		bill += CalTaxPTI(pti)
	}
	bill = bill[:len(bill)-1]
	return fmt.Sprintf(`{"taxes": [%s]}`, bill)
}
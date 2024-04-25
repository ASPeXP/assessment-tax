package tax

import (
	"fmt"
	// "os"
	// "path/filepath"
	"strconv"
	"strings"

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
// func getCSVPath(c echo.Context, fileName string ) (string, error){
// 	root, err := os.Getwd()
// 	if err != nil {
// 		return "", err 
// 	}

// 	csvPath := filepath.Join(root, fileName)
// 	mappedPath, err := c.MapPath(csvPath)
// 	if err != nil {
// 		return "", err 
// 	}

// 	return mappedPath, nil 
// }
func GetTaxCSV(filePath string ) string {
	text := ReadCSV(filePath)
	var bill string 
	for i, each_ln := range text {
		if i == 0 {
			continue
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
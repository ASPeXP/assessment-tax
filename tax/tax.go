package tax

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
)

type PersonalTaxInfo struct {
	Income           float64
	Wht              float64
	PersonalDeducted float64
	Donation         float64
	KReceipt float64
}

type TaxData struct {
	Tax float64 `json:"tax"`
	TaxLevel []TaxLevel `json:"taxLevel"`
}
type TaxLevel struct {
	Level string `json:"level"`
	Tax float64	`json:"tax"`
}

type TaxDataCSV struct {
	TotalIncome float64 `json:"totalIncome"`
	Tax float64 `json:"tax"`
}
func getCSVPath(filename string) (string, error){
	root, err := os.Getwd()
	if err != nil {
		return "", err 
	}

	csvPath := filepath.Join(root, filename)
	return csvPath, nil 
}
func ReadCSV(filePath string ) []string {
		csvPath, err := getCSVPath(filePath)
		if err != nil {
			log.Printf("Failed to get CSV file path: %v", err )
		}

		file, err := os.Open(csvPath)
		if err != nil {
			log.Fatalf("Error: failed to open %v", err )
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanLines)
		var text []string 
		for scanner.Scan(){
			text = append(text, scanner.Text())
		}
	
		file.Close()
	
	return text
}
func getLevelTax(tlv []TaxLevel) string {
	var s string 
	
	for _, t := range tlv {
		s += fmt.Sprintf(`{"level": "%s","tax": %.1f},`,  t.Level, t.Tax)
	}

	return 	s[:len(s)-1]
	// return s
	
}
// func getTaxCSV(tdc []TaxDataCSV) string {
// 	var s string 
	
// 	for _, t := range tdc {
// 		s += fmt.Sprintf(`{"level": "%.1f","tax": %.1f},`,  t.TotalIncome, t.Tax)
// 	}

// 	return 	s[:len(s)-1]
// }

func CalTaxPTITaxLevel(pti PersonalTaxInfo) string {
	
	var tlv []TaxLevel
	var bodyStr []string 
	taxZero := 0.0
	pti.Income = (pti.Income - pti.PersonalDeducted) - pti.Donation - pti.KReceipt
	taxVal := 0.0
	
	if pti.Income > 2000000 {
		
		pti.Income = pti.Income - 2000000
		tax := pti.Income * 0.35
		taxVal += tax
		tl := TaxLevel{
			Level: "2,000,001 ขึ้นไป",
			Tax: tax,
		}
		tlv = append(tlv, tl)

		pti.Income = 2000000 - 1000000
		tax = pti.Income * 0.20
		taxVal += tax
		tl = TaxLevel{
			Level: "1,000,001-2,000,000",
			Tax: tax,
		}
		tlv = append(tlv, tl)

		pti.Income = 1000000 - 500000
		tax = pti.Income * 0.15
		taxVal += tax
		tl = TaxLevel{
			Level: "500,001-1,000,000",
			Tax: tax,
		}
		tlv = append(tlv, tl)

		pti.Income = 500000 - 150000
		tax = pti.Income * 0.10
		taxVal += tax
		taxVal -= pti.Wht
		tl = TaxLevel{
			Level: "150,001-500,000",
			Tax: tax,
		}
		tlv = append(tlv, tl)

		tl = TaxLevel{
			Level: "0-150,000",
			Tax: 0.0,
		}
		tlv = append(tlv, tl)
		

		// bodyStr += getLevelTax(tlv)
		bodyStr = append(bodyStr, getLevelTax(tlv))
		return fmt.Sprintf(`{"tax": %.1f,"taxLevel": %s}`, taxVal, bodyStr)
	}

	if pti.Income > 1000000 && pti.Income <= 2000000 {
		

		tl := TaxLevel{
			Level: "2,000,001 ขึ้นไป",
			Tax: 0.0,
		}
		tlv = append(tlv, tl)

		pti.Income = pti.Income - 1000000
		tax := pti.Income * 0.20
		taxVal += tax
		tl = TaxLevel{
			Level: "1,000,001-2,000,000",
			Tax: tax,
		}
		tlv = append(tlv, tl)

		pti.Income = 1000000 - 500000
		tax = pti.Income * 0.15
		taxVal += tax
		tl = TaxLevel{
			Level: "500,001-1,000,000",
			Tax: tax,
		}
		tlv = append(tlv, tl)

		pti.Income = 500000 - 150000
		tax = pti.Income * 0.10
		taxVal += tax
		taxVal -= pti.Wht
		tl = TaxLevel{
			Level: "150,001-500,000",
			Tax: tax,
		}
		tlv = append(tlv, tl)

		tl = TaxLevel{
			Level: "0-150,000",
			Tax: 0.0,
		}
		
		tlv = append(tlv, tl)
		
		// bodyStr += getLevelTax(tlv)
		bodyStr = append(bodyStr, getLevelTax(tlv))
		return fmt.Sprintf(`{"tax": %.1f,"taxLevel": %s}`, taxVal, bodyStr)
	}
	if pti.Income > 500000 && pti.Income <= 1000000 {

		tl := TaxLevel{
			Level: "2,000,001 ขึ้นไป",
			Tax: 0.0,
		}
		tlv = append(tlv, tl)

		tl = TaxLevel{
			Level: "1,000,001-2,000,000",
			Tax: 0.0,
		}
		tlv = append(tlv, tl)

		pti.Income = pti.Income - 500000
		tax := pti.Income * 0.15
		taxVal += tax
		tl = TaxLevel{
			Level: "500,001-1,000,000",
			Tax: tax,
		}
		tlv = append(tlv, tl)

		pti.Income = 500000 - 150000
		tax = pti.Income * 0.10
		taxVal += tax
		taxVal -= pti.Wht
		tl = TaxLevel{
			Level: "150,001-500,000",
			Tax: tax,
		}
		tlv = append(tlv, tl)

		
		tl = TaxLevel{
			Level: "0-150,000",
			Tax: 0.0,
		}
		tlv = append(tlv, tl)
		

		// bodyStr += getLevelTax(tlv)
		bodyStr = append(bodyStr, getLevelTax(tlv))
		return fmt.Sprintf(`{"tax": %.1f,"taxLevel": %s}`, taxVal, bodyStr)
	}
	if pti.Income > 150000 && pti.Income <= 500000 {
	
		tl := TaxLevel{
			Level: "2,000,001 ขึ้นไป",
			Tax: taxZero,
		}
		tlv = append(tlv, tl)

		tl = TaxLevel{
			Level: "1,000,001-2,000,000",
			Tax: taxZero,
		}
		tlv = append(tlv, tl)

		tl = TaxLevel{
			Level: "500,001-1,000,000",
			Tax: taxZero,
		}
		tlv = append(tlv, tl)

		pti.Income = pti.Income - 150000
		tax := pti.Income * 0.10
		taxVal += tax
		taxVal -= pti.Wht
		tl = TaxLevel{
			Level: "150,001-500,000",
			Tax: tax,
		}
		tlv = append(tlv, tl)

		tl = TaxLevel{
		Level: "0-150,000",
		Tax: taxZero,
		}
		tlv = append(tlv, tl)
		
		// bodyStr += getLevelTax(tlv)
		bodyStr = append(bodyStr, getLevelTax(tlv))
		return fmt.Sprintf(`{"tax": %.1f,"taxLevel": %s}`, taxVal, bodyStr)
	}
	if pti.Income <= 150000  {
		
		tax := 0.0
		taxVal += tax
		tl := TaxLevel{
			Level: "0-150,000",
			Tax: taxVal,
		}
		tlv = append(tlv, tl)
		// bodyStr += getLevelTax(tlv)
		bodyStr = append(bodyStr, getLevelTax(tlv))
		return fmt.Sprintf(`{"tax": %.1f,"taxLevel": %s}`, taxVal, bodyStr)
	}


	return fmt.Sprintf(`{"tax": %.1f,"taxLevel": %s}`, taxVal, bodyStr)
}

func CalTaxPTI(pti PersonalTaxInfo) string {

	totalIncome := pti.Income
	pti.Income = (pti.Income - pti.PersonalDeducted) - pti.Donation - pti.KReceipt
	
	taxVal := 0.00
	if pti.Income > 2000000 {
		pti.Income = pti.Income - 2000000
		tax := pti.Income * 0.35
		taxVal += tax

		pti.Income = 2000000 - 1000000
		tax = pti.Income * 0.20
		taxVal += tax

		pti.Income = 1000000 - 500000
		tax = pti.Income * 0.15
		taxVal += tax

		pti.Income = 500000 - 150000
		tax = pti.Income * 0.10
		taxVal += tax
		taxVal -= pti.Wht
		var retStr string 
		if taxVal >= 0 {
			retStr = fmt.Sprintf(`{"totalIncome": %.1f,"tax": %.1f},`, totalIncome, taxVal)
		}
		if taxVal < 0 {
			retStr = fmt.Sprintf(`{"totalIncome": %.1f,"tax": %.1f, "taxRefund": %.1f},`, totalIncome, 0.0, math.Abs(taxVal))
		}

		return retStr
	}

	if pti.Income > 1000000 && pti.Income <= 2000000 {
		pti.Income = pti.Income - 1000000
		tax := pti.Income * 0.20
		taxVal += tax

		pti.Income = 1000000 - 500000
		tax = pti.Income * 0.15
		taxVal += tax

		pti.Income = 500000 - 150000
		tax = pti.Income * 0.10
		taxVal += tax
		taxVal -= pti.Wht
		var retStr string 
		if taxVal >= 0 {
			retStr = fmt.Sprintf(`{"totalIncome": %.1f,"tax": %.1f},`, totalIncome, taxVal)
		}
		if taxVal < 0 {
			retStr = fmt.Sprintf(`{"totalIncome": %.1f,"tax": %.1f, "taxRefund": %.1f},`, totalIncome, 0.0, math.Abs(taxVal))
		}

		return retStr
	}
	if pti.Income > 500000 && pti.Income <= 1000000 {
		pti.Income = pti.Income - 500000
		tax := pti.Income * 0.15
		taxVal += tax

		pti.Income = 500000 - 150000
		tax = pti.Income * 0.10
		taxVal += tax
		taxVal -= pti.Wht
		var retStr string 
		if taxVal >= 0 {
			retStr = fmt.Sprintf(`{"totalIncome": %.1f,"tax": %.1f},`, totalIncome, taxVal)
		}
		if taxVal < 0 {
			retStr = fmt.Sprintf(`{"totalIncome": %.1f,"tax": %.1f, "taxRefund": %.1f},`, totalIncome, 0.0, math.Abs(taxVal))
		}

		return retStr
	}
	if pti.Income > 150000 && pti.Income <= 500000 {
		pti.Income = pti.Income - 150000
		tax := pti.Income * 0.10
		taxVal += tax
		taxVal -= pti.Wht
		
		var retStr string 
		if taxVal >= 0 {
			retStr = fmt.Sprintf(`{"totalIncome": %.1f,"tax": %.1f},`, totalIncome, taxVal)
		}
		if taxVal < 0 {
			retStr = fmt.Sprintf(`{"totalIncome": %.1f,"tax": %.1f, "taxRefund": %.1f},`, totalIncome, 0.0, math.Abs(taxVal))
		}

		return retStr
	}

	// return fmt.Sprintf(`{"totalIncome": %.1f,"tax": %.1f },`, totalIncome, taxVal)
	var retStr string 
		if taxVal >= 0 {
			retStr = fmt.Sprintf(`{"totalIncome": %.1f,"tax": %.1f},`, totalIncome, taxVal)
		}
		if taxVal < 0 {
			retStr = fmt.Sprintf(`{"totalIncome": %.1f,"tax": %.1f, "taxRefund": %.1f},`, totalIncome, 0.0, math.Abs(taxVal))
		}

	return retStr
}

// func CalTax(amount float64, wht float64, personalDeducted float64) string {

// 	amount = (amount - personalDeducted)
// 	taxVal := 0.00
// 	if amount > 2000000 {
// 		amount = amount - 2000000
// 		tax := amount * 0.35
// 		taxVal += tax

// 		amount = 2000000 - 1000000
// 		tax = amount * 0.20
// 		taxVal += tax

// 		amount = 1000000 - 500000
// 		tax = amount * 0.15
// 		taxVal += tax

// 		amount = 500000 - 150000
// 		tax = amount * 0.10
// 		taxVal += tax
// 		taxVal -= wht
// 		return fmt.Sprintf(`{"tax": %.2f }`, taxVal)
// 	}

// 	if amount > 1000000 && amount <= 2000000 {
// 		amount = amount - 1000000
// 		tax := amount * 0.20
// 		taxVal += tax

// 		amount = 1000000 - 500000
// 		tax = amount * 0.15
// 		taxVal += tax

// 		amount = 500000 - 150000
// 		tax = amount * 0.10
// 		taxVal += tax
// 		taxVal -= wht
// 		return fmt.Sprintf(`{"tax": %.2f }`, taxVal)
// 	}
// 	if amount > 500000 && amount <= 1000000 {
// 		amount = amount - 500000
// 		tax := amount * 0.15
// 		taxVal += tax

// 		amount = 500000 - 150000
// 		tax = amount * 0.10
// 		taxVal += tax
// 		taxVal -= wht
// 		return fmt.Sprintf(`{"tax": %.2f }`, taxVal)
// 	}
// 	if amount > 150000 && amount <= 500000 {
// 		amount = amount - 150000
// 		tax := amount * 0.10
// 		taxVal += tax
// 		taxVal -= wht
// 		return fmt.Sprintf(`{"tax": %.2f }`, taxVal)
// 	}

// 	return fmt.Sprintf(`{"tax": %.2f }`, taxVal)
// }

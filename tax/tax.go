package tax

import (
	"fmt"
)

type PersonalTaxInfo struct {
	Income           float64
	Wht              float64
	PersonalDeducted float64
	Donation         float64
}

type TaxLevel struct {
	level string 
	tax float64
}

func getLevelTax(tlv []TaxLevel) string {
	var s string 
	
	for _, t := range tlv {
		s += fmt.Sprintf(`{"level": "%s","tax": %.1f},`,  t.level, t.tax)
		// s += fmt.Sprintf(`{"level": "%s","tax": %.1f}`,  t.level, t.tax)
	}

	return 	s[:len(s)-1]
	// return s
	
}

func CalTaxPTITaxLevel(pti PersonalTaxInfo) string {
	
	var tlv []TaxLevel
	var bodyStr []string 
	taxZero := 0.0

	pti.Income = (pti.Income - pti.PersonalDeducted) - pti.Donation
	taxVal := 0.0
	
	if pti.Income > 2000000 {
		
		pti.Income = pti.Income - 2000000
		tax := pti.Income * 0.35
		taxVal += tax
		tl := TaxLevel{
			level: "2,000,001 ขึ้นไป",
			tax: tax,
		}
		tlv = append(tlv, tl)

		pti.Income = 2000000 - 1000000
		tax = pti.Income * 0.20
		taxVal += tax
		tl = TaxLevel{
			level: "1,000,001-2,000,000",
			tax: tax,
		}
		tlv = append(tlv, tl)

		pti.Income = 1000000 - 500000
		tax = pti.Income * 0.15
		taxVal += tax
		tl = TaxLevel{
			level: "500,001-1,000,000",
			tax: tax,
		}
		tlv = append(tlv, tl)

		pti.Income = 500000 - 150000
		tax = pti.Income * 0.10
		taxVal += tax
		tl = TaxLevel{
			level: "150,001-500,000",
			tax: tax,
		}
		tlv = append(tlv, tl)

		tl = TaxLevel{
			level: "0-150,000",
			tax: 0.0,
		}
		tlv = append(tlv, tl)
		
		taxVal -= pti.Wht

		// bodyStr += getLevelTax(tlv)
		bodyStr = append(bodyStr, getLevelTax(tlv))
		return fmt.Sprintf(`{"tax": %.1f,"taxLevel": %s}`, taxVal, bodyStr)
	}

	if pti.Income > 1000000 && pti.Income <= 2000000 {
		

		tl := TaxLevel{
			level: "2,000,001 ขึ้นไป",
			tax: 0.0,
		}
		tlv = append(tlv, tl)

		pti.Income = pti.Income - 1000000
		tax := pti.Income * 0.20
		taxVal += tax
		tl = TaxLevel{
			level: "1,000,001-2,000,000",
			tax: tax,
		}
		tlv = append(tlv, tl)

		pti.Income = 1000000 - 500000
		tax = pti.Income * 0.15
		taxVal += tax
		tl = TaxLevel{
			level: "500,001-1,000,000",
			tax: tax,
		}
		tlv = append(tlv, tl)

		pti.Income = 500000 - 150000
		tax = pti.Income * 0.10
		taxVal += tax
		taxVal -= pti.Wht
		tl = TaxLevel{
			level: "150,001-500,000",
			tax: tax,
		}
		tlv = append(tlv, tl)

		tl = TaxLevel{
			level: "0-150,000",
			tax: 0.0,
		}
		tlv = append(tlv, tl)
		// bodyStr += getLevelTax(tlv)
		bodyStr = append(bodyStr, getLevelTax(tlv))
		return fmt.Sprintf(`{"tax": %.1f,"taxLevel": %s}`, taxVal, bodyStr)
	}
	if pti.Income > 500000 && pti.Income <= 1000000 {

		tl := TaxLevel{
			level: "2,000,001 ขึ้นไป",
			tax: 0.0,
		}
		tlv = append(tlv, tl)

		tl = TaxLevel{
			level: "1,000,001-2,000,000",
			tax: 0.0,
		}
		tlv = append(tlv, tl)

		// tl = TaxLevel{
		// 	level: "1,000,001-2,000,000",
		// 	tax: 0.0,
		// }
		// tlv = append(tlv, tl)

		pti.Income = pti.Income - 500000
		tax := pti.Income * 0.15
		taxVal += tax
		tl = TaxLevel{
			level: "500,001-1,000,000",
			tax: tax,
		}
		tlv = append(tlv, tl)

		pti.Income = 500000 - 150000
		tax = pti.Income * 0.10
		taxVal += tax
		taxVal -= pti.Wht
		tl = TaxLevel{
			level: "150,001-500,000",
			tax: tax,
		}
		tlv = append(tlv, tl)

		
		tl = TaxLevel{
			level: "0-150,000",
			tax: 0.0,
		}
		tlv = append(tlv, tl)
		

		// bodyStr += getLevelTax(tlv)
		bodyStr = append(bodyStr, getLevelTax(tlv))
		return fmt.Sprintf(`{"11tax": %.1f,"taxLevel": %s}`, taxVal, bodyStr)
	}
	if pti.Income > 150000 && pti.Income <= 500000 {
	
		tl := TaxLevel{
			level: "2,000,001 ขึ้นไป",
			tax: taxZero,
		}
		tlv = append(tlv, tl)

		tl = TaxLevel{
			level: "1,000,001-2,000,000",
			tax: taxZero,
		}
		tlv = append(tlv, tl)

		tl = TaxLevel{
			level: "500,001-1,000,000",
			tax: taxZero,
		}
		tlv = append(tlv, tl)

		pti.Income = pti.Income - 150000
		tax := pti.Income * 0.10
		taxVal += tax
		taxVal -= pti.Wht
		tl = TaxLevel{
			level: "150,001-500,000",
			tax: tax,
		}
		tlv = append(tlv, tl)

		tl = TaxLevel{
		level: "0-150,000",
		tax: taxZero,
		}
		tlv = append(tlv, tl)
		
		// bodyStr += getLevelTax(tlv)
		bodyStr = append(bodyStr, getLevelTax(tlv))
		return fmt.Sprintf(`{"tax": %.1f,"taxLevel": %s}`, taxVal, bodyStr)
	}
	if pti.Income <= 150000  {
		
		tax := 0.0
		taxVal += tax
		// taxVal -= pti.Wht
		tl := TaxLevel{
			level: "0-150,000",
			tax: tax,
		}
		tlv = append(tlv, tl)
		// bodyStr += getLevelTax(tlv)
		bodyStr = append(bodyStr, getLevelTax(tlv))
		return fmt.Sprintf(`{"tax": %.1f,"taxLevel": %s}`, taxVal, bodyStr)
	}


	return fmt.Sprintf(`{"tax": %.1f,"taxLevel": %s}`, taxVal, bodyStr)
}

func CalTaxPTI(pti PersonalTaxInfo) string {

	pti.Income = (pti.Income - pti.PersonalDeducted) - pti.Donation
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
		return fmt.Sprintf(`{"tax": %.2f }`, taxVal)
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
		return fmt.Sprintf(`{"tax": %.2f }`, taxVal)
	}
	if pti.Income > 500000 && pti.Income <= 1000000 {
		pti.Income = pti.Income - 500000
		tax := pti.Income * 0.15
		taxVal += tax

		pti.Income = 500000 - 150000
		tax = pti.Income * 0.10
		taxVal += tax
		taxVal -= pti.Wht
		return fmt.Sprintf(`{"tax": %.2f }`, taxVal)
	}
	if pti.Income > 150000 && pti.Income <= 500000 {
		pti.Income = pti.Income - 150000
		tax := pti.Income * 0.10
		taxVal += tax
		taxVal -= pti.Wht
		return fmt.Sprintf(`{"tax": %.2f }`, taxVal)
	}

	return fmt.Sprintf(`{"tax": %.2f }`, taxVal)
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

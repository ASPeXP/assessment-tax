package tax

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPOSTPersonal(t *testing.T) {

	t.Run("given amount not more than 150000 should return calculate tax as zero", func(t *testing.T) {
		amount := 70000.0
		want := amount
		pd := PersonalDeduct(amount)
		var b map[string]float64
		err := json.Unmarshal([]byte(pd), &b)
		if err != nil {
			fmt.Println("error:", err)
		}
		
		got := b["personalDeduction"]
		assert.Equal(t, want, got, "It's should return 70000.0 when I give 70000.0")
	})
	
}
func TestTax(t *testing.T) {

	t.Run("given amount not more than 150000 should return calculate tax as zero", func(t *testing.T) {
		// amount := 149999.00
		// wht := 0.0
		// personalDeducted := 60000.00
		var pti = PersonalTaxInfo{
			Income:           149999.00,
			Wht:              0.00,
			PersonalDeducted: 60000.00,
			Donation:         0.00,
		}
		want := 0.0

		tax := CalTaxPTITaxLevel(pti)
		var taxData TaxData
		err := json.Unmarshal([]byte(tax), &taxData)
		if err != nil {
			fmt.Println("error:", err)
		}
		got := taxData.Tax
		// tax := CalTax(amount, wht, personalDeducted)
		// tax := CalTaxPTI(pti)
		// var b map[string]float64
		// err := json.Unmarshal([]byte(tax), &b)
		// if err != nil {
		// 	fmt.Println("error:", err)
		// }

		// got := b["tax"]
		assert.Equal(t, want, got, "It's should return 0 when I give 149999")
	})

	t.Run("given amount more than 150000 but less than 500002 should return calculate tax as 10 percent", func(t *testing.T) {
		// amount := 500001.00
		// wht := 0.0
		// personalDeducted := 60000.00
		var pti = PersonalTaxInfo{
			Income:           500001.00,
			Wht:              0.00,
			PersonalDeducted: 60000.00,
			Donation:         0.00,
		}

		want := 29000.10

		tax := CalTaxPTITaxLevel(pti)
		var taxData TaxData
		err := json.Unmarshal([]byte(tax), &taxData)
		if err != nil {
			fmt.Println("error:", err)
		}
		got := taxData.Tax
		// tax := CalTax(amount, wht, personalDeducted)
		// tax := CalTaxPTI(pti)
		// var b map[string]float64
		// err := json.Unmarshal([]byte(tax), &b)
		// if err != nil {
		// 	fmt.Println("error:", err)
		// }

		// got := b["tax"]
		assert.Equal(t, want, got, "It's should return 29000.10 when I give 500001.00")
	})

	t.Run("given amount more than 660000 should return calculate tax as 50000", func(t *testing.T) {
		// amount := 660000.00
		// wht := 0.0
		// personalDeducted := 60000.00

		var pti = PersonalTaxInfo{
			Income:           660000.00,
			Wht:              0.00,
			PersonalDeducted: 60000.00,
			Donation:         0.00,
		}

		want := 50000.00

		tax := CalTaxPTITaxLevel(pti)
		var taxData TaxData
		err := json.Unmarshal([]byte(tax), &taxData)
		if err != nil {
			fmt.Println("error:", err)
		}
		got := taxData.Tax
		// tax := CalTax(amount, wht, personalDeducted)
		// tax := CalTaxPTI(pti)
		// var b map[string]float64
		// err := json.Unmarshal([]byte(tax), &b)
		// if err != nil {
		// 	fmt.Println("error:", err)
		// }

		// got := b["tax"]
		assert.Equal(t, want, got, "It's should return 50000.00 when I give 660000.00")
	})

	t.Run("given amount  500000 should return calculate tax as 29000.00", func(t *testing.T) {
		// amount := 500000.00
		// wht := 0.0
		// personalDeducted := 60000.00
		var pti = PersonalTaxInfo{
			Income:           500000.00,
			Wht:              0.00,
			PersonalDeducted: 60000.00,
			Donation:         0.00,
		}

		want := 29000.00

		tax := CalTaxPTITaxLevel(pti)
		var taxData TaxData
		err := json.Unmarshal([]byte(tax), &taxData)
		if err != nil {
			fmt.Println("error:", err)
		}
		got := taxData.Tax
		// tax := CalTax(amount, wht, personalDeducted)
		// tax := CalTaxPTI(pti)
		// var b map[string]float64
		// err := json.Unmarshal([]byte(tax), &b)
		// if err != nil {
		// 	fmt.Println("error:", err)
		// }

		// got := b["tax"]
		assert.Equal(t, want, got, "It's should return 29000.00 when I give 500000.00")
	})

	t.Run("given amount  1000000 should return calculate tax as 15% after deducted", func(t *testing.T) {
		// amount := 1000000.00
		// wht := 0.0
		// personalDeducted := 60000.00

		var pti = PersonalTaxInfo{
			Income:           1000000.00,
			Wht:              0.00,
			PersonalDeducted: 60000.00,
			Donation:         0.00,
		}

		want := 101000.00

		tax := CalTaxPTITaxLevel(pti)
		var taxData TaxData
		err := json.Unmarshal([]byte(tax), &taxData)
		if err != nil {
			fmt.Println("error:", err)
		}
		got := taxData.Tax
		// tax := CalTax(amount, wht, personalDeducted)
		// tax := CalTaxPTI(pti)
		// var b map[string]float64
		// err := json.Unmarshal([]byte(tax), &b)
		// if err != nil {
		// 	fmt.Println("error:", err)
		// }

		// got := b["tax"]
		assert.Equal(t, want, got, "It's should return 101000.00 when I give 1000000.00")
	})
	t.Run("given amount  1900000 should return calculate tax as 20% after deducted", func(t *testing.T) {
		// amount := 1900000.00
		// wht := 0.0
		// personalDeducted := 60000.00

		var pti = PersonalTaxInfo{
			Income:           1900000.00,
			Wht:              0.00,
			PersonalDeducted: 60000.00,
			Donation:         0.00,
		}

		want := 278000.00

		tax := CalTaxPTITaxLevel(pti)
		var taxData TaxData
		err := json.Unmarshal([]byte(tax), &taxData)
		if err != nil {
			fmt.Println("error:", err)
		}
		got := taxData.Tax
		// tax := CalTax(amount, wht, personalDeducted)
		// tax := CalTaxPTI(pti)
		// var b map[string]float64
		// err := json.Unmarshal([]byte(tax), &b)
		// if err != nil {
		// 	fmt.Println("error:", err)
		// }

		// got := b["tax"]
		assert.Equal(t, want, got, "It's should return 278000.00 when I give 1900000.00")
	})

	t.Run("given amount  2100000 should return calculate tax as 20% after deducted", func(t *testing.T) {
		// amount := 2100000.00
		// wht := 0.0
		// personalDeducted := 60000.00
		var pti = PersonalTaxInfo{
			Income:           2100000.00,
			Wht:              0.00,
			PersonalDeducted: 60000.00,
			Donation:         0.00,
		}

		want := 324000.00

		tax := CalTaxPTITaxLevel(pti)
		var taxData TaxData
		err := json.Unmarshal([]byte(tax), &taxData)
		if err != nil {
			fmt.Println("error:", err)
		}
		got := taxData.Tax
		// tax := CalTax(amount, wht, personalDeducted)
		// tax := CalTaxPTI(pti)
		// var b map[string]float64
		// err := json.Unmarshal([]byte(tax), &b)
		// if err != nil {
		// 	fmt.Println("error:", err)
		// }

		// got := b["tax"]
		assert.Equal(t, want, got, "It's should return 324000.00 when I give 2100000.00")
	})

	t.Run("given amount not more than 150000 with wht 25000 should return calculate tax as zero", func(t *testing.T) {
		// amount := 149999.00
		// wht := 25000.00
		// personalDeducted := 60000.00
		var pti = PersonalTaxInfo{
			Income:           149999.00,
			Wht:              25000.00,
			PersonalDeducted: 60000.00,
			Donation:         0.00,
		}

		want := 0.0

		tax := CalTaxPTITaxLevel(pti)
		var taxData TaxData
		err := json.Unmarshal([]byte(tax), &taxData)
		if err != nil {
			fmt.Println("error:", err)
		}
		got := taxData.Tax
		// tax := CalTax(amount, wht, personalDeducted)
		// tax := CalTaxPTI(pti)
		// var b map[string]float64
		// err := json.Unmarshal([]byte(tax), &b)
		// if err != nil {
		// 	fmt.Println("error:", err)
		// }

		// got := b["tax"]
		assert.Equal(t, want, got, "It's should return 0 when I give 149999 with wht 25000")
	})

	t.Run("given amount more than 150000 but less than 500002 with wht 25000 should return calculate tax as 4000.10", func(t *testing.T) {
		// amount := 500001.00
		// wht := 25000.00
		// personalDeducted := 60000.00
		var pti = PersonalTaxInfo{
			Income:           500001.00,
			Wht:              25000.00,
			PersonalDeducted: 60000.00,
			Donation:         0.00,
		}

		want := 4000.10
		tax := CalTaxPTITaxLevel(pti)
		var taxData TaxData
		err := json.Unmarshal([]byte(tax), &taxData)
		if err != nil {
			fmt.Println("error:", err)
		}
		got := taxData.Tax
		// tax := CalTax(amount, wht, personalDeducted)
		// tax := CalTaxPTI(pti)
		// var b map[string]float64
		// err := json.Unmarshal([]byte(tax), &b)
		// if err != nil {
		// 	fmt.Println("error:", err)
		// }

		// got := b["tax"]
		assert.Equal(t, want, got, "It's should return 4000.10 when I give 500001.00")
	})
	t.Run("given amount  500000 with wht 25000 should return calculate tax as 4000.00", func(t *testing.T) {
		// amount := 500000.00
		// wht := 25000.00
		// personalDeducted := 60000.00
		var pti = PersonalTaxInfo{
			Income:           500000.00,
			Wht:              25000.00,
			PersonalDeducted: 60000.00,
			Donation:         0.00,
		}

		want := 4000.00

		tax := CalTaxPTITaxLevel(pti)
		var taxData TaxData
		err := json.Unmarshal([]byte(tax), &taxData)
		if err != nil {
			fmt.Println("error:", err)
		}
		got := taxData.Tax
		// tax := CalTax(amount, wht, personalDeducted)
		// tax := CalTaxPTI(pti)
		// var b map[string]float64
		// err := json.Unmarshal([]byte(tax), &b)
		// if err != nil {
		// 	fmt.Println("error:", err)
		// }

		// got := b["tax"]
		assert.Equal(t, want, got, "It's should return 4000.00 when I give amount 500000.00 and wht 25000")
	})

	t.Run("given amount more than 660000 with whf 25000 should return calculate tax as 25000", func(t *testing.T) {
		// amount := 660000.00
		// wht := 25000.00
		// personalDeducted := 60000.00
		var pti = PersonalTaxInfo{
			Income:           660000.00,
			Wht:              25000.00,
			PersonalDeducted: 60000.00,
			Donation:         0.00,
		}

		want := 25000.00

		tax := CalTaxPTITaxLevel(pti)
		var taxData TaxData
		err := json.Unmarshal([]byte(tax), &taxData)
		if err != nil {
			fmt.Println("error:", err)
		}
		got := taxData.Tax
		// tax := CalTax(amount, wht, personalDeducted)
		// tax := CalTaxPTI(pti)
		// var b map[string]float64
		// err := json.Unmarshal([]byte(tax), &b)
		// if err != nil {
		// 	fmt.Println("error:", err)
		// }

		// got := b["tax"]
		assert.Equal(t, want, got, "It's should return 25000.00 when I give 660000.00 with 25000 wht")
	})
	t.Run("given amount  1000000 with wht 25000 should return calculate tax as 76000", func(t *testing.T) {
		// amount := 1000000.00
		// wht := 25000.00
		// personalDeducted := 60000.00
		var pti = PersonalTaxInfo{
			Income:           1000000.00,
			Wht:              25000.00,
			PersonalDeducted: 60000.00,
			Donation:         0.00,
		}

		want := 76000.00

		tax := CalTaxPTITaxLevel(pti)
		var taxData TaxData
		err := json.Unmarshal([]byte(tax), &taxData)
		if err != nil {
			fmt.Println("error:", err)
		}
		got := taxData.Tax

		// tax := CalTax(amount, wht, personalDeducted)
		// tax := CalTaxPTI(pti)
		// var b map[string]float64
		// err := json.Unmarshal([]byte(tax), &b)
		// if err != nil {
		// 	fmt.Println("error:", err)
		// }

		// got := b["tax"]
		assert.Equal(t, want, got, "It's should return 76000.00 when I give 1000000.00 with wht 25000")
	})

	t.Run("given amount  1900000 with wht 25000 should return calculate tax as 253000", func(t *testing.T) {
		// amount := 1900000.00
		// wht := 25000.00
		// personalDeducted := 60000.00
		var pti = PersonalTaxInfo{
			Income:           1900000.00,
			Wht:              25000.00,
			PersonalDeducted: 60000.00,
			Donation:         0.00,
		}

		want := 253000.00

		tax := CalTaxPTITaxLevel(pti)
		var taxData TaxData
		err := json.Unmarshal([]byte(tax), &taxData)
		if err != nil {
			fmt.Println("error:", err)
		}
		got := taxData.Tax
		
		assert.Equal(t, want, got, "It's should return 253000.00 when I give 1900000.00")
	})

	t.Run("given amount  2100000 with wht 25000 should return calculate tax as 299000", func(t *testing.T) {
		// amount := 2100000.00
		// wht := 25000.00
		// personalDeducted := 60000.00
		var pti = PersonalTaxInfo{
			Income:           2100000.00,
			Wht:              25000.00,
			PersonalDeducted: 60000.00,
			Donation:         0.00,
		}

		want := 299000.00

		tax := CalTaxPTITaxLevel(pti)
		var taxData TaxData
		err := json.Unmarshal([]byte(tax), &taxData)
		if err != nil {
			fmt.Println("error:", err)
		}
		got := taxData.Tax
		assert.Equal(t, want, got, "It's should return 299000.00 when I give 2100000.00 with 25000 wht")
	})

	t.Run("given amount 500000 with donation 100000 should return calculate tax as 19000", func(t *testing.T) {
		// amount := 500000.00
		// wht := 0.00
		// personalDeducted := 60000.00
		// donation := 100000
		var pti = PersonalTaxInfo{
			Income:           500000.00,
			Wht:              0.00,
			PersonalDeducted: 60000.00,
			Donation:         100000.00,
		}

		want := 19000.00
		tax := CalTaxPTITaxLevel(pti)
		var taxData TaxData
		err := json.Unmarshal([]byte(tax), &taxData)
		if err != nil {
			fmt.Println("error:", err)
		}
		got := taxData.Tax
		assert.Equal(t, want, got, "It's should return 19000.00 when I give 500000.00 with 100000 donation")
	})
}

func TestTaxWithResponseBody(t *testing.T) {
	var pti = PersonalTaxInfo{
			Income:           500000.0,
			Wht:              0.0,
			PersonalDeducted: 60000.0,
			Donation:         100000.0,
		}
	bill := CalTaxPTITaxLevel(pti)
	exp := `{"tax": 19000.0,"taxLevel": [{"level": "2,000,001 ขึ้นไป","tax": 0.0},{"level": "1,000,001-2,000,000","tax": 0.0},{"level": "500,001-1,000,000","tax": 0.0},{"level": "150,001-500,000","tax": 19000.0},{"level": "0-150,000","tax": 0.0}]}`

	bill = strings.TrimSpace(bill)
	exp = strings.ReplaceAll(exp, "\n", "")
	exp = strings.TrimSpace(exp)
	if exp != bill {
		t.Errorf("expect %q but got %q", exp, bill)
	}

	
}

func TestTaxWithAmount2100000WithResponseBody(t *testing.T) {
	var pti = PersonalTaxInfo{
			Income:           2100000.0,
			Wht:              0.0,
			PersonalDeducted: 60000.0,
			Donation:         0.0,
		}
	bill := CalTaxPTITaxLevel(pti)
	exp := `{"tax": 324000.0,"taxLevel": [{"level": "2,000,001 ขึ้นไป","tax": 14000.0},{"level": "1,000,001-2,000,000","tax": 200000.0},{"level": "500,001-1,000,000","tax": 75000.0},{"level": "150,001-500,000","tax": 35000.0},{"level": "0-150,000","tax": 0.0}]}`

	bill = strings.TrimSpace(bill)
	exp = strings.ReplaceAll(exp, "\n", "")
	exp = strings.TrimSpace(exp)
	if exp != bill {
		t.Errorf("expect %q but got %q", exp, bill)
	}

	
}

func TestTaxWithAmount1900000WithResponseBody(t *testing.T) {
	var pti = PersonalTaxInfo{
			Income:           1900000.0,
			Wht:              0.0,
			PersonalDeducted: 60000.0,
			Donation:         0.0,
		}
	bill := CalTaxPTITaxLevel(pti)
	exp := `{"tax": 278000.0,"taxLevel": [{"level": "2,000,001 ขึ้นไป","tax": 0.0},{"level": "1,000,001-2,000,000","tax": 168000.0},{"level": "500,001-1,000,000","tax": 75000.0},{"level": "150,001-500,000","tax": 35000.0},{"level": "0-150,000","tax": 0.0}]}`

	bill = strings.TrimSpace(bill)
	exp = strings.ReplaceAll(exp, "\n", "")
	exp = strings.TrimSpace(exp)
	if exp != bill {
		t.Errorf("expect %q but got %q", exp, bill)
	}

	
}
func TestTaxWithAmount1000000WithResponseBody(t *testing.T) {
	var pti = PersonalTaxInfo{
			Income:           1000000.0,
			Wht:              0.00,
			PersonalDeducted: 60000.0,
			Donation:         0.0,
		}
	bill := CalTaxPTITaxLevel(pti)
	exp := `{"tax": 101000.0,"taxLevel": [{"level": "2,000,001 ขึ้นไป","tax": 0.0},{"level": "1,000,001-2,000,000","tax": 0.0},{"level": "500,001-1,000,000","tax": 66000.0},{"level": "150,001-500,000","tax": 35000.0},{"level": "0-150,000","tax": 0.0}]}`

	bill = strings.TrimSpace(bill)
	exp = strings.ReplaceAll(exp, "\n", "")
	exp = strings.TrimSpace(exp)
	if exp != bill {
		t.Errorf("expect %q but got %q", exp, bill)
	}

	
}


package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"syscall"

	// "github.com/aspexp/assessment-tax/postgres"
	"github.com/aspexp/assessment-tax/tax"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RequestBody struct {
	TotalIncome float64     `json:"totalIncome"`
	Wht         float64     `json:"wht"`
	Allowances  []Allowance `json:"allowances"`
}
type Allowance struct {
	Type   string  `json:"allowanceType"`
	Amount float64 `json:"amount"`
}

type UploadBody struct {
	TotalIncome float64 `json:"totalIncome"`
	Wht         float64 `json:"wht"`
	Donation    float64 `json:"donation"`
}
type PDRequestBody struct {
	Amount float64 `json:"amount"`
}

type TaxData struct {
	Tax F64 `json:"tax"`
	TaxLevel []TaxLevel `json:"taxLevel"`
	
}

type TaxLevel struct {
	Level string `json:"level`
	Tax float64 `json:"tax"`
}


type F64 float64
func (f *F64) UnmarshalJSON(b []byte) error {
	var v float64
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err 
	}
	*f = F64(v)
	return nil 
}


func main() {

	e := echo.New()
	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})
	
	e.POST("/tax/calculations", func(c echo.Context) error {

		var body RequestBody
		err := c.Bind(&body)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		var deductAmount []float64
		var pti tax.PersonalTaxInfo
		if len(body.Allowances) == 1 {
			for _, bd := range body.Allowances {
						if bd.Type == "k-receipt" {
							if bd.Amount > 50000.0 {
								deductAmount = append(deductAmount, 50000.0)
							}
							if bd.Amount <= 50000 && bd.Amount >= 1 {
								deductAmount = append(deductAmount, bd.Amount)
							}
								pti = tax.PersonalTaxInfo{
								Income:           body.TotalIncome,
								Wht:              body.Wht,
								PersonalDeducted: 60000.0,
								Donation:         0.0,
								KReceipt:         deductAmount[0],
							}
						}
						if bd.Type == "donation" {

							if bd.Amount > 100000.0 {
								deductAmount = append(deductAmount, 100000.0)
							}
							if bd.Amount <= 100000  {
								deductAmount = append(deductAmount, bd.Amount)
							}

							pti = tax.PersonalTaxInfo{
								Income:           body.TotalIncome,
								Wht:              body.Wht,
								PersonalDeducted: 60000.0,
								Donation:         deductAmount[0],
								KReceipt:         0.0,
							}
						}
					}
		}else{
		// if len(body.Allowances) > 1 {
			for _, bd := range body.Allowances {
				if bd.Type == "k-receipt" {
						if bd.Amount > 50000 {
							deductAmount = append(deductAmount, 50000)
						}
						if bd.Amount <= 50000 && bd.Amount >= 1 {
							deductAmount = append(deductAmount, bd.Amount)
						}
							
					}
					if bd.Type == "donation" {

						deductAmount = append(deductAmount, bd.Amount)
						
					}
				}
			
				 pti = tax.PersonalTaxInfo{
						Income:           body.TotalIncome,
						Wht:              body.Wht,
						PersonalDeducted: 60000.0,
						Donation:         deductAmount[1],
						KReceipt:         deductAmount[0],
					}
					
		}
		
		result := tax.TaxHandler(pti)
		var data TaxData
		err = json.Unmarshal([]byte(result), &data )
		// err = data.UnmarshalJSON([]byte(result) )
		if err != nil {
			fmt.Println("Error unmarshalling JSON:", err )
			return c.JSON(http.StatusInternalServerError, err )
		}
		sort.Slice(data.TaxLevel, func(i, j int) bool {
			return i > j 
		})

		reversedJSON, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Println("Error mashalling JSON:", err )
			return c.JSON(http.StatusInternalServerError, err )
		}

		// var data2 map[string]interface{}
		// json.Unmarshal([]byte(result), &data2)
		// return c.JSON(http.StatusOK, data2)

		var data1 map[string]interface{}
		json.Unmarshal([]byte(reversedJSON), &data1)
		return c.JSON(http.StatusOK, data1)


	})
	e.POST("/admin/deductions/personal", func(c echo.Context) error {
		var pd PDRequestBody
		err := c.Bind(&pd)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		retStr := tax.InsertPersonalDeduct(pd.Amount)

		var data map[string]interface{}
		json.Unmarshal([]byte(retStr), &data)
		return c.JSON(http.StatusOK, data)
	})
	e.POST("/admin/deductions/k-receipt", func(c echo.Context) error {
		var pd PDRequestBody
		err := c.Bind(&pd)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		retStr := tax.InsertKReceiptDeduct(pd.Amount)
		var data map[string]interface{}
		json.Unmarshal([]byte(retStr), &data)
		return c.JSON(http.StatusOK, data)
	})
	e.POST("/tax/calculations/upload-csv", func(c echo.Context) error {
		result := tax.GetTaxCSV("/tax/upload/taxes.csv")
		var data map[string]interface{}
		json.Unmarshal([]byte(result), &data)
		return c.JSON(http.StatusOK, data)
	})

	serverPort := ":" + os.Getenv("PORT")
	go func() {
		e.Logger.Fatal(e.Start(serverPort))
	}()
	log.Println("server started.")

	<-stop
	log.Println("shutting down the server")

}

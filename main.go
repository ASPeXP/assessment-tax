package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	// "github.com/aspexp/assessment-tax/postgres"
	"github.com/aspexp/assessment-tax/tax"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
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

func main() {

	e := echo.New()
	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

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
		for _, bd := range body.Allowances {
			if bd.Type == "k-receipt" {
				if bd.Amount > 100000 {
					deductAmount = append(deductAmount, 100000)
				}
				if bd.Amount <= 100000 && bd.Amount >= 1 {
					deductAmount = append(deductAmount, bd.Amount)
				}
			}
			if bd.Type == "donation" {

				deductAmount = append(deductAmount, bd.Amount)

			}
		}
		var pti = tax.PersonalTaxInfo{
			Income:           body.TotalIncome,
			Wht:              body.Wht,
			PersonalDeducted: 60000.00,
			Donation:         deductAmount[1],
			KReceipt:         deductAmount[0],
		}
		result := tax.TaxHandler(pti)
		fmt.Println(result)
		fmt.Printf("donation %.1f", pti.Donation)
		fmt.Printf("k-receipt %.1f", pti.KReceipt)
		return c.JSON(http.StatusOK, result)

	})
	e.POST("/admin/deductions/personal", func(c echo.Context) error {
		var pd PDRequestBody
		err := c.Bind(&pd)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		retStr := tax.InsertPersonalDeduct(pd.Amount)

		return c.JSON(http.StatusOK, retStr)
	})
	e.POST("/admin/deductions/k-receipt", func(c echo.Context) error {
		var pd PDRequestBody
		err := c.Bind(&pd)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		retStr := tax.InsertKReceiptDeduct(pd.Amount)

		return c.JSON(http.StatusOK, retStr)
	})
	e.POST("/tax/calculations/upload-csv", func(c echo.Context) error {
		result := tax.GetTaxCSV("/tax/upload/taxes.csv")

		return c.JSON(http.StatusOK, result)
	})

	serverPort := ":" + os.Getenv("PORT")
	go func() {
		e.Logger.Fatal(e.Start(serverPort))
	}()
	log.Println("server started.")

	<-stop
	log.Println("shutting down the server")

}

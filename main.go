package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/aspexp/assessment-tax/tax"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
)

type RequestBody struct {
	TotalIncome float64 `json:"totalIncome"`
  	Wht  float64 `json:"wht"`
	Allowances []Allowance `json:"allowances"`
}
type Allowance struct {
	
	Type string  `json:"allowanceType"`
	Amount float64 `json:"amount"`
 
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
	e.POST("/tax/calculation", func( c echo.Context) error {

		var body RequestBody
		err := c.Bind(&body)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error()) 
		}

		var pti = tax.PersonalTaxInfo{
			Income: body.TotalIncome,
			Wht: body.Wht,
			PersonalDeducted: 60000.00,
			Donation: body.Allowances[0].Amount,
		}
		result := tax.TaxHandler(pti)
		fmt.Println(result)
		fmt.Printf("donation %.2f", pti.Donation)
		return c.JSON(http.StatusOK, result )
		
	})
	e.POST("/admin/deductions/personal", func( c echo.Context) error {
		var pd PDRequestBody 
		err := c.Bind(&pd)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error()) 
		}
		fmt.Println(pd)
		retStr := fmt.Sprintf(`{"personalDeduction": %.1f }`, pd.Amount)
		return c.JSON(http.StatusOK,retStr)
	})

	serverPort := ":" + os.Getenv("PORT")
	go func(){
	e.Logger.Fatal(e.Start(serverPort))
	}()
	log.Println("server started.")

	<-stop 
	log.Println("shutting down the server")

}

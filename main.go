package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/aspexp/assessment-tax/tax"
	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RequestBody struct {
TotalIncome float64 `json:"totalIncome"`
  	Wht  float64 `json:"wht"`
	Allowances []Allowance `json:"allowances"`
}
type Allowance struct {
	
	AllowanceType string  `json:"allowanceType"`
	Allowance_amount float64 `json:"allowance_amount"`
 
  
}
func main() {
	
	
	e := echo.New()
	  // Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

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

		personalDeducted := 60000.00
		result := tax.TaxHandler(body.TotalIncome, body.Wht, personalDeducted)

		return c.JSON(http.StatusOK, result )
	})

	serverPort := ":" + os.Getenv("PORT")
	go func(){
	e.Logger.Fatal(e.Start(serverPort))
	}()
	log.Println("server started.")

	<-stop 
	log.Println("shutting down the server")

}

package main

import (
	"log"

	"github.com/labstack/echo/v4"

	booking_setup "github.com/antosdaniel/go-presentation-collaborating-on-architecture/modules/booking/setup"
)

func main() {
	booking, err := booking_setup.New()
	if err != nil {
		log.Fatalf("Could not set up booking module: %v", err)
	}

	e := echo.New()
	booking.RegisterRoutes(e.Group("/booking"))

	e.Logger.Fatal(e.Start(":8000"))
}

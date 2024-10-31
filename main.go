package main

import (
	"log"

	booking_setup "github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/booking/setup"
	no_setup "github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/no"
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/pkg/metrics"
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/pkg/setup"
)

func main() {
	s := setup.New()

	no, err := no_setup.New()
	if err != nil {
		log.Fatalf("Could not set up no module: %v", err)
	}

	booking, err := booking_setup.New(s.Enqueueer, no.API)
	if err != nil {
		log.Fatalf("Could not set up booking module: %v", err)
	}

	go booking.ListenForEvents(s.Dequeueer)
	go no.PeriodicRetry()

	metrics.RegisterRoutes(s.Router)
	booking.RegisterRoutes(s.Router.Group("/booking"))
	s.StartServer()
}

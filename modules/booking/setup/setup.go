package booking_setup

import (
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/modules/booking"
	"github.com/labstack/echo/v4"
)

func New() (BookingSetup, error) {
	return BookingSetup{}, nil
}

type BookingSetup struct {
}

func (b BookingSetup) RegisterRoutes(g *echo.Group) {
	g.POST("/book", func(c echo.Context) error {
		booking.Book()
		return nil
	})
}

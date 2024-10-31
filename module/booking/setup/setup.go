package booking_setup

import (
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/booking"
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/no/api"
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/pkg/queue"
	"github.com/labstack/echo/v4"
)

func New(enqueue queue.Enqueue, noAPI api.API) (BookingSetup, error) {
	return BookingSetup{
		manager:       booking.New(enqueue),
		notifications: noAPI,
	}, nil
}

type BookingSetup struct {
	manager       booking.Manager
	notifications api.API
}

func (b BookingSetup) RegisterRoutes(g *echo.Group) {
	g.POST("/book", func(c echo.Context) error {
		b.manager.Book()
		return nil
	})
}

func (b BookingSetup) ListenForEvents(dequeue queue.Dequeue) {
	queue.ListenFor(dequeue, booking.BookingRequested, func() {
		b.notifications.SendNotification("template-id", "user-id")
	})
}

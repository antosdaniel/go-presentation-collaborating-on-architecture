package booking_setup

import (
	"log/slog"

	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/booking"
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/module/notification"
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/pkg/queue"
	"github.com/labstack/echo/v4"
)

func New(enqueue queue.Enqueue, noAPI notification.API) (BookingSetup, error) {
	return BookingSetup{
		manager:       booking.New(enqueue),
		notifications: noAPI,
	}, nil
}

type BookingSetup struct {
	manager       booking.Manager
	notifications notification.API
}

func (b BookingSetup) RegisterRoutes(g *echo.Group) {
	g.POST("/request", func(c echo.Context) error {
		b.manager.RequestBooking()
		return nil
	})
}

func (b BookingSetup) ListenForEvents(dequeue queue.Dequeue) {
	queue.ListenFor(dequeue, booking.BookingRequested, func() {
		err := b.notifications.SendNotification("template-id", "user-id")
		if err != nil {
			slog.Error("Sending booking requested notification failed", "error", err)
		}
	})
}

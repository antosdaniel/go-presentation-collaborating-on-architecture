package booking

import (
	"log/slog"

	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/pkg/queue"
)

type Manager struct {
	enqueue queue.Enqueue
}

func New(enqueue queue.Enqueue) Manager {
	return Manager{enqueue: enqueue}
}

func (m Manager) Book() {
	slog.Info("Booking successfully requested")

	m.enqueue.Enqueue(BookingRequested)
}

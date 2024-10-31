package setup

import (
	"github.com/antosdaniel/go-presentation-collaborating-on-architecture/pkg/queue"
	"github.com/labstack/echo/v4"
)

func New() Setup {
	return Setup{
		Router:    echo.New(),
		Enqueueer: queue.NewEnqueuer(),
		Dequeueer: queue.NewDequeuer(),
	}
}

type Setup struct {
	Router    *echo.Echo
	Enqueueer queue.Enqueue
	Dequeueer queue.Dequeue
}

func (s Setup) StartServer() {
	s.Router.Logger.Fatal(s.Router.Start(":8000"))
}

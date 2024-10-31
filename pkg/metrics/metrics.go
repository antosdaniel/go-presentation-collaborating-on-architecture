package metrics

import "github.com/labstack/echo/v4"

func RegisterRoutes(g *echo.Echo) {
	g.GET("/metrics", func(c echo.Context) error {
		return nil
	})
}

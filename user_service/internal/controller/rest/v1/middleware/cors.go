package middleware

import "github.com/labstack/echo/v4"

func CORS(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Access-control-allow-origin", "*")
		return next(c)
	}
}

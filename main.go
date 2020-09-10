package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/thaiPost/getToken", GetToken)
	e.POST("/thaiPost/getTracking", GetTracking)
	e.Logger.Fatal(e.Start(":1323"))
}

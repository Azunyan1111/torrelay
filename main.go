package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK,`tor -socksport 0 -orport 80 -dirport 443 -exitrelay 0 -nickname YourNickName -contactinfo YourEmail -exitpolicy "reject *:*"`)
	})
	e.Start(":8081")
}
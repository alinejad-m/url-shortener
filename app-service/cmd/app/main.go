package main

import (
	"app-service/internal/rest"
	"fmt"
	"github.com/labstack/echo/v4"
)

const (
	defaultTimeout = 30
	defaultAddress = ":9090"
)

func main() {
	fmt.Println("App Main Started...")

	e := echo.New()
	e.Use(middleware.CORS)
}

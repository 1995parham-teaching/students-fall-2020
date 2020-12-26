package main

import (
	"fmt"

	"github.com/aut-ce/students.go/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/hello", handler.Hello)

	s := handler.Student{}
	e.POST("/student", s.Create)

	if err := e.Start("0.0.0.0:8080"); err != nil {
		fmt.Println(err)
	}
}

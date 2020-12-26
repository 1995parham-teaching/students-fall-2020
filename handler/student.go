package handler

import (
	"fmt"
	"net/http"

	"github.com/aut-ce/students.go/model"
	"github.com/labstack/echo/v4"
)

type Student struct{}

type request struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (s *Student) Create(c echo.Context) error {
	var req request

	if err := c.Bind(&req); err != nil {
		fmt.Println(err)

		return echo.ErrBadRequest
	}

	fmt.Println(model.Student{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		ID:        0,
	})

	return c.JSON(http.StatusCreated, "")
}

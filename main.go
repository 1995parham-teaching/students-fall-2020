package main

import (
	"fmt"
	"log"

	"github.com/aut-ce/students.go/handler"
	"github.com/aut-ce/students.go/model"
	"github.com/aut-ce/students.go/store"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	db, err := gorm.Open(sqlite.Open("students.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// disabling gorm logs
	// db.Logger.LogMode(logger.Silent)

	if err := db.Migrator().AutoMigrate(&model.Student{}); err != nil {
		log.Fatal(err)
	}

	e.GET("/hello", handler.Hello)

	s := handler.NewStudent(store.NewSQLStudent(db))
	e.POST("/student", s.Create)
	e.GET("/student", s.Find)

	if err := e.Start("0.0.0.0:8080"); err != nil {
		fmt.Println(err)
	}
}

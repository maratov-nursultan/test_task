package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"test_task/internal/database"
	"test_task/internal/handler"
	"test_task/internal/service"
)

func main() {
	db, err := database.ConnectPostgres()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	factory := service.NewService(db)
	if err != nil {
		log.Fatalf("Kaput factory is not working %v", err)
	}

	newHandler := handler.NewHandler(factory.GetUserManager())

	e := echo.New()

	e.GET("/iin_check/:iin", newHandler.CheckIin)
	e.POST("/people/info", newHandler.CreateUser)
	e.GET("/people/info/iin/:iin", newHandler.GetUserByIin)
	e.GET("/people/info/phone/:name", newHandler.ListUserByName)

	e.Logger.Fatal(e.Start(":80"))

}

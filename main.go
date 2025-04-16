package main

import (
	"github.com/labstack/echo/v4"
	"github.com/maratov-nursultan/profile/internal/database"
	"github.com/maratov-nursultan/profile/internal/handler"
	"github.com/maratov-nursultan/profile/internal/service"
	"log"
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

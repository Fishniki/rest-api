package main

import (
	"fmt"
	"rest-api/internal/api"
	"rest-api/internal/config"
	"rest-api/internal/connection"
	"rest-api/internal/repository"
	"rest-api/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	db := connection.GetDatabase(cnf.Database)

	app := fiber.New()

	customerRepository := repository.NewCustomer(db)

	customerService := service.NewCustomer(customerRepository)

	api.NewCustomer(app, customerService)

	// app.Get("/developers", develoPers)
	_ =app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)

	fmt.Println("Server is running on port " + cnf.Server.Port)
}


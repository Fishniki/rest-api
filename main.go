package main

import (
	"fmt"
	"net/http"
	"rest-api/dto"
	"rest-api/internal/api"
	"rest-api/internal/config"
	"rest-api/internal/connection"
	"rest-api/internal/repository"
	"rest-api/internal/service"

	jwtMid "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	db := connection.GetDatabase(cnf.Database)

	app := fiber.New()
	jwtMid := jwtMid.New(jwtMid.Config{
		SigningKey: jwtMid.SigningKey{
			Key: []byte(cnf.Jwt.Key),
		},
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			return ctx.Status(http.StatusUnauthorized).
				JSON(dto.CreateResponsError("endpoint perlu token JWT, silahkan login terlebih dahulu"))
		},
	})

	customerRepository := repository.NewCustomer(db)
	userRepository := repository.NewUser(db)
	bookRepository := repository.NewBook(db)
	bookStockRepository := repository.NewBookStock(db)
	journalRepository := repository.NewJournal(db)
	mediaRepository := repository.NewMedia(db)

	customerService := service.NewCustomer(customerRepository)
	authService := service.NewAuth(cnf, userRepository)
	bookService := service.NewBook(bookRepository, bookStockRepository)
	bookStockService := service.NewBookStock(bookRepository, bookStockRepository)
	journalService := service.NewJournal(journalRepository, bookRepository, bookStockRepository, customerRepository)
	mediaService := service.NewMedia(cnf, mediaRepository)
	

	api.NewCustomer(app, customerService, jwtMid)
	api.NewAuth(app, authService)
	api.NewBook(app, bookService, jwtMid)
	api.NewBookStock(app, bookStockService, jwtMid)
	api.NewJournal(app, journalService, jwtMid)
	api.NewMedia(app, mediaService, jwtMid, cnf)

	// app.Get("/developers", develoPers)
	_ =app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)

	fmt.Println("Server is running on port " + cnf.Server.Port)
}


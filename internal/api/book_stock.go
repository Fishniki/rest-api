package api

import (
	"context"
	"net/http"
	"rest-api/domain"
	"rest-api/dto"
	"rest-api/internal/utility"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type bookStockApi struct {
	bookStockService domain.BookStockService
}

func NewBookStock(app *fiber.App, bookStockService domain.BookStockService, authMid fiber.Handler) {
	bsa := bookStockApi{
		bookStockService: bookStockService,
	}

	app.Post("/book-stocks", authMid, bsa.Create)
	app.Delete("/book-stocks", authMid, bsa.Delete)

}


func (ba bookStockApi) Create(ctx *fiber.Ctx) error {

	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var req dto.CreateBookStockRequest
	if err := ctx.BodyParser(&req); err != nil {
		return  ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := utility.Validate(req)
	if len(fails) > 0 {
		return  ctx.Status(http.StatusBadRequest).
				JSON(dto.CreateResponsErrorData("validation failed", fails))
	}

	err := ba.bookStockService.Create(c, req)
	if err != nil {
		return  ctx.Status(http.StatusInternalServerError).
				JSON(dto.CreateResponsError(err.Error()))
	}

	return  ctx.Status(http.StatusOK).
			JSON(dto.CreateResponsSucces("succses"))

}

func (ba bookStockApi) Delete(ctx *fiber.Ctx) error {

	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()


	codeStr := ctx.Query("code")
	if len(codeStr) < 1 {
		return ctx.Status(http.StatusBadRequest).
				JSON(dto.CreateResponsError("parameter code wajib diisi"))
	}
	codes := strings.Split(codeStr, ";")

	err := ba.bookStockService.Delete(c, dto.DeleteBookStockRequest{Codes: codes})
	if err != nil {
		return  ctx.Status(http.StatusInternalServerError).
				JSON(dto.CreateResponsError(err.Error()))
	}

	return  ctx.Status(http.StatusCreated).
			JSON(dto.CreateResponsSucces("succses"))

}
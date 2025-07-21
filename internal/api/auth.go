package api

import (
	"context"
	"net/http"
	"rest-api/domain"
	"rest-api/dto"
	"rest-api/internal/utility"
	"time"

	"github.com/gofiber/fiber/v2"
)

type authApi struct {
	authService domain.AuthService
}

func NewAuth(app *fiber.App, authService domain.AuthService) {
	aa := authApi{
		authService: authService,
	}

	app.Post("/auth", aa.Login)
	app.Post("/register", aa.Register)
}

func (aa authApi) Login(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var req dto.AuthRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}
	res, err := aa.authService.Login(c, req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
			JSON(dto.CreateResponsError(err.Error()))
	}

	return  ctx.Status(http.StatusOK).
			JSON(dto.CreateResponsSucces(res))

}

func (aa authApi) Register(ctx *fiber.Ctx) error {

	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var req dto.RegisterRequest
	if err := ctx.BodyParser(&req); err != nil {
		return  ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := utility.Validate(req)
	if len(fails) > 0 {
		return  ctx.Status(http.StatusBadRequest).
				JSON(dto.CreateResponsErrorData("Validasi gagal", fails))
	}

	err := aa.authService.Register(c, req)
	if err != nil {
		return  ctx.Status(http.StatusInternalServerError).
				JSON(dto.CreateResponsError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).
			JSON(dto.CreateResponsSucces("data berhasil terkirim"))
}
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

type customerAPI struct {
	customerService domain.CustomerService 
}

func NewCustomer(app *fiber.App, customerService domain.CustomerService){
	ca := customerAPI{
		customerService: customerService,
	}

	app.Get("/customers/", ca.Index)
	app.Post("/customers/", ca.Create)
	app.Put("/customers/:customer_id/", ca.Update)
	app.Delete("/customers/:customer_id/", ca.Delete)
	app.Get("/customers/:customer_id/", ca.Show)

}

func (ca customerAPI) Index(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	res, err := ca.customerService.Index(c)
	if err != nil {
		return  ctx.Status(http.StatusInternalServerError).
				JSON(dto.CreateResponsError(err.Error()))
	}

	return  ctx.JSON(dto.CreateResponsSucces(res))
}

func (ca customerAPI) Create(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var req  dto.CreateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := utility.Validate(req)
	if len(fails) > 0 {
		return  ctx.Status(http.StatusBadRequest).
				JSON(dto.CreateResponsErrorData("validation failed", fails))
	}

	err := ca.customerService.Create(c, req) 
	if err != nil {
		return  ctx.Status(http.StatusInternalServerError).
				JSON(dto.CreateResponsError(err.Error()))
	}

	return  ctx.Status(http.StatusCreated).
			JSON(dto.CreateResponsSucces("succses"))
}

func (ca customerAPI) Update(ctx *fiber.Ctx) error {

	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	var req dto .UpdateCustomerRequest
	if err := ctx.BodyParser(&req); err != nil {
		return  ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	fails := utility.Validate(req)
	if len(fails) > 0 {
		return  ctx.Status(http.StatusBadRequest).
				JSON(dto.CreateResponsErrorData("validation error", fails))
	}

	//customer
	req.ID = ctx.Params("customer_id")
	err := ca.customerService.Update(c, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
				JSON(dto.CreateResponsError(err.Error()))
	}

	return  ctx.Status(http.StatusOK).
			JSON(dto.CreateResponsSucces("Succses"))

}


func (ca customerAPI) Delete(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30 * time.Second)
	defer cancel()

	id := ctx.Params("customer_id")
	err := ca.customerService.Delete(c, id)
	if err != nil {
		return  ctx.Status(http.StatusInternalServerError).
				JSON(dto.CreateResponsError(err.Error()))
	}

	return  ctx.Status(http.StatusNoContent).
			JSON(dto.CreateResponsSucces("Succses"))

}

func (ca customerAPI) Show(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 30 * time.Second)
	defer cancel()

	id := ctx.Params("customer_id")
	res, err := ca.customerService.Show(c, id)
	if err != nil {
		return  ctx.Status(http.StatusInternalServerError).
				JSON(dto.CreateResponsError(err.Error()))
	}

	return  ctx.Status(http.StatusOK).
			JSON(dto.CreateResponsSucces(res))
}

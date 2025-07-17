package api

import (
	"context"
	"net/http"
	"path/filepath"
	"rest-api/domain"
	"rest-api/dto"
	"rest-api/internal/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type mediaApi struct {
	cnf *config.Config
	mediaService domain.MediaService
}

func NewMedia(app *fiber.App, mediaService domain.MediaService, authMid fiber.Handler, cnf *config.Config) {
	ma := mediaApi{
		cnf: cnf,
		mediaService: mediaService,
	}

	app.Post("/media", authMid, ma.Create)
	app.Static("/media", cnf.Storage.BasePath)

}

func (ma mediaApi) Create(ctx *fiber.Ctx) error {

	c, cancel := context.WithTimeout(ctx.Context(), 10 * time.Second)
	defer cancel()

	file, err := ctx.FormFile("media")
	if err != nil {
		return ctx.SendStatus(http.StatusBadRequest)	
	}

	filename := uuid.NewString() + file.Filename
	path := filepath.Join(ma.cnf.Storage.BasePath, filename)

	err = ctx.SaveFile(file, path)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
				JSON(dto.CreateResponsError(err.Error()))
	}

	res, err := ma.mediaService.Create(c, dto.CreateMediaRequest{
		Path: filename,
	})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).
				JSON(dto.CreateResponsError(err.Error()))
	}

	return ctx.Status(http.StatusCreated).
			JSON(dto.CreateResponsSucces(res))

}


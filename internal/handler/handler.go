package handler

import (
	"github.com/didsqq/ZeroAgency_Test/internal/service"
	"github.com/gofiber/fiber"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *fiber.App {
	app := fiber.New()

	app.Get("/list", h.GetAll)
	app.Post("/edit/:Id", h.Update)

	return app
}

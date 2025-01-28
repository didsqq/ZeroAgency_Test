package handler

import "github.com/gofiber/fiber"

func (h *Handler) GetAll(c *fiber.Ctx) {

}

func (h *Handler) Update(c *fiber.Ctx) {
	id := c.Params("Id")
}

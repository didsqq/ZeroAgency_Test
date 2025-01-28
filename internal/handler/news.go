package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/didsqq/ZeroAgency_Test/internal/models"
	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetAll(c *fiber.Ctx) {
	news, err := h.services.GetAll()
	if err != nil {
		logrus.Warn("getall error")
		c.Status(http.StatusBadRequest).SendString("getall error")
		return
	}

	if err := c.JSON(news); err != nil {
		logrus.Warn("json serialization error")
	}
}

func (h *Handler) Update(c *fiber.Ctx) {
	newsId, err := strconv.Atoi(c.Params("Id"))
	if err != nil {
		logrus.Warn("invalid list id param")
		c.Status(http.StatusBadRequest).SendString("invalid list id param")
		return
	}

	var input models.News
	if err := c.BodyParser(&input); err != nil {
		logrus.Warn("body parse error")
		c.Status(http.StatusBadRequest).SendString("body parse error")
		return
	}

	if err = h.services.News.Update(newsId, input); err != nil {
		msg := fmt.Sprintf("news update error: %v", err)
		logrus.Warn(msg)
		c.Status(http.StatusBadRequest).SendString(msg)
		return
	}

	c.JSON(fiber.Map{"message": "Success"})
}

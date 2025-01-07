package handlers

import (
	"context"
	"github.com/MrScotch679/book-tickets-backend/models"
	"github.com/gofiber/fiber/v2"
	"time"
)

type EventHandler struct {
	repository models.EventRepository
}

func (h *EventHandler) GetMany(c *fiber.Ctx) error {
	context, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	events, err := h.repository.GetMany(context)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(events)
}

func (h *EventHandler) GetOne(c *fiber.Ctx) error {
	return nil
}

func (h *EventHandler) CreateOne(c *fiber.Ctx) error {
	return nil
}

func NewEventHandler(router fiber.Router, repository models.EventRepository) {
	handler := &EventHandler{
		repository,
	}

	router.Get("/", handler.GetMany)
	router.Get("/:id", handler.GetOne)
	router.Post("/", handler.CreateOne)
}

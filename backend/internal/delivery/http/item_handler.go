package deliveryHttp

import (
	"strconv"
	"test_tablelink/internal/domain"
	"test_tablelink/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type ItemHandler struct {
	uc usecase.ItemUsecase
}

func NewItemHandler(uc usecase.ItemUsecase) *ItemHandler {
	return &ItemHandler{uc: uc}
}

func (h *ItemHandler) List(c *fiber.Ctx) error {
	limit := c.Query("limit", "10")
	offset := c.Query("offset", "0")

	l, _ := strconv.Atoi(limit)
	o, _ := strconv.Atoi(offset)

	items, err := h.uc.GetAll(c.Context(), l, o)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(items)
}

func (h *ItemHandler) Create(c *fiber.Ctx) error {
	var item domain.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if item.Name == "" || item.Price <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
	}

	if err := h.uc.Create(c.Context(), &item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true, "uuid": item.UUID})
}

func (h *ItemHandler) Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	var item domain.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	item.UUID = uuid

	if err := h.uc.Update(c.Context(), &item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true})
}

func (h *ItemHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	if err := h.uc.HardDelete(c.Context(), uuid); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true})
}

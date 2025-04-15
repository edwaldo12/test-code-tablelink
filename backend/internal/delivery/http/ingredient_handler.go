package deliveryHttp

import (
	"strconv"
	"test_tablelink/internal/domain"
	"test_tablelink/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type IngredientHandler struct {
	uc usecase.IngredientUsecase
}

func NewIngredientHandler(uc usecase.IngredientUsecase) *IngredientHandler {
	return &IngredientHandler{uc: uc}
}

func (h *IngredientHandler) List(c *fiber.Ctx) error {
	limit := c.Query("limit", "10")
	offset := c.Query("offset", "0")

	l, _ := strconv.Atoi(limit)
	o, _ := strconv.Atoi(offset)

	ingredients, err := h.uc.GetAll(c.Context(), l, o)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(ingredients)
}

func (h *IngredientHandler) Create(c *fiber.Ctx) error {
	var ing domain.Ingredient
	if err := c.BodyParser(&ing); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if ing.Name == "" || ing.Type < 0 || ing.Type > 2 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Missing required fields"})
	}

	if err := h.uc.Create(c.Context(), &ing); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"success": true, "uuid": ing.UUID})
}

func (h *IngredientHandler) Update(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	var ing domain.Ingredient
	if err := c.BodyParser(&ing); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	ing.UUID = uuid

	if err := h.uc.Update(c.Context(), &ing); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true})
}

func (h *IngredientHandler) Delete(c *fiber.Ctx) error {
	uuid := c.Params("uuid")

	if err := h.uc.HardDelete(c.Context(), uuid); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"success": true})
}

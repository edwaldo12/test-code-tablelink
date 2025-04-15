package main

import (
	"log"

	http "test_tablelink/internal/delivery/http"
	"test_tablelink/internal/repository"
	"test_tablelink/internal/usecase"
	"test_tablelink/pkg/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	dbPool, err := config.NewPostgresPool()
	if err != nil {
		log.Fatalf("Error establishing DB connection: %v", err)
	}
	defer dbPool.Close()

	ingredientRepo := repository.NewIngredientRepo(dbPool)
	itemRepo := repository.NewItemRepo(dbPool)

	ingredientUsecase := usecase.NewIngredientUsecase(ingredientRepo)
	itemUsecase := usecase.NewItemUsecase(itemRepo)

	ingredientHandler := http.NewIngredientHandler(ingredientUsecase)
	itemHandler := http.NewItemHandler(itemUsecase)

	app.Get("/ingredients", ingredientHandler.List)
	app.Post("/ingredients", ingredientHandler.Create)
	app.Put("/ingredients/:uuid", ingredientHandler.Update)
	// app.Delete("/ingredients/:uuid", ingredientHandler.HardDelete)
	app.Get("/items", itemHandler.List)
	app.Post("/items", itemHandler.Create)
	app.Put("/items/:uuid", itemHandler.Update)
	// app.Delete("/items/:uuid", itemHandler.HardDelete)

	app.Listen(":3001")
}

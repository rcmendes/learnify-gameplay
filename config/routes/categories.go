package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rcmendes/learnify/gameplay/adapters/entrypoints/rest"
	"github.com/rcmendes/learnify/gameplay/core/ucs/ports"
)

func LoadCategoriesRoutes(
	app *fiber.App,
	createCategoryUC ports.CreateCategory,
	findAllCategoriesUC ports.FindAllCategories) {

	ctrl := rest.NewCategoryController(createCategoryUC, findAllCategoriesUC)

	group := app.Group("/categories")
	group.Post("/", ctrl.Create)
	group.Get("/", ctrl.ListAll)
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rcmendes/learnify/gameplay/adapters/entrypoints/rest"
	"github.com/rcmendes/learnify/gameplay/core/ucs/ports"
)

func LoadQuizzesRoutes(
	app *fiber.App,
	findAllQuizzesUC ports.FindAllQuizzes,
	findQuizByCategoryNameUC ports.FindQuizByCategoryName,
	findQuizUC ports.FindQuiz) {

	ctrl := rest.NewQuizController(
		findAllQuizzesUC,
		findQuizByCategoryNameUC,
		findQuizUC)

	group := app.Group("/quizzes")
	group.Get("/", ctrl.ListAll)
	group.Get("/:id", ctrl.FindOneByID)
	group.Get("/:id/image", ctrl.GetImageByID)
	group.Get("/:id/audio", ctrl.GetAudioByID)
	// quizzesGroup.Delete("/:uuid", ctrl.DeleteByUUID)

}

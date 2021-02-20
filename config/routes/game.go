package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rcmendes/learnify/gameplay/adapters/entrypoints/rest"
	"github.com/rcmendes/learnify/gameplay/core/ucs/ports"
)

func LoadGameRoutes(
	app *fiber.App,
	createGameUC ports.CreateGame,
	validateAnswerGameQuizUC ports.ValidateAnswerGameQuiz,
	findOneNotPlayedGameQuizUC ports.FindOneNotPlayedGameQuiz) {

	ctrl := rest.NewGameController(createGameUC, validateAnswerGameQuizUC, findOneNotPlayedGameQuizUC)

	gameGroup := app.Group("/game")
	gameGroup.Post("/", ctrl.Create)
	gameGroup.Get("/:gameID/quizzes/next", ctrl.NexQuiz)
	gameGroup.Put("/:gameID/quizzes/:quizID/validate", ctrl.ValidateAnswer)
}

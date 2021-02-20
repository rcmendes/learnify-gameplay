package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rcmendes/learnify/gameplay/core/entities"
	"github.com/rcmendes/learnify/gameplay/core/ucs/ports"
)

type GameController interface {
	Create(c *fiber.Ctx) error
	ValidateAnswer(c *fiber.Ctx) error
	NexQuiz(c *fiber.Ctx) error
}

type gameController struct {
	create         ports.CreateGame
	validateAnswer ports.ValidateAnswerGameQuiz
	nextQuiz       ports.FindOneNotPlayedGameQuiz
}

func NewGameController(
	createGameUC ports.CreateGame,
	validateAnswerGameQuizUC ports.ValidateAnswerGameQuiz,
	findOneNotPlayedGameQuizUC ports.FindOneNotPlayedGameQuiz,
) GameController {
	return &gameController{
		create:         createGameUC,
		validateAnswer: validateAnswerGameQuizUC,
		nextQuiz:       findOneNotPlayedGameQuizUC,
	}
}

//CreateGameRequest defines the contract of a request for creating a Game.
type CreateGameRequest struct {
	CategoryID uuid.UUID `json:"category_id"`
	PlayerID   uuid.UUID `json:"player_id"`
}

//CreateGameResponse defines the contract of the response of creating a Game.
type CreateGameResponse struct {
	ID uuid.UUID `json:"id"`
}

//GetGameByIDRequest defines the contract of the request for loading a Game.
// type GetGameByIDRequest struct {
// 	ID uuid.UUID `json:"id"`
// }

//GetGameByIDResponse defines the contract of the response of the loaded a Game.
// type GetGameByIDResponse struct {
// 	Question string      `json:"question"`
// 	Images   []MediaLink `json:"images"`
// 	Audio    MediaLink   `json:"audio"`
// 	// Images   []uuid.UUID `json:"images"`
// 	// Audio    uuid.UUID   `json:"audio"`
// }

func (ctrl *gameController) Create(c *fiber.Ctx) error {
	request := new(CreateGameRequest)

	if err := c.BodyParser(request); err != nil {
		//TODO handle error
		return err
	}

	newGame := entities.NewGameData{
		PlayerID:   request.PlayerID,
		CategoryID: request.CategoryID,
	}
	id, err := ctrl.create.Create(newGame)

	if err != nil {
		//TODO Handle Error
		return fiber.ErrInternalServerError
	}

	response := CreateGameResponse{
		ID: *id,
	}

	return c.Status(201).JSON(response)
}

// func (ctrl *GameController) GetGameByID(c *fiber.Ctx) error {
// 	// gameID := c.Params("id")

// 	ids := []string{"44e420ad-c5fd-4d26-9ed5-2838f2450147",
// 		"93ac3e8e-e82b-4a4f-adf8-f09d4857b0e0",
// 		"32ed9040-1c14-4875-9ab3-c504dc7a962c",
// 		"5509e717-a4de-4457-b3b9-5b1676e591e4"}

// 	options := make([]ucs.MediaLink, 0, 4)

// 	for _, id := range ids {
// 		media := ucs.MediaLink{
// 			ID:  uuid.MustParse(id),
// 			URI: fmt.Sprintf("%s/quizzes/%s/image", c.BaseURL(), id),
// 		}
// 		options = append(options, media)
// 	}

// 	audio := ucs.MediaLink{
// 		ID:  uuid.MustParse("44e420ad-c5fd-4d26-9ed5-2838f2450147"),
// 		URI: fmt.Sprintf("%s/quizzes/%s/audio", c.BaseURL(), "44e420ad-c5fd-4d26-9ed5-2838f2450147"),
// 	}

// 	response := ucs.GetGameByIDResponse{
// 		Question: "What is this animal?",
// 		Audio:    audio,
// 		Images:   options,
// 	}

// 	return c.JSON(response)

// }

func (ctrl *gameController) ValidateAnswer(c *fiber.Ctx) error {
	return c.JSON("")
}

func (ctrl *gameController) NexQuiz(c *fiber.Ctx) error {
	return c.JSON("")
}

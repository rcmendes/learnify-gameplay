package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rcmendes/learnify/gameplay/core/entities"
	"github.com/rcmendes/learnify/gameplay/core/ucs/ports"
)

//TODO Evaluate the usage of Context for managing cancelling request for example.

//QuizController defines the endpoints of a Quiz controller.
type QuizController interface {
	ListAll(c *fiber.Ctx) error
	FindOneByID(c *fiber.Ctx) error
	DeleteByID(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	GetImageByID(c *fiber.Ctx) error
	GetAudioByID(c *fiber.Ctx) error
}

type quizController struct {
	findAll            ports.FindAllQuizzes
	findByCategoryName ports.FindQuizByCategoryName
	findQuiz           ports.FindQuiz
}

// NewQuizController creates a Quiz controller.
func NewQuizController(
	findAllQuizzesUC ports.FindAllQuizzes,
	findQuizByCategoryNameUC ports.FindQuizByCategoryName,
	findQuizUC ports.FindQuiz,
) QuizController {
	return &quizController{
		findAll:            findAllQuizzesUC,
		findByCategoryName: findQuizByCategoryNameUC,
		findQuiz:           findQuizUC,
	}
}

//QuizDTO defines the data returned when fetching a Quiz entity.
type QuizDTO struct {
	ID       uuid.UUID `json:"uuid"`
	Category string    `json:"category"`
	Palavra  string    `json:"palavra"`
	Mot      string    `json:"mot"`
}

func NewQuizDTO(q *entities.Quiz) *QuizDTO {
	return &QuizDTO{
		ID:       q.ID,
		Category: q.Category,
		Palavra:  q.Palavra,
		Mot:      q.Mot,
	}
}

//TODO Handle errors

func contextPath(c *fiber.Ctx) string {
	baseURL := c.BaseURL()
	context := c.Route().Path

	return baseURL + context + "/"
}

func (controller *quizController) ListAll(c *fiber.Ctx) error {
	category := c.Query("category", "")

	var err error
	var quizzes entities.QuizList
	if category == "" {
		quizzes, err = controller.findAll.FindAll()
		if err != nil {
			return err
		}
	} else {
		quizzes, err = controller.findByCategoryName.FindByCategoryName(category)
	}

	if err != nil {
		return err
	}

	var list []*QuizDTO
	for _, q := range quizzes {
		list = append(list, NewQuizDTO(q))
	}

	return c.JSON(list)
}

func (controller *quizController) FindOneByID(c *fiber.Ctx) error {
	//TODO handle missing or invalid data
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		//TODO Handle error
		return err
	}

	quiz, err := controller.findQuiz.FindByID(id)
	if err != nil {
		//TODO handle error
		return err
	}

	return c.JSON(quiz)
}

func (controller *quizController) DeleteByID(c *fiber.Ctx) error {
	return c.SendString("Delete a Quiz by UUID")
}

func (controller *quizController) Create(c *fiber.Ctx) error {
	// createQuizInput := services.CreateQuizInput{}

	// if err := c.BodyParser(&createQuizInput); err != nil {
	// 	c.SendStatus(http.StatusBadRequest)
	// 	return err
	// }

	// output, err := controller.quizSrv.Create(createQuizInput)
	// if err != nil {
	// 	c.SendStatus(http.StatusInternalServerError)
	// 	return err
	// }

	// return c.Status(http.StatusCreated).JSON(output)
	return c.SendString("Create a Quiz")
}

func (controller *quizController) GetAudioByID(c *fiber.Ctx) error {
	//TODO handle missing or invalid data
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return err
	}

	audio, err := controller.findQuiz.GetAudioByID(id)
	if err != nil {
		//TODO handle error
		return err
	}

	contentType := contentTypeFromAudio(audio)
	c.Response().Header.Add("Content-type", contentType)

	if _, err := c.Write(*audio.Data); err != nil {
		return err
	}

	return nil
}

func (controller *quizController) GetImageByID(c *fiber.Ctx) error {
	//TODO handle missing or invalid data
	idParam := c.Params("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		return err
	}

	image, err := controller.findQuiz.GetImageByID(id)
	if err != nil {
		//TODO handle error
		return err
	}

	contentType := contentTypeFromImage(image)
	c.Response().Header.Add("Content-type", contentType)

	if _, err := c.Write(*image.Data); err != nil {
		return err
	}

	return nil
}

func contentTypeFromImage(image *entities.MediaInfo) string {
	if image.Png() {
		return "image/png"
	}

	if image.Jpeg() {
		return "image/jpeg"
	}

	if image.Gif() {
		return "image/gif"
	}

	return "application/octet-stream"
}

func contentTypeFromAudio(audio *entities.MediaInfo) string {
	if audio.Mp3() {
		return "audio/mpeg"
	}

	if audio.Ogg() {
		return "audio/ogg"
	}

	return "application/octet-stream"
}

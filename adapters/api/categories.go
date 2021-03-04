package rest

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rcmendes/learnify-gameplay/core/entities"
	"github.com/rcmendes/learnify-gameplay/core/ucs/ports"
)

//CreateCategoryRequest defines the basic data to create a Category entity.
type CreateCategoryRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (req *CreateCategoryRequest) isValid() bool {
	return req.Name != ""
}

//CategoryDTOResponse defines the response structure of a Category.
type CategoryDTOResponse struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

func NewCategoryDTOResponse(category *entities.Category) *CategoryDTOResponse {
	return &CategoryDTOResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

func (input *CreateCategoryRequest) ToEntity() *entities.NewCategory {
	return &entities.NewCategory{
		Name:        input.Name,
		Description: input.Description,
	}
}

//CreateCategoryResponse defines the response data of a Category entity creation.
type CreateCategoryResponse struct {
	ID uuid.UUID `json:"id"`
}

//CategoryDTO defines the data returned when fetching a Category entity.
type CategoryDTO struct {
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
}

//CategoryController defines the endpoints of a Category controller.
type CategoryController interface {
	ListAll(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

type categoryController struct {
	create  ports.CreateCategory
	findAll ports.FindAllCategories
}

//TODO Handle errors

// NewCategoryController creates a Category Controller.
func NewCategoryController(create ports.CreateCategory, findAll ports.FindAllCategories) CategoryController {
	return &categoryController{
		create,
		findAll,
	}
}

func (ctrl *categoryController) ListAll(c *fiber.Ctx) error {
	categories, err := ctrl.findAll.FindAll()
	if err != nil {
		return err
	}

	list := make([]*CategoryDTOResponse, 0, len(categories))
	for _, c := range categories {
		list = append(list, NewCategoryDTOResponse(c))
	}

	return c.JSON(list)
}

func (ctrl *categoryController) Create(c *fiber.Ctx) error {
	payload := new(CreateCategoryRequest)

	if err := c.BodyParser(payload); err != nil || !payload.isValid() {
		return fiber.ErrBadRequest
	}

	if err := c.BodyParser(&payload); err != nil {
		c.SendStatus(http.StatusBadRequest)
		return err
	}

	id, err := ctrl.create.Create(*payload.ToEntity())
	//TODO Handle Error
	if err != nil {
		c.SendStatus(http.StatusInternalServerError)
		return err
	}

	output := CreateCategoryResponse{
		ID: *id,
	}

	return c.Status(http.StatusCreated).JSON(output)
}

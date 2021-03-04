package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/rcmendes/learnify-gameplay/adapters/repository/postgres/models"
	"github.com/rcmendes/learnify-gameplay/core/entities"
	"github.com/rcmendes/learnify-gameplay/core/ucs/ports"
)

type quizRepository struct {
	connectFn func() *pg.DB
}

//TODO Customize errors

//NewQuizPostgresRepository creates a Quiz repository instance.
func NewQuizPostgresRepository() ports.QuizRepository {
	return &quizRepository{
		connectFn: Connect,
	}
}

func (repo *quizRepository) ListAll() (entities.QuizList, error) {
	conn := repo.connectFn()

	var list []models.QuizModel

	if err := conn.Model(&list).Select(); err != nil {
		return nil, err
	}

	quizzes := *convertQuizModelListToQuizEntityList(&list)

	return quizzes, nil
}

func (repo *quizRepository) FindQuizByCategoryID(categoryID entities.CategoryID) (entities.QuizList, error) {
	conn := repo.connectFn()

	var list []models.QuizModel

	if err := conn.Model(&list).
		Where("category_id = ?", categoryID).
		Select(); err != nil {

		return nil, err
	}

	quizzes := *convertQuizModelListToQuizEntityList(&list)

	return quizzes, nil
}

func (repo *quizRepository) FindQuizzesSameCategory(id entities.QuizID) (entities.QuizList, error) {
	conn := repo.connectFn()

	//TODO improve this code for better performance. 2 Queries => 1 query.

	// TODO refactor for handle error
	quiz, err := repo.GetQuizByID(id)
	if err != nil {
		return nil, err
	}

	var list []models.QuizModel

	if err := conn.Model(&list).
		Where("category_id = ?", quiz.CategoryID).
		Select(); err != nil {

		return nil, err
	}

	quizzes := *convertQuizModelListToQuizEntityList(&list)

	return quizzes, nil
}

func (repo *quizRepository) GetQuizByID(id entities.QuizID) (*entities.Quiz, error) {
	conn := repo.connectFn()

	var model models.QuizModel

	if err := conn.Model(&model).
		Where("id = ?", id).
		First(); err != nil {

		return nil, err
	}

	quiz := model.ToEntity()

	return quiz, nil
}

func convertQuizModelListToQuizEntityList(list *[]models.QuizModel) *entities.QuizList {
	quizzes := make(entities.QuizList, 0, len(*list))
	for _, qm := range *list {
		quiz := qm.ToEntity()
		quizzes = append(quizzes, quiz)
	}

	return &quizzes
}

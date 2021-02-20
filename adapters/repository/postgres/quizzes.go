package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/rcmendes/learnify/gameplay/adapters/repository/postgres/models"
	"github.com/rcmendes/learnify/gameplay/core/entities"
	"github.com/rcmendes/learnify/gameplay/core/ucs/ports"
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
	defer conn.Close()

	var list []models.QuizModel

	if err := conn.Model(&list).Select(); err != nil {
		return nil, err
	}

	var quizzes entities.QuizList
	for _, qm := range list {
		quiz := qm.To()
		quizzes = append(quizzes, &quiz)
	}

	return quizzes, nil
}

func (repo *quizRepository) FindQuizByCategoryName(categoryName string) (entities.QuizList, error) {
	conn := repo.connectFn()
	defer conn.Close()

	var list []models.QuizModel

	if err := conn.Model(&list).
		Where("category = ?", categoryName).
		Select(); err != nil {

		return nil, err
	}

	var quizzes entities.QuizList
	for _, qm := range list {
		quiz := qm.To()
		quizzes = append(quizzes, &quiz)
	}

	return quizzes, nil
}

func (repo *quizRepository) FindQuizzesSameCategory(id entities.QuizID) (entities.QuizList, error) {
	conn := repo.connectFn()
	defer conn.Close()

	//TODO improve this code for better performance. 2 Queries => 1 query.

	// TODO refactor for handle error
	quiz, err := repo.GetQuizByID(id)
	if err != nil {
		return nil, err
	}

	var list []models.QuizModel

	if err := conn.Model(&list).
		Where("category = ?", quiz.Category).
		Select(); err != nil {

		return nil, err
	}

	var quizzes entities.QuizList
	for _, qm := range list {
		quiz := qm.To()
		quizzes = append(quizzes, &quiz)
	}

	return quizzes, nil
}

func (repo *quizRepository) GetQuizByID(id entities.QuizID) (*entities.Quiz, error) {
	conn := repo.connectFn()
	defer conn.Close()

	var model models.QuizModel

	if err := conn.Model(&model).
		Where("id = ?", id).
		First(); err != nil {

		return nil, err
	}

	quiz := model.To()

	return &quiz, nil
}

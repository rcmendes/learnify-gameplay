package ucs

import (
	"github.com/rcmendes/learnify/gameplay/core/entities"
	"github.com/rcmendes/learnify/gameplay/core/ucs/ports"
)

type findQuiz struct {
	quizRepo  ports.QuizRepository
	imageRepo ports.ImageRepository
	audioRepo ports.AudioRepository
}

//MakeFindQuiz create a FindQuiz instance.
func MakeFindQuiz(
	quizRepo ports.QuizRepository,
	imageRepo ports.ImageRepository,
	audioRepo ports.AudioRepository,
) ports.FindQuiz {

	return &findQuiz{
		quizRepo,
		imageRepo,
		audioRepo,
	}
}

//TODO Handle errors or create custom ones.

func (uc *findQuiz) FindByID(id entities.QuizID) (*entities.Quiz, error) {
	quiz, err := uc.quizRepo.GetQuizByID(id)

	if err != nil {
		//TODO Handle error
		return nil, err
	}

	return quiz, nil
}

func (uc *findQuiz) GetImageByID(id entities.QuizID) (*entities.MediaInfo, error) {
	quiz, err := uc.quizRepo.GetQuizByID(id)

	if err != nil {
		//TODO Handle error
		return nil, err
	}

	data, err := uc.imageRepo.GetImageByFilename(quiz.Image.Name)
	if err != nil {
		//TODO Handle error
		return nil, err
	}

	quiz.Image.Data = data

	return &quiz.Image, nil
}

func (uc *findQuiz) GetAudioByID(id entities.QuizID) (*entities.MediaInfo, error) {
	quiz, err := uc.quizRepo.GetQuizByID(id)

	if err != nil {
		//TODO Handle error
		return nil, err
	}

	data, err := uc.audioRepo.GetAudioByFilename(quiz.Audio.Name)
	if err != nil {
		//TODO Handle error
		return nil, err
	}

	quiz.Audio.Data = data

	return &quiz.Audio, nil
}

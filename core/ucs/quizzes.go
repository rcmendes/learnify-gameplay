package ucs

// type findAllQuizzes struct {
// 	quizRepo  ports.QuizRepository
// 	imageRepo storage.ImageRepository
// 	audioRepo storage.AudioRepository
// }

// //MakeFindAllQuizzes creates a List All Quizzes Use Case instance.
// func MakeFindAllQuizzes(
// 	quizRepo ports.QuizRepository,
// 	imageRepo storage.ImageRepository,
// 	audioRepo storage.AudioRepository,
// ) ports.FindAllQuizzes {

// 	return &findAllQuizzes{
// 		quizRepo,
// 		imageRepo,
// 		audioRepo,
// 	}
// }

// //TODO Handle errors or create custom ones.

// func (svc *quizService) ListQuizzesByCategoryName(categoryName string) (entities.QuizList, error) {
// 	list, err := svc.quizRepo.FindByCategoryName(categoryName)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return list, nil
// }

// func (svc *quizService) GetQuizByID(id entities.QuizID) (*entities.Quiz, error) {
// 	quiz, err := svc.quizRepo.GetQuizByID(id)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return quiz, nil
// }

// func (svc *quizService) FindQuizzesSameCategory(id entities.QuizID) (entities.QuizList, error) {
// 	list, err := svc.quizRepo.FindQuizzesSameCategory(id)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return list, nil
// }

// func (svc *quizService) GetQuizImageByID(id entities.QuizID, loadData bool) (*entities.MediaInfo, error) {
// 	quiz, err := svc.quizRepo.GetQuizByID(id)

// 	if err != nil {
// 		//TODO handle error
// 		return nil, err
// 	}

// 	if loadData {
// 		quiz.Image.Data, err = svc.imageRepo.GetImageByFilename(quiz.Image.Name)

// 		if err != nil {
// 			//TODO handle error
// 			return nil, err
// 		}
// 	}

// 	return &quiz.Image, nil
// }

// func (svc *quizService) GetQuizAudioByID(id entities.QuizID, loadData bool) (*entities.MediaInfo, error) {
// 	quiz, err := svc.quizRepo.GetQuizByID(id)

// 	if err != nil {
// 		//TODO handle error
// 		return nil, err
// 	}

// 	if loadData {
// 		quiz.Audio.Data, err = svc.audioRepo.GetAudioByFilename(quiz.Audio.Name)

// 		if err != nil {
// 			//TODO handle error
// 			return nil, err
// 		}
// 	}
// 	return &quiz.Audio, nil
// }

package filesystem

import (
	"io/ioutil"
	"path/filepath"

	"github.com/rcmendes/learnify/gameplay/core/ucs/ports"
)

type imageFSRepository struct {
	basePath string
}

type audioFSRepository struct {
	basePath string
}

//TODO Customize errors

// NewImageFSRepository creates a File System based Image repository instance.
func NewImageFSRepository(basePath string) ports.ImageRepository {
	return &imageFSRepository{basePath: basePath}
}

// NewAudioFSRepository creates a File System based Audio repository instance.
func NewAudioFSRepository(basePath string) ports.AudioRepository {
	return &audioFSRepository{basePath: basePath}
}

func (repo *imageFSRepository) GetImageByFilename(filename string) (*[]byte, error) {
	filePath := filepath.Join(repo.basePath, filename)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		//TODO Handle error
		return nil, err
	}

	return &data, nil
}

func (repo *audioFSRepository) GetAudioByFilename(filename string) (*[]byte, error) {
	filePath := filepath.Join(repo.basePath, filename)

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		//TODO Handle error
		return nil, err
	}

	return &data, nil
}

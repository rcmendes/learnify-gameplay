package entities

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

//QuizID defines the type of the ID of a Quiz
type QuizID = uuid.UUID

//Quiz defines the properties of a Quiz database entity.
type Quiz struct {
	ID       QuizID
	Category string
	Palavra  string
	Mot      string
	Image    MediaInfo
	Audio    MediaInfo
}

type QuizList = []*Quiz

//MediaInfo defines the data structure information of a Media type.
type MediaInfo struct {
	Name string
	Data *[]byte
}

func (q *Quiz) String() string {
	return fmt.Sprintf("Quiz <category=%s, "+
		"palavra=%s, mot=%s, image=%s, audio=%s>",
		q.Category, q.Palavra, q.Mot, q.Image.Name, q.Audio.Name)
}

func extractExtensionFromFilename(filename string) string {
	extension := filepath.Ext(filename)
	if strings.Index(extension, ".") == 0 {
		extension = extension[1:]
	}

	return strings.ToLower(extension)
}

func (mi *MediaInfo) Unknown() bool {
	extension := extractExtensionFromFilename(mi.Name)
	return extension == ""
}

func (mi *MediaInfo) Png() bool {
	extension := extractExtensionFromFilename(mi.Name)
	return extension == "png"
}

func (mi *MediaInfo) Jpeg() bool {
	extension := extractExtensionFromFilename(mi.Name)
	return extension == "jpeg" || extension == "jpg"
}

func (mi *MediaInfo) Gif() bool {
	extension := extractExtensionFromFilename(mi.Name)
	return extension == "gif"
}

func (mi *MediaInfo) Mp3() bool {
	extension := extractExtensionFromFilename(mi.Name)
	return extension == "mp3"
}

func (mi *MediaInfo) Ogg() bool {
	extension := extractExtensionFromFilename(mi.Name)
	return extension == "ogg"
}

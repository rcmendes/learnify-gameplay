package ports

//ImageRepository defines the contract of an Image Repository.
type ImageRepository interface {
	GetImageByFilename(filename string) (*[]byte, error)
}

//AudioRepository defines the contract of an Audio Repository.
type AudioRepository interface {
	GetAudioByFilename(filename string) (*[]byte, error)
}

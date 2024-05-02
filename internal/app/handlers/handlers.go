package handlers

type Storage interface {
	GetByURLPath(string) string
	SetByOriginalURL(string) string
}

type Handlers struct {
	storage Storage
}

func NewHandlers(storage Storage) *Handlers {
	return &Handlers{
		storage: storage,
	}
}

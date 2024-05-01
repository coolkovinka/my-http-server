package storage

type Storage struct{}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) GetByURLPath(_ string) string {
	return ""
}

func (s *Storage) SetByOriginalURL(_ string) string {
	return ""
}

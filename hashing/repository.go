package hashing

import "errors"

type Repository interface {
	CheckHashExists(hash string) (bool, error)
	GetHashData(hash string) (string, error)
	HashData(data string) (string, error)
}

// InMemoryRepository представляет репозиторий в памяти
type InMemoryRepository struct {
	hashes map[string]string
}

// NewInMemoryRepository создает новый экземпляр репозитория в памяти
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		hashes: make(map[string]string),
	}
}

// GetHashData реализует метод GetHashData интерфейса Repository
func (r *InMemoryRepository) GetHashData(hash string) (string, error) {
	data, ok := r.hashes[hash]
	if !ok {
		return "", errors.New("hash not found")
	}
	return data, nil
}

// HashData реализует метод HashData интерфейса Repository
func (r *InMemoryRepository) HashData(data string) (string, error) {
	// Ваш код для создания хэша из переданных данных
	// В этом примере просто возвращаем данные как хэш
	return data, nil
}

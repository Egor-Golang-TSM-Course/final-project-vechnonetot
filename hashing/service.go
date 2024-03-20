package hashing

import (
	"context"
)

// CheckHashRequest представляет собой запрос на проверку существования хэша
type CheckHashRequest struct {
	Hash string
}

// CheckHashResponse представляет собой ответ на запрос проверки существования хэша
type CheckHashResponse struct {
	Exists bool
}

// GetHashRequest представляет запрос на получение хэша
type GetHashRequest struct {
	Hash string
}

// GetHashResponse представляет ответ на запрос получения хэша
type GetHashResponse struct {
	Data string
}

// CreateHashRequest представляет запрос на создание хэша
type CreateHashRequest struct {
	Data string
}

// CreateHashResponse представляет ответ на запрос создания хэша
type CreateHashResponse struct {
	Hash string
}

// HashingService представляет собой сервис для работы с хешированием
type HashingService struct {
	repo Repository
}

// CheckHash проверяет существует ли хэш полезных данных в репозитории
func (s *HashingService) CheckHash(ctx context.Context, req *CheckHashRequest) (*CheckHashResponse, error) {
	// Вызываем метод CheckHashExists репозитория для проверки существования хэша
	exists, err := s.repo.CheckHashExists(req.Hash)
	if err != nil {
		return nil, err
	}

	// Возвращаем ответ клиенту в виде структуры CheckHashResponse
	return &CheckHashResponse{Exists: exists}, nil
}

// GetHash возвращает хэш существующих полезных данных из репозитория
func (s *HashingService) GetHash(ctx context.Context, req *GetHashRequest) (*GetHashResponse, error) {
	// Вызываем метод GetHashData репозитория для получения данных по хэшу
	data, err := s.repo.GetHashData(req.Hash)
	if err != nil {
		return nil, err
	}

	// Возвращаем ответ клиенту в виде структуры GetHashResponse
	return &GetHashResponse{Data: data}, nil
}

// CreateHash создает и сохраняет хэш для новых полезных данных в репозитории
func (s *HashingService) CreateHash(ctx context.Context, req *CreateHashRequest) (*CreateHashResponse, error) {
	// Вызываем метод HashData репозитория для создания хэша для переданных данных
	hash, err := s.repo.HashData(req.Data)
	if err != nil {
		return nil, err
	}

	// Возвращаем ответ клиенту в виде структуры CreateHashResponse
	return &CreateHashResponse{Hash: hash}, nil
}

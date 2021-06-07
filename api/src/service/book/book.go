package book

import (
	"candidate/api/src/dto/response"
	"candidate/api/src/models"
	"candidate/api/src/repository/book"
	"context"
	"net/http"
)

// Service is the interface spec expected of a book service
type Service interface {
	GetBooksAvailableForSale(ctx context.Context) ([]models.Book, *response.Error)
	GetAllBooks(ctx context.Context) ([]models.Book, *response.Error)
}

// Manager is the implementation of book service
type Manager struct {
	bookRepository book.Repository
}

// NewManager initializes book service
func NewManager(bookRepository book.Repository) *Manager {
	return &Manager{
		bookRepository: bookRepository,
	}
}

// GetBooksAvailableForSale service function to get BookAvailableForSale
func (m *Manager) GetBooksAvailableForSale(ctx context.Context) ([]models.Book, *response.Error) {
	books, err := m.bookRepository.GetBooksAvailableForSale(ctx)
	if err != nil {
		return nil, &response.Error{Code: http.StatusInternalServerError, Description: err.Error()}
	}
	return books, nil
}

// GetAllBooks service function to get All Books
func (m *Manager) GetAllBooks(ctx context.Context) ([]models.Book, *response.Error) {
	books, err := m.bookRepository.GetAllBooks(ctx)
	if err != nil {
		return nil, &response.Error{Code: http.StatusInternalServerError, Description: err.Error()}
	}
	return books, nil
}

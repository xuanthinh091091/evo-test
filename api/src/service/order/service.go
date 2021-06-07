package order

import (
	"candidate/api/src/dto/response"
	"candidate/api/src/models"
	"candidate/api/src/repository/order"
	"context"
	"net/http"
)

// Service is the interface spec expected of a order service
type Service interface {
	OrderBook(ctx context.Context, userID, bookID int64, quantity int, shippingDetais string) *response.Error
	GetAllOrders(ctx context.Context) ([]models.Order, *response.Error)
}

// Manager is the implementation of order service
type Manager struct {
	orderRepository order.Repository
}

// NewManager initializes order service
func NewManager(orderRepository order.Repository) *Manager {
	return &Manager{
		orderRepository: orderRepository,
	}
}

// OrderBook make book order
func (m *Manager) OrderBook(ctx context.Context, userID, bookID int64, quantity int, shippingDetais string) *response.Error {
	err := m.orderRepository.OrderBook(ctx, userID, bookID, quantity, shippingDetais)
	if err != nil {
		return &response.Error{Code: http.StatusInternalServerError, Description: err.Error()}
	}

	return nil
}

// GetAllOrders service function to get AllOrders
func (m *Manager) GetAllOrders(ctx context.Context) ([]models.Order, *response.Error) {
	books, err := m.orderRepository.GetAllOrders(ctx)
	if err != nil {
		return nil, &response.Error{Code: http.StatusInternalServerError, Description: err.Error()}
	}
	return books, nil
}

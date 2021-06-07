package user

import (
	"candidate/api/src/dto/response"
	"candidate/api/src/models"
	"candidate/api/src/repository/user"
	"context"
	"net/http"
)

// Service is the interface spec expected of a user service
type Service interface {
	GetUser(ctx context.Context, userName, passWord string) (models.User, *response.Error)
	GetUserByID(ctx context.Context, userID int64) (models.User, error)
}

// Manager is the implementation of user service
type Manager struct {
	userRepository user.Repository
}

// NewManager initializes user service
func NewManager(userRepository user.Repository) *Manager {
	return &Manager{
		userRepository: userRepository,
	}
}

// GetUser get user by user name and password
func (m *Manager) GetUser(ctx context.Context, userName, passWord string) (models.User, *response.Error) {
	user, err := m.userRepository.GetUser(ctx, userName, passWord)
	if err != nil {
		return models.User{}, &response.Error{Code: http.StatusInternalServerError, Description: err.Error()}
	}

	return user, nil
}

// GetUserByID get user by id
func (m *Manager) GetUserByID(ctx context.Context, userID int64) (models.User, error) {
	user, err := m.userRepository.GetUserByID(ctx, userID)
	if err != nil {
		return models.User{}, &response.Error{Code: http.StatusInternalServerError, Description: err.Error()}
	}

	return user, nil
}

package user

import (
	"api-go/internal/domain"
	"context"
	"log"
)

type (
	Service interface {
		Create(ctx context.Context, firstName, lastName, email string) (*domain.User, error)
		GetAll(ctx context.Context) ([]domain.User, error)
		GetUser(ctx context.Context, id uint64) (*domain.User, error)
	}

	service struct {
		log  *log.Logger
		repo Repository
	}
)

// Create implements Service.
func (s *service) Create(ctx context.Context, firstName string, lastName string, email string) (*domain.User, error) {
	user := &domain.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// GetAll implements Service.
func (s *service) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := s.repo.GetAll(ctx)

	if err != nil {
		return nil, err
	}
	return users, nil
}

func NewService(l *log.Logger, repo Repository) Service {
	return &service{
		log:  l,
		repo: repo,
	}
}

func (s service) GetUser(ctx context.Context, id uint64) (*domain.User, error) {
	return s.repo.GetUser(ctx, id)
}

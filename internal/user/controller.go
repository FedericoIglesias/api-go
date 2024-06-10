package user

import (
	"context"
)

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	Endpoints struct {
		Create     Controller
		GetAll     Controller
		GetUser    Controller
		UpdateUser Controller
	}

	GetReq struct {
		ID uint64
	}
	CreateReq struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
	}
	UpdateReq struct {
		ID        uint64
		FirstName *string `json:"firstName"`
		LastName  *string `json:"lastName"`
		Email     *string `json:"email"`
	}
)

func MakeEndpoits(ctx context.Context, s Service) Endpoints {
	return Endpoints{
		Create:     makeCreateEndpoints(s),
		GetAll:     makeGetAllEndpoints(s),
		GetUser:    makeGetUserEndpoints(s),
		UpdateUser: makeUpdateUserEndpoint(s),
	}
}

func makeGetAllEndpoints(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users, err := s.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		return users, nil
	}
}

func makeCreateEndpoints(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(CreateReq)

		if req.FirstName == "" {
			return nil, ErrFirstNameRequired
		}
		if req.LastName == "" {
			return nil, ErrLasNameRequired
		}
		if req.Email == "" {
			return nil, ErrEmailRequired
		}
		user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}

func makeGetUserEndpoints(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetReq)
		return s.GetUser(ctx, req.ID)
	}
}

func makeUpdateUserEndpoint(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateReq)

		if req.FirstName != nil && *req.FirstName == "" {
			return nil, ErrFirstNameRequired
		}

		if req.LastName != nil && *req.LastName == "" {
			return nil, ErrLasNameRequired
		}

		if req.Email != nil && *req.Email == "" {
			return nil, ErrEmailRequired
		}

		err := s.UpdateUser(ctx, req.ID, req.FirstName, req.LastName, req.Email)

		if err != nil {
			return nil, err
		}
		return nil, nil
	}

}

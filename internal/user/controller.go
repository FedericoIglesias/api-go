package user

import (
	"context"
	"errors"
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
		Create Controller
		GetAll Controller
		GetUser    Controller
	}

	GetReq struct {
		ID uint64
	}
	CreateReq struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
	}
)

func MakeEndpoits(ctx context.Context, s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoints(s),
		GetAll: makeGetAllEndpoints(s),
		GetUser:    makeGetUserEndpoints(s),
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
			return nil, errors.New("fist name is required")
		}
		if req.LastName == "" {
			return nil, errors.New("lastname is required")
		}
		if req.Email == "" {
			return nil, errors.New("email is required")
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

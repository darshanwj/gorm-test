package internal

import (
	"context"
	"darshanwj/gorm-test/internal/model"

	"github.com/go-kit/kit/endpoint"
)

type homeRequest struct{}

type homeResponse struct {
	Message string `json:"msg"`
}

type getUserRequest struct{}

type getUserResponse struct {
	User model.User `json:"user"`
}

func makeHomeEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		return homeResponse{Message: "gokit test service"}, nil
	}
}

func makeGetUserEndpoint(us UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		user := us.GetUser(ctx, 10)
		return getUserResponse{User: user}, nil
	}
}

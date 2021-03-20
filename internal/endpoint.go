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

func makeHomeEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		return homeResponse{Message: "gokit test service"}, nil
	}
}

type getUserRequest struct {
	Id int
}

type getUserResponse struct {
	User model.User `json:"user"`
}

func makeGetUserEndpoint(us UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		gur := request.(getUserRequest)
		user := us.GetUser(ctx, uint(gur.Id))
		return getUserResponse{User: user}, nil
	}
}

type getUsersRequest struct{}

type getUsersResponse struct {
	Users []model.User `json:"users"`
}

func makeGetUsersEndpoint(us UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		users := us.GetUsers(ctx)
		return getUsersResponse{Users: users}, nil
	}
}

type createUserRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type createUserResponse struct {
	User model.User `json:"user"`
}

func makeCreateUserEndpoint(us UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		cur := request.(createUserRequest)
		user := us.CreateUser(ctx, cur)
		return createUserResponse{User: user}, nil
	}
}

type getCommentsRequest struct {
	UserId int
}

type getCommentsResponse struct {
	Comments []model.Comment `json:"comments"`
}

func makeGetCommentsEndpoint(us UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		gcr := request.(getCommentsRequest)
		comments := us.GetComments(ctx, gcr)
		return getCommentsResponse{Comments: comments}, nil
	}
}

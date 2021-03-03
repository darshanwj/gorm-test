package internal

import (
	"context"

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

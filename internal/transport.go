package internal

import (
	"context"
	"encoding/json"
	"net/http"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(us UserService) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/").Handler(transport.NewServer(makeHomeEndpoint(), decodeHomeRequest, encodeResponse))
	r.Methods("GET").Path("/user").Handler(transport.NewServer(makeGetUserEndpoint(us), decodeGetUserRequest, encodeResponse))
	return r
}

func decodeHomeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return homeRequest{}, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return getUserRequest{}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

package internal

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(us UserService) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/").Handler(transport.NewServer(makeHomeEndpoint(), decodeHomeRequest, encodeResponse))
	r.Methods("GET").Path("/user/{id}").Handler(transport.NewServer(makeGetUserEndpoint(us), decodeGetUserRequest, encodeResponse))
	r.Methods("GET").Path("/users").Handler(transport.NewServer(makeGetUsersEndpoint(us), decodeGetUsersRequest, encodeResponse))
	r.Methods("POST").Path("/user").Handler(transport.NewServer(makeCreateUserEndpoint(us), decodeCreateUserRequest, encodeResponse))
	return r
}

func decodeHomeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return homeRequest{}, nil
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userId, ok := vars["id"]
	if !ok {
		return nil, errors.New("bad route")
	}

	id, err := strconv.Atoi(userId)
	if err != nil {
		return nil, err
	}

	return getUserRequest{Id: id}, nil
}

func decodeGetUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return getUsersRequest{}, nil
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req createUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

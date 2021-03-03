package internal

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPHandler(us UserService) http.Handler {
	r := mux.NewRouter()
	r.Methods("GET").Path("/").Handler(httptransport.NewServer(makeHomeEndpoint(), decodeHomeRequest, encodeResponse))
	return r
}

func decodeHomeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return homeRequest{}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

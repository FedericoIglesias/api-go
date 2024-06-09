package handler

import (
	"api-go/internal/user"
	"api-go/pkg/transport"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func NewUserHTTPServer(ctx context.Context, router *http.ServeMux, endpoint user.Endpoints) {
	router.HandleFunc("/users", UserServer(ctx, endpoint))
}

func UserServer(ctx context.Context, enpoints user.Endpoints) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		tran := transport.New(w, r, ctx)

		switch r.Method {
		case http.MethodGet:
			tran.Server(
				transport.Endpoint(enpoints.GetAll),
				decodeGetAllUser,
				encodeResponse,
				encodeError,
			)
			return 
		case http.MethodPost:
			tran.Server(
				transport.Endpoint(enpoints.Create),
				decodeCreateUser,
				encodeResponse,
				encodeError,
			)
			return
		default:
			InvalidMethod(w)
		}
	}
}

func decodeGetAllUser(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeCreateUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req user.CreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("Invalid request format: '%v'", err.Error())
	}
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	data, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	status := http.StatusOK
	w.WriteHeader(status)
	w.Header().Set("Conten-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `{"sataus": %d,"data":%s}`, status, data)
	return nil
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	status := http.StatusInternalServerError
	w.WriteHeader(status)
	w.Header().Set("Conten-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, `{"sataus": %d,"message":%s}`, status, err.Error())
}

func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"sataus": %d,"message": "method doesn't exist"}`, status)
}

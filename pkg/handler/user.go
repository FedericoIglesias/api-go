package handler

import (
	"api-go/internal/user"
	"api-go/pkg/transport"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func NewUserHTTPServer(ctx context.Context, router *http.ServeMux, endpoint user.Endpoints) {
	router.HandleFunc("/users/", UserServer(ctx, endpoint))
}

func UserServer(ctx context.Context, enpoints user.Endpoints) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		url := r.URL.Path
		log.Println(r.Method, ": ", url)

		path, pathSize := transport.Clean(url)

		params := make(map[string]string)

		if pathSize == 4 && path[2] != "" {
			params["userID"] = path[2]
		}

		ctx = context.WithValue(ctx, "params", params)

		tran := transport.New(w, r, ctx)

		var end user.Controller
		var deco func(ctx context.Context, r *http.Request) (interface{}, error)

		switch r.Method {
		case http.MethodGet:
			switch pathSize {
			case 3:
				end = enpoints.GetAll
				deco = decodeGetAllUser
			case 4:
				end = enpoints.GetUser
				deco = decodeGetUser
			}
		case http.MethodPost:
			switch pathSize {
			case 3:
				end = enpoints.Create
				deco = decodeCreateUser
			}
		}
		if end != nil && deco != nil {
			tran.Server(
				transport.Endpoint(end),
				deco,
				encodeResponse,
				encodeError,
			)
		} else {
			InvalidMethod(w)
		}
	}
}

func decodeGetUser(ctx context.Context, r *http.Request) (interface{}, error) {
	params := ctx.Value("params").(map[string]string)

	id, err := strconv.ParseUint(params["userID"],10,61)
	
	if err!= nil{
		return nil, err
	}
	
	fmt.Println(params)
	fmt.Println(params["userID"])


	return user.GetReq{
		ID:id,
	}, nil
}
func decodeGetAllUser(ctx context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func decodeCreateUser(ctx context.Context, r *http.Request) (interface{}, error) {
	var req user.CreateReq
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, fmt.Errorf("invalid request format: '%v'", err.Error())
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

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	const port int = 8080
	http.HandleFunc("/users", UserServer)
	fmt.Printf("Server up in port: %d\n", port)
	log.Fatal((http.ListenAndServe(":8080", nil)))
}

type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

var users []User
var maxID uint64

func init() {
	users = []User{
		{
			ID:        1,
			FirstName: "Pepe",
			LastName:  "Coco",
			Email:     "@algo",
		},
		{
			ID:        2,
			FirstName: "Cacho",
			LastName:  "Goxila",
			Email:     "@otraCosa",
		},
		{
			ID:        3,
			FirstName: "Armando",
			LastName:  "Banquito",
			Email:     "@otroGato",
		},
	}
	maxID = 3
}

func UserServer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetAllUser(w)
	case http.MethodPost:
		decode := json.NewDecoder(r.Body)
		var u User
		if err := decode.Decode(&u); err != nil {
			MsgResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		PostUser(w, u)
	default:
		InvalidMethod(w)
	}
}

func GetAllUser(w http.ResponseWriter) {
	DataResponse(w, http.StatusOK, users)
}

func DataResponse(w http.ResponseWriter, status int, users interface{}) {
	value, err := json.Marshal(users)
	if err != nil {
		MsgResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"sataus": %d,"data":%s}`, status, value)
}

func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"sataus": %d,"message":%s}`, status, message)
}

func PostUser(w http.ResponseWriter, data interface{}) {
	user := data.(User)

	if user.FirstName == ""  {
		MsgResponse(w, http.StatusBadRequest, "Fist name invalid")
		return
	}
	if user.LastName == ""  {
		MsgResponse(w, http.StatusBadRequest, "Lastname invalid")
		return
	}
	if user.Email == ""  {
		MsgResponse(w, http.StatusBadRequest, "Email invalid")
		return
	}
	maxID++
	user.ID = maxID
	users = append(users, user)
	DataResponse(w, http.StatusCreated, user)
}

func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusNotFound
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"sataus": %d,"message": "method doesn't exist"}`, status)
}

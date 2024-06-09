package main

import (
	"api-go/internal/domain"
	"api-go/internal/user"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	server:= http.NewServeMux()

	db := user.DB{
		Users: []domain.User{
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
		},
		MaxUserID: 3,
	}

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	repo := user.NewRepo(db, logger)
	service:= user.NewService(logger,repo)
	
	ctx := context.Background()

	server.HandleFunc("/users", user.MakeEndpoits(ctx, service))
	fmt.Println("Server up in port: 8080")
	log.Fatal((http.ListenAndServe(":8080", server)))
}
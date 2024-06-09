package main

import (
	"api-go/internal/user"
	"api-go/pkg/bootstrap"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	server:= http.NewServeMux()

	
	repo := user.NewRepo(bootstrap.NewDB(), bootstrap.NewLogger())
	service:= user.NewService(bootstrap.NewLogger(),repo)
	
	ctx := context.Background()

	server.HandleFunc("/users", user.MakeEndpoits(ctx, service))
	fmt.Println("Server up in port: 8080")
	log.Fatal((http.ListenAndServe(":8080", server)))
}
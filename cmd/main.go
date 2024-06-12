package main

import (
	"api-go/internal/user"
	"api-go/pkg/bootstrap"
	"api-go/pkg/handler"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	_ = godotenv.Load()

	server:= http.NewServeMux()

	port := os.Getenv("PORT")

	db, err:= bootstrap.NewDB()

	if err!= nil{
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping();err != nil{
		log.Fatal(err)
	}
	
	repo := user.NewRepo(db, bootstrap.NewLogger())
	service:= user.NewService(bootstrap.NewLogger(),repo)
	
	ctx := context.Background()
	handler.NewUserHTTPServer(ctx, server, user.MakeEndpoits(ctx, service))

	
	fmt.Println("Server up in port: ", port)
	log.Fatal((http.ListenAndServe(":8080", server)))
}
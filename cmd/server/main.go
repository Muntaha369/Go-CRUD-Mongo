package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Muntaha369/Go-CRUD-Mongo/internal/db"
	"github.com/Muntaha369/Go-CRUD-Mongo/internal/repository"
)

func main() {
	client := db.DB{
		Db: db.ConnectDB(),
	}
	router := http.NewServeMux()

	operation := repository.NewService(client)

	router.HandleFunc("GET /api/sendall", operation.GetAll())
	router.HandleFunc("POST /api/createnew", operation.WriteTO())

	server := http.Server{
		Addr:    "localhost:3004",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("There is an error")
	}
	fmt.Println("Server is up and running at :", server.Addr)
}
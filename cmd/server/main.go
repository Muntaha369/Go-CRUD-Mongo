package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Muntaha369/Go-CRUD-Mongo/internal/db"
)

func main() {
	client := db.DB{
		Db: db.ConnectDB(),
	}
	router := http.NewServeMux()

	router.HandleFunc("GET /api/sendall", client.GetAll())
	router.HandleFunc("POST /api/createnew", client.WriteTO())

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
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Muntaha369/Go-CRUD-Mongo/internal/config"
	"github.com/Muntaha369/Go-CRUD-Mongo/internal/db"
	"github.com/Muntaha369/Go-CRUD-Mongo/internal/repository"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	client := db.DB{
		Db: db.ConnectDB(cfg.Database.URI),
	}
	router := http.NewServeMux()

	operation := repository.NewService(client)

	router.HandleFunc("GET /api/sendall", operation.GetAll())
	router.HandleFunc("POST /api/createnew", operation.CreateNew())
	router.HandleFunc("GET /api/getuser/{name}", operation.GetByName())
	router.HandleFunc("PUT /api/updateuser", operation.UpdateUser())

	server := http.Server{
		Addr:    ":" + strconv.Itoa(cfg.Server.Port),
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("There is an error")
	}
	fmt.Println("Server is up and running at :", server.Addr)
}
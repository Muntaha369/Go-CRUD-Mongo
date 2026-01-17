package repository

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Muntaha369/Go-CRUD-Mongo/internal/db"
	"github.com/Muntaha369/Go-CRUD-Mongo/internal/model"
	response "github.com/Muntaha369/Go-CRUD-Mongo/internal/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Service struct{
	DB db.DB
}

func NewService(db db.DB) *Service{
	return &Service{DB: db}
}

func (h *Service) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cursor, err := h.DB.Db.Collection("users").Find(context.TODO(), bson.M{})
		if err != nil {
			response.WriteJson(w, 404, response.Genralerror(err))

		}
		var users []bson.M
		if err = cursor.All(context.TODO(), &users); err != nil {
			response.WriteJson(w, 404, response.Genralerror(errors.New("Cant establish a connection")))
			return
		}
		response.WriteJson(w, 200, users)
	}	
}

func (h *Service) WriteTO() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			response.WriteJson(w, 404, response.Genralerror(err))
		}
		_, err := h.DB.Db.Collection("users").InsertOne(context.TODO(), user)

		if err != nil {
			response.WriteJson(w, 404, response.Genralerror(err))
		}
		response.WriteJson(w, 200, map[string]string{"message": "User created successfully"})
	}
}
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

type Service struct {
	DB db.DB
}

type UpdatedUser struct {
	Newname string
	Oldname string
}

func NewService(db db.DB) *Service {
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

func (h *Service) CreateNew() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			response.WriteJson(w, 404, response.Genralerror(err))
		}
		_, err := h.DB.Db.Collection("users").InsertOne(context.TODO(), user)

		if err != nil {
			response.WriteJson(w, 404, response.Genralerror(err))
			return
		}
		response.WriteJson(w, 200, map[string]string{"message": "User created successfully"})
	}
}

func (h *Service) GetByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")

		res := h.DB.Db.Collection("users").FindOne(context.TODO(), name)

		if res != nil {
			response.WriteJson(w, 404, map[string]string{"message": "Cant seem to find a user with taht name"})
			return
		}

		response.WriteJson(w, 200, map[string]string{"message": "User exists"})
	}
}

func (h *Service) UpdateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user UpdatedUser

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			response.WriteJson(w, 404, response.Genralerror(err))
		}

		update := bson.M{
			"$set": bson.M{
				"name": user.Newname,
			},
		}

		filter := bson.M{"name": user.Oldname}

		_, err := h.DB.Db.Collection("users").UpdateOne(context.TODO(), filter, update)

		if err != nil {
			response.WriteJson(w, 404, response.Genralerror(err))
			return
		}

		response.WriteJson(w, 203, map[string]string{"message": "user updated successfully"})
	}
}

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
		var user model.User

		filter := bson.M{"name": name}
		err := h.DB.Db.Collection("users").FindOne(context.TODO(), filter).Decode(&user)

		if err != nil {
			if err.Error() == "mongo: no documents in result" {
				response.WriteJson(w, http.StatusNotFound, response.Genralerror(errors.New("user not found")))
				return
			}
			response.WriteJson(w, http.StatusInternalServerError, response.Genralerror(err))
			return
		}

		response.WriteJson(w, http.StatusOK, user)
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

func (h *Service) DeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user struct {
			Name string `json:"name"`
		}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			response.WriteJson(w, 400, response.Genralerror(err))
			return
		}

		if user.Name == "" {
			response.WriteJson(w, 400, map[string]string{"error": "name is required"})
			return
		}

		filter := bson.M{"name": user.Name}

		result, err := h.DB.Db.Collection("users").DeleteOne(context.TODO(), filter)

		if err != nil {
			response.WriteJson(w, 500, response.Genralerror(err))
			return
		}

		if result.DeletedCount == 0 {
			response.WriteJson(w, 404, map[string]string{"error": "user not found"})
			return
		}

		response.WriteJson(w, 200, map[string]string{"message": "user deleted"})
	}
}

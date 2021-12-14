package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetTaskResponse struct {
	ID   string `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
}

func (h *Handler) handlerGetTask(w http.ResponseWriter, r *http.Request) {
	vars :=  mux.Vars(r)

	objectID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	res := h.mongo.Database("db").Collection("tasks").FindOne(r.Context(), bson.D{
		{"_id",objectID},
	})

	err = res.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	var t GetTaskResponse
	err = res.Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	bb, err := json.Marshal(t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bb)
}

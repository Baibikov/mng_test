package main

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) handlerDeleteTask(w http.ResponseWriter, r *http.Request) {
	vars :=  mux.Vars(r)

	objectID, err := primitive.ObjectIDFromHex(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	res, err := h.mongo.Database("db").Collection("tasks").DeleteOne(r.Context(), bson.D{
		{"_id",objectID},
	}, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}


	w.Write([]byte(strconv.Itoa(int(res.DeletedCount))))
}


package main

import (
	"encoding/json"
	"net/http"
)


type CreateRequest struct {
	Name string `bson:"name" json:"name"`
}

func (h *Handler) handlerPostTaskHandler(w http.ResponseWriter, r *http.Request) {
	var t CreateRequest

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusBadRequest,
		)
		return
	}

	_, err = h.mongo.Database("db").Collection("tasks").InsertOne(r.Context(), t)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}

	w.Write([]byte("OK"))
}

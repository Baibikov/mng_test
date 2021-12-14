package main

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

type GetTasksResponse []GetTaskResponse

func (h *Handler) handlerGetTasks(w http.ResponseWriter, r *http.Request) {

	c, err := h.mongo.Database("db").Collection("tasks").Find(r.Context(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	err = c.Err()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	var tasks GetTasksResponse
	defer c.Close(r.Context())
	for c.Next(r.Context()) {
		var t GetTaskResponse
		err = c.Decode(&t)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		tasks = append(tasks, t)
	}

	bb, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bb)
}

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Handler struct {
	mongo *mongo.Client
}

func NewHandler(client *mongo.Client) *Handler {
	return &Handler{
		mongo: client,
	}
}


func main() {
	ctx  := context.Background()
	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://root:123@localhost:27017"),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Disconnect(ctx)

	h := NewHandler(client)

	r := mux.NewRouter()

	r.HandleFunc("/api/task/{id}", h.handlerGetTask).Methods(http.MethodGet)
	r.HandleFunc("/api/task/{id}", h.handlerDeleteTask).Methods(http.MethodDelete)
	r.HandleFunc("/api/task", h.handlerPostTaskHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/tasks", h.handlerGetTasks)

	err = http.ListenAndServe(":9090", r)
	if err != nil {
		log.Fatalln(err)
	}
}
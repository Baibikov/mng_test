package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	Task Task
}

type Task interface {
	Create(ctx context.Context) error
	Get(ctx context.Context) error
	GetAll(ctx context.Context) error
}

func New(client *mongo.Client) *Storage {
	return &Storage{}
}

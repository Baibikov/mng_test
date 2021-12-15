package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"mongou/internal/app/models"
	"mongou/internal/app/repository/mg"
)

type Storage struct {
	Task Task
}

type Task interface {
	Create(ctx context.Context, task models.Task) (string, error)
	Get(ctx context.Context, id string) (models.Task, error)
	GetAll(ctx context.Context) ([]models.Task, error)
}

func New(client *mongo.Client) *Storage {
	return &Storage{
		Task: mg.New(client),
	}
}

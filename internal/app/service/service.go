package service

import (
	"context"

	"mongou/internal/app/models"
	"mongou/internal/app/repository"
)

type UseCase struct {
	storage *repository.Storage
}

func New(storage *repository.Storage) *UseCase {
	return &UseCase{storage: storage}
}

func (u *UseCase) Create(ctx context.Context, task models.Task) (string, error) {
	return u.storage.Task.Create(ctx, task)
}

func (u *UseCase) Get(ctx context.Context, id string) (models.Task, error) {
	return u.storage.Task.Get(ctx, id)
}

func (u *UseCase) GetAll(ctx context.Context) ([]models.Task, error) {
	return u.storage.Task.GetAll(ctx)
}

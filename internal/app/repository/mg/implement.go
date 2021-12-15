package mg

import (
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"mongou/internal/app/models"
)

type ImplementTask struct {
	client *mongo.Client
}

func New(client *mongo.Client) *ImplementTask {
	return &ImplementTask{
		client: client,
	}
}

func (i *ImplementTask) Create(ctx context.Context, task models.Task) (string, error) {
	res, err := i.client.Database("db").Collection("tasks").InsertOne(ctx, task)
	if err != nil {
		return "", err
	}

	id, _ := res.InsertedID.(string)
	return id, err
}

func (i *ImplementTask) Get(ctx context.Context, id string) (models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Task{}, errors.Wrapf(
			err,
			"getting object id from hex: %s",
			id,
		)
	}
	res := i.client.Database("db").Collection("tasks").FindOne(ctx, bson.D{
		{"_id", objID},
	})

	err = res.Err()
	if err != nil {
		return models.Task{}, err
	}

	var t models.Task
	err = res.Decode(&t)
	if err != nil {
		return models.Task{}, err
	}

	return t, nil
}

func (i *ImplementTask) GetAll(ctx context.Context) ([]models.Task, error) {
	res, err := i.client.Database("db").Collection("tasks").Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	err = res.Err()
	if err != nil {
		return nil, err
	}

	t := make([]models.Task, 0)
	for res.Next(ctx) {
		var tt models.Task
		err = res.Decode(&tt)
		if err != nil {
			return nil, errors.Wrap(
				err,
				"decode from collection result",
			)
		}

		t = append(t, tt)
	}

	return t, nil
}

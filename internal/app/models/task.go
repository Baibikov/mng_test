package models

type Task struct {
	ID   string `bson:"_id"`
	Name string `bson:"name"`
}

package models

import (
	"github.com/globalsign/mgo/bson"
)

// Task is the structure for a task.

type Task struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	Completed   bool          `bson:"completed" json:"completed"`
}

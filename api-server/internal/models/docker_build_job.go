package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DockerBuildJobModel struct {
	Id     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Status string             `bson:"status" json:"status"`
}

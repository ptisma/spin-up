package repositories

import (
	"context"
	"log"

	"api-server/internal/infrastructures/db"
	"api-server/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DockerBuildJobRepository struct {
	DB db.DB
}

func (r *DockerBuildJobRepository) CreateDockerBuildJob(status string) (string, error) {
	var err error

	job := models.DockerBuildJobModel{Status: status}

	log.Println("test")

	collection := r.DB.GetClient().Database("test").Collection("jobs")

	insertResult, err := collection.InsertOne(context.TODO(), job)

	if err != nil {
		return "", err
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)

	res, _ := insertResult.InsertedID.(primitive.ObjectID)

	return res.String(), err

}

func (r *DockerBuildJobRepository) GetDockerBuildJobByID(id string) (models.DockerBuildJobModel, error) {

	var job models.DockerBuildJobModel

	collection := r.DB.GetClient().Database("test").Collection("jobs")

	filter := bson.M{"_id": id}

	err := collection.FindOne(context.TODO(), filter).Decode(&job)

	return job, err

}

package services

import (
	"api-server/internal/interfaces"
	"api-server/internal/models"
	"log"
)

type DockerBuildJobService struct {
	I1 interfaces.IDockerBuildJobRepository
}

func (s *DockerBuildJobService) CreateDockerBuildJob(status string) (string, error) {
	log.Println("test")
	return s.I1.CreateDockerBuildJob(status)
}

func (s *DockerBuildJobService) GetDockerBuildJobByID(id string) (models.DockerBuildJobModel, error) {
	return s.I1.GetDockerBuildJobByID(id)
}

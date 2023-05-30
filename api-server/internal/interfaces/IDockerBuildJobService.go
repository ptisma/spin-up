package interfaces

import (
	"api-server/internal/models"
)

type IDockerBuildJobService interface {
	CreateDockerBuildJob(status string) (string, error)
	GetDockerBuildJobByID(id string) (models.DockerBuildJobModel, error)
}

package interfaces

import (
	"api-server/internal/models"
)

type IDockerBuildJobRepository interface {
	CreateDockerBuildJob(status string) (string, error)
	GetDockerBuildJobByID(id string) (models.DockerBuildJobModel, error)
}

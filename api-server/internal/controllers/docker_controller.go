package controllers

import (
	"api-server/internal/interfaces"
	"api-server/internal/models/requestModels"
	"api-server/internal/models/viewModels"
	"encoding/json"
	"log"
	"net/http"
)

type DockerController struct {
	I1 interfaces.IDockerBuildJobService
}

func (c *DockerController) Test(w http.ResponseWriter, r *http.Request) {

	var dockerImgBuild requestModels.DockerImgBuildModel

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&dockerImgBuild)

	if err != nil {
		log.Panicln(err)
	}

	log.Println(dockerImgBuild)

	json.NewEncoder(w).Encode(viewModels.DockerImgBuildModel{JobID: "1234"})

}

func (c *DockerController) CreateDockerBuildJob(w http.ResponseWriter, r *http.Request) {

	log.Println("test")

	id, err := c.I1.CreateDockerBuildJob("pending")
	if err != nil {
		log.Panicln(err)
	}

	json.NewEncoder(w).Encode(viewModels.DockerImgBuildModel{JobID: id})

}

func (c *DockerController) GetDockerBuidJob(w http.ResponseWriter, r *http.Request) {

	var id string
	dockerBuildJob, err := c.I1.GetDockerBuildJobByID(id)
	if err != nil {
		log.Panicln(err)
	}

	json.NewEncoder(w).Encode(dockerBuildJob)

}

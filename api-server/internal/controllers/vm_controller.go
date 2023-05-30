package controllers

import (
	"api-server/internal/interfaces"
	"encoding/json"
	"log"
	"net/http"
)

type VmController struct {
	interfaces.IVMService
}

func (c *VmController) CreateContainerBuidJob(w http.ResponseWriter, r *http.Request) {

}

func (c *VmController) GetHealth(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Status OK"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}

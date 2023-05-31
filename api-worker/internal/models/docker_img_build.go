package models

type DockerImgBuildModel struct {
	GitURI         string
	GitSecretToken string
	ConfigPath     string
	PackerScript   string
}

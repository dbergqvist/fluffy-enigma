package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	SanaAPI struct {
		BaseURL     string `json:"base_url"`
		AccessToken string `json:"access_token"`
	} `json:"sana_api"`
	ScriveAPI struct {
		BaseURL     string `json:"base_url"`
		AccessToken string `json:"access_token"`
	} `json:"scrive_api"`
	TeacherEmail string `json:"teacher_email"`
}

var AppConfig Config

func LoadConfig(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&AppConfig)
}

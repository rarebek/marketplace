package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		AuthHost           string
		EventHost          string
		MedalHost          string
		AthleteHost        string
		AiService          string
		ServerAddress      string
		StreamHost         string
		CompanyServiceHost string
	}
)

const (
	OtpSecret = "some_secret"
	SignKey   = "nodirbek"
)

func (c *Config) Load() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	c.ServerAddress = os.Getenv("SERVER_ADDRESS")
	c.CompanyServiceHost = os.Getenv("COMPANY_SERVICE_HOST")
	return nil
}

func New() (*Config, error) {
	var cnfg Config
	if err := cnfg.Load(); err != nil {
		return nil, err
	}
	return &cnfg, nil
}

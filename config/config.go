package config

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

type GoogleDriveConfig struct {
	ApiKey string
}

type BackupConfig struct {
	Cron        string `validate:"required,cron"`
	Zip         bool
	GoogleDrive GoogleDriveConfig
}

type MongoDbConfig struct {
	Host     string
	Port     int
	Database string
}

type Config struct {
	Backup  BackupConfig
	Mongodb MongoDbConfig
}

func (c *Config) Load(path string) error {

	content, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, c)

	if err != nil {
		return err
	}

	return nil
}

func (c *Config) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(c)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

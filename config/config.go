package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type GoogleDriveConfig struct {
	ApiKey string
}

type BackupConfig struct {
	Cron        string
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

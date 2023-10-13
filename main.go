package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/laurence-trippen/go-mongodb-backup/archive"
	"github.com/laurence-trippen/go-mongodb-backup/config"
	"github.com/laurence-trippen/go-mongodb-backup/mongodump"
)

func main() {

	config := config.Config{}
	err := config.Load("data/single_config.yaml")
	if err != nil {
		log.Fatal("Couldn't load config!")
	}

	err = config.Validate()
	if err != nil {
		log.Fatal("Config isn't valid!")
	}

	path, err := mongodump.CheckExecutable()
	if err != nil {
		log.Fatal("Couldn't find executable 'mongodump'. Please install mongodump!")
	}

	fmt.Println("mongodump path: ", path)

	mongodump.PrintVersion()

	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Cron(config.Backup.Cron).Do(func() {
		folderName := archive.FolderNamePrefix + time.Now().Format(archive.FolderDateFormat)

		fmt.Println("start backing up ", config.Mongodb.Database, " to ", folderName+".zip")

		mongodump.Dump(config.Mongodb.Database, folderName)

		if config.Backup.Zip {

			err := archive.ZipArchiveFolder(folderName)
			if err != nil {
				fmt.Println(err)
			}

		}

		fmt.Println("backup done!")
	})

	scheduler.StartBlocking()

}

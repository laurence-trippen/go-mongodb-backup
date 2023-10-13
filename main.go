package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/laurence-trippen/go-mongodb-backup/config"
)

func main() {

	c := config.Config{}
	err := c.Load("data/single_config.yaml")
	if err != nil {
		log.Fatal("Couldn't load config!")
	}

	err = c.Validate()
	if err != nil {
		log.Fatal("Config isn't valid!")
	}

	fmt.Println(c)

	// Check for mongodump bin
	cmd := exec.Command("mongodump", "--version")

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		log.Fatal("Couldn't find executable 'mongodump'. Please install mongodump!")
	}

	fmt.Println("Found mongodump:")
	fmt.Println(out.String())

	// archive.CreateArchiveFolder()

	scheduler := gocron.NewScheduler(time.UTC)

	scheduler.Cron(c.Backup.Cron).Do(func() {
		fmt.Println("Do Backup")
	})

	scheduler.StartBlocking()

}

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
	err := c.Load("single_config.yaml")
	if err != nil {
		log.Fatal("Couldn't load config!")
	}

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

	cron := gocron.NewScheduler(time.UTC)

	cron.Every(5).Seconds().Do(func() {
		fmt.Println("Hello")
	})

	cron.StartBlocking()

}

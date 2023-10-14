package googledrive

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
)

const serviceAccountCredentials string = "private/google-credentials.json"

var srv *drive.Service

func Connect() {
	ctx := context.Background()

	var err error
	srv, err = drive.NewService(ctx, option.WithCredentialsFile(serviceAccountCredentials), option.WithScopes(drive.DriveScope))
	if err != nil {
		log.Fatalf("Warning: Unable to create drive Client %v", err)
	}
}

func UploadFile(path string) error {
	if srv != nil {
		return errors.New("drive service isn't initialized. Please call googledrive.Connect()")
	}

	file, err := os.Open(path)
	if err != nil {
		return errors.New("couldn't open file for upload: " + path)
	}
	defer file.Close()

	info, _ := file.Stat()

	driveFile := &drive.File{Name: info.Name()}

	res, err := srv.Files.
		Create(driveFile).
		Media(file).
		ProgressUpdater(func(now, size int64) { fmt.Printf("%d, %d\r", now, size) }).
		Do()

	if err != nil {
		return errors.New("failed uploading " + info.Name() + " to google drive!")
	}

	fmt.Printf("New file id: %s\n", res.Id)

	return nil
}

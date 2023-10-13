package archive

import (
	"os"
)

const FolderNamePrefix string = "backup-"
const FolderDateFormat string = "2006-01-02T15-04-05"

func CreateArchiveFolder() (string, error) {
	archiveName := FolderNamePrefix + FolderDateFormat

	err := os.Mkdir(archiveName, os.ModeAppend)

	if err != nil {
		return "", err
	}

	return archiveName, nil
}

func ZipArchiveFolder() {

}

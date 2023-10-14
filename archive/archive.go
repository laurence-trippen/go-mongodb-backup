package archive

import (
	"archive/zip"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
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

func ZipArchiveFolder(folderPath string) (string, error) {

	zipPath := folderPath + ".zip"

	archive, err := os.Create(zipPath)

	if err != nil {
		return zipPath, err
	}

	defer archive.Close()

	zipWriter := zip.NewWriter(archive)
	defer zipWriter.Close()

	err = filepath.Walk(folderPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		fmt.Println(path, info.Size())

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		writer, err := zipWriter.Create(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, file)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Println(err)
	}

	return zipPath, nil
}

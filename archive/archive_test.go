package archive_test

import (
	"os"
	"testing"

	"github.com/laurence-trippen/go-mongodb-backup/archive"
)

func TestCreateArchiveFolder(t *testing.T) {
	folderName, err := archive.CreateArchiveFolder()

	if err != nil {
		t.Fatal(err)
	}

	_, err = os.Stat(folderName)

	if err != nil {
		t.Fatal(err)
	}
}

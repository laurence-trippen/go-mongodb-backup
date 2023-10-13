package mongodump

import (
	"bytes"
	"fmt"
	"os/exec"
)

func PrintVersion() error {
	cmd := exec.Command("mongodump", "--version")

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println(out.String())

	return nil
}

func CheckExecutable() (string, error) {
	return exec.LookPath("mongodump")
}

func Dump(database string, path string) error {
	cmd := exec.Command("mongodump", "-d", database, "-o", path)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println(out.String())

	return nil
}

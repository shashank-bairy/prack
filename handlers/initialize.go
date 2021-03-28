package handlers

import (
	"io"
	"os"
	"path/filepath"

	"github.com/BA1RY/prack/storage"
	"github.com/BA1RY/prack/utils"
)

func generatePrackYAML() error {
	path := filepath.Join(utils.ProjectPath, "template", "prack.yaml")

	in, err := os.Open(path)
	if err != nil {
		return err
	}
	defer in.Close()

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	out, err := os.Create(filepath.Join(cwd, "prack.yaml"))
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func HandleInit() error {
	dbCon, err := GetDBCon()
	if err != nil {
		return err
	}

	err = storage.CreateProjectTable(dbCon.db)
	if err != nil {
		return err
	}

	err = generatePrackYAML()
	if err != nil {
		return err
	}

	return nil
}

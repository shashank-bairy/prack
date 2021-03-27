package handlers

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/BA1RY/prack/models"
	"github.com/BA1RY/prack/storage"
	"gopkg.in/yaml.v2"
)

type ProjectBlock struct {
	Project models.Project `yaml:"project,flow"`
}

func readPrackYaml(pb *ProjectBlock) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	yamlFile, err := ioutil.ReadFile(filepath.Join(cwd, "prack.yaml"))
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, pb)
	if err != nil {
		return err
	}

	return nil
}

func HandleAdd() error {
	pb := new(ProjectBlock)
	err := readPrackYaml(pb)
	if err != nil {
		return err
	}

	dbCon, err := GetDBCon()
	if err != nil {
		return err
	}

	err = storage.AddProject(dbCon.db, pb.Project)
	if err != nil {
		return err
	}

	// dev
	storage.DisplayTables(dbCon.db)

	return nil
}

func HandleRemove(alias string) error {
	// pb := new(ProjectBlock)
	// err := readPrackYaml(pb)
	// if err != nil {
	// 	return err
	// }

	dbCon, err := GetDBCon()
	if err != nil {
		return err
	}

	isPresent, err := storage.IsProjectPresent(dbCon.db, alias)
	if err != nil {
		return err
	}
	if !isPresent {
		return fmt.Errorf("deletion not possible as project is not present")
	}

	err = storage.RemoveProject(dbCon.db, alias)
	if err != nil {
		return err
	}

	// dev
	storage.DisplayTables(dbCon.db)

	return nil
}

func HandleUpdate() error {
	pb := new(ProjectBlock)
	err := readPrackYaml(pb)
	if err != nil {
		return err
	}

	dbCon, err := GetDBCon()
	if err != nil {
		return err
	}

	isPresent, err := storage.IsProjectPresent(dbCon.db, pb.Project.Alias)
	if err != nil {
		return err
	}

	if !isPresent {
		return fmt.Errorf("update not possible as project is not present")
	}

	err = storage.RemoveProject(dbCon.db, pb.Project.Alias)
	if err != nil {
		return err
	}

	err = storage.AddProject(dbCon.db, pb.Project)
	if err != nil {
		return err
	}

	// dev
	storage.DisplayTables(dbCon.db)

	return nil
}

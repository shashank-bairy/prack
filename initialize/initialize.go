package initialize

import (
	"io"
	"os"
	"path/filepath"

	"github.com/BA1RY/prack/utils"
)

func GeneratePrackYAML() error {

	path := filepath.Join(utils.ProjectPath, "./template/prack.yaml")

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

// type commandBlock struct {
// 	Alias    string
// 	Commands []string
// }

// type Project struct {
// 	Name          string
// 	Alias         string
// 	Description   string
// 	Tags          []string
// 	CommandBlocks []commandBlock
// }

// var project = Project{
// 	Name:        "Project Name",
// 	Alias:       "project_alias",
// 	Description: "Project description",
// 	Tags: []string{
// 		"tag1",
// 		"tag2",
// 	},
// 	CommandBlocks: []commandBlock{
// 		{
// 			Alias: "command_alias_1",
// 			Commands: []string{
// 				"command 1",
// 				"command 2",
// 				"command 3",
// 			},
// 		},
// 		{
// 			Alias: "command_alias_2",
// 			Commands: []string{
// 				"command 1",
// 				"command 2",
// 			},
// 		},
// 	},
// }

// func GeneratePrackYAML() error {
// 	paths := []string{
// 		filepath.Join(utils.ProjectPath, "./template/prack.tmpl"),
// 	}

// 	t := template.Must(template.New("prack.tmpl").ParseFiles(paths...))

// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		return err
// 	}

// 	f, err := os.Create(filepath.Join(cwd, "prack.yaml"))
// 	if err != nil {
// 		return err
// 	}

// 	err = t.Execute(f, project)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

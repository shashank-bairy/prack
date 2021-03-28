package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/BA1RY/prack/storage"
)

func executeCommands(commands []string) error {
	for _, command := range commands {
		args := strings.Split(command, " ")
		_, isPresent := os.LookupEnv(args[0])
		if isPresent {
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				return err
			}
		} else {
			switch args[0] {
			case "cd":
				os.Chdir(args[1])
			case "rm":
				if len(args[1:]) == 1 {
					os.Remove(args[1])
				} else {
					return fmt.Errorf("cannot delete file")
				}
			}
		}
	}

	return nil
}

func HandleOpen(args []string) error {

	if len(args) != 1 && len(args) != 2 {
		return fmt.Errorf("please provide the project and command block alias")
	}

	projectAlias := args[0]
	var cbAlias string

	dbCon, err := GetDBCon()
	if err != nil {
		return err
	}

	if len(args) == 1 {
		cbAliasList, err := storage.GetCommandBlocks(dbCon.db, projectAlias)
		if err != nil {
			return err
		}
		if len(cbAliasList) == 1 {
			cbAlias = cbAliasList[0]
		} else {
			return fmt.Errorf("please provide the command block name")
		}
	} else {
		cbAlias = args[1]
	}

	commands, err := storage.GetCommands(dbCon.db, projectAlias, cbAlias)
	if err != nil {
		return err
	}

	err = executeCommands(commands)
	if err != nil {
		return err
	}

	return nil
}

package handlers

import (
	"fmt"

	"github.com/BA1RY/prack/storage"
)

func HandleList() error {
	dbCon, err := GetDBCon()
	if err != nil {
		return err
	}

	projectMap, err := storage.GetProjects(dbCon.db)
	if err != nil {
		return err
	}

	if len(projectMap) == 0 {
		fmt.Println("No projects are present")
	} else {
		i := 0
		fmt.Printf("SL No\tProject Alias\tProject Name\n")
		for alias, name := range projectMap {
			fmt.Printf("%d\t%s\t%s\n", i+1, alias, name)
			i++
		}
	}

	return nil
}

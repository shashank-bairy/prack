package handlers

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BA1RY/prack/utils"
)

type DBCon struct {
	db *sql.DB
}

func GetDBCon() (*DBCon, error) {
	dbPath := filepath.Join(utils.ProjectPath, "project.db")
	if !utils.FileExists(dbPath) {
		_, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
	}

	db, err := sql.Open("sqlite3", fmt.Sprintf("file:%s?_foreign_keys=on", dbPath))
	if err != nil {
		return nil, err
	}

	dbCon := &DBCon{db: db}

	return dbCon, nil
}

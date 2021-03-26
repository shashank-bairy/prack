package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BA1RY/prack/models"
	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
)

func executeQuery(db *sql.DB, query string) error {
	statement, err := db.Prepare(query)
	if err != nil {
		return err
	}
	statement.Exec()
	defer statement.Close()
	return nil
}

func DisplayTables(db *sql.DB) error {
	rows, err := db.Query(`SELECT * FROM project`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		p := new(models.ProjectTable)
		rows.Scan(&p.UUID, &p.Name, &p.Alias, &p.Description)
		fmt.Println(p)
	}

	fmt.Println()

	rows, err = db.Query(`SELECT * FROM tag`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		t := new(models.TagTable)
		rows.Scan(&t.UUID, &t.Label, &t.ProjectID)
		fmt.Println(t)
	}

	fmt.Println()

	rows, err = db.Query(`SELECT * FROM command_block`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		cb := new(models.CommandBlockTable)
		rows.Scan(&cb.UUID, &cb.Alias, &cb.ProjectID)
		fmt.Println(cb)
	}

	fmt.Println()

	rows, err = db.Query(`SELECT * FROM command`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		cmd := new(models.CommandTable)
		rows.Scan(&cmd.UUID, &cmd.Cmd, &cmd.Position, &cmd.CBlockID)
		fmt.Println(cmd)
	}

	return nil
}

func CreateProjectTable(db *sql.DB) error {
	createProjectTableSQL := `CREATE TABLE IF NOT EXISTS project (
		"uuid" TEXT NOT NULL PRIMARY KEY,
		"name" TEXT NOT NULL,
		"alias" TEXT NOT NULL UNIQUE,
		"description" TEXT
	);`

	createTagTableSQL := `CREATE TABLE IF NOT EXISTS tag (
		"uuid" TEXT NOT NULL PRIMARY KEY,
		"label" TEXT NOT NULL,
		"project_id" TEXT NOT NULL,
		FOREIGN KEY (project_id) REFERENCES project (uuid) ON DELETE CASCADE,
		UNIQUE (label, project_id)
	);`

	createCommandBlockTableSQL := `CREATE TABLE IF NOT EXISTS command_block (
		"uuid" TEXT NOT NULL PRIMARY KEY,
		"alias" TEXT NOT NULL,
		"project_id" TEXT NOT NULL,
		FOREIGN KEY (project_id) REFERENCES project (uuid) ON DELETE CASCADE,
		UNIQUE (alias, project_id)
	);`

	createCommandTableSQL := `CREATE TABLE IF NOT EXISTS command (
		"uuid" TEXT NOT NULL PRIMARY KEY,
		"cmd" TEXT NOT NULL,
		"position" INTEGER NOT NULL,
		"cblock_id" TEXT NOT NULL,
		FOREIGN KEY (cblock_id) REFERENCES command_block (uuid) ON DELETE CASCADE
	);`

	log.Println("Creating project table...")
	err := executeQuery(db, createProjectTableSQL)
	if err != nil {
		return err
	}

	log.Println("Creating tag table...")
	err = executeQuery(db, createTagTableSQL)
	if err != nil {
		return err
	}

	log.Println("Creating commandBlock table...")
	err = executeQuery(db, createCommandBlockTableSQL)
	if err != nil {
		return err
	}

	log.Println("Creating command table...")
	err = executeQuery(db, createCommandTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func AddProject(db *sql.DB, p models.Project) error {
	puuid := uuid.NewV4().String()
	project := &models.ProjectTable{
		UUID:        puuid,
		Name:        p.Name,
		Alias:       p.Alias,
		Description: p.Description,
	}

	insertProjectSQL := fmt.Sprintf(`INSERT INTO project
		(uuid, name, alias, description)
		VALUES
		("%s", "%s", "%s", "%s")
	;`, project.UUID, project.Name, project.Alias, project.Description)

	log.Println("Adding project data to database...")
	err := executeQuery(db, insertProjectSQL)
	if err != nil {
		return err
	}

	insertTagsSQL := `INSERT INTO tag
		(uuid, label, project_id)
		VALUES
	`
	for _, tag := range p.Tags {
		tuuid := uuid.NewV4().String()
		insertTagsSQL += fmt.Sprintf(`("%s", "%s", "%s"),`, tuuid, tag, puuid)
	}
	insertTagsSQL = insertTagsSQL[:len(insertTagsSQL)-1]

	log.Println("Adding tag data to database...")
	err = executeQuery(db, insertTagsSQL)
	if err != nil {
		return err
	}

	insertCommandBlockSQL := `INSERT INTO command_block
		(uuid, alias, project_id)
		VALUES
	`
	insertCommandsSQL := `INSERT INTO command
		(uuid, cmd, position, cblock_id)
		VALUES
	`
	for _, cb := range p.CommandBlocks {
		cbuuid := uuid.NewV4().String()
		insertCommandBlockSQL += fmt.Sprintf(`("%s", "%s", "%s"),`, cbuuid, cb.Alias, puuid)
		for i, cmd := range cb.Commands {
			cuuid := uuid.NewV4().String()
			insertCommandsSQL += fmt.Sprintf(`("%s", "%s", %d, "%s"),`, cuuid, cmd, i+1, cbuuid)
		}
	}
	insertCommandBlockSQL = insertCommandBlockSQL[:len(insertCommandBlockSQL)-1]
	insertCommandsSQL = insertCommandsSQL[:len(insertCommandsSQL)-1]

	log.Println("Adding command block data to database...")
	err = executeQuery(db, insertCommandBlockSQL)
	if err != nil {
		return err
	}

	log.Println("Adding commands data to database...")
	err = executeQuery(db, insertCommandsSQL)
	if err != nil {
		return err
	}

	return nil
}

func IsProjectPresent(db *sql.DB, alias string) (bool, error) {
	err := db.QueryRow("SELECT 1 FROM project WHERE alias = ?", alias).Scan(&alias)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func RemoveProject(db *sql.DB, alias string) error {
	deleteProjectSQL := fmt.Sprintf(`DELETE FROM project WHERE alias="%s"`, alias)
	err := executeQuery(db, deleteProjectSQL)
	if err != nil {
		return err
	}
	return nil
}

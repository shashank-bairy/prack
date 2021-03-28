package storage

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BA1RY/prack/models"
	_ "github.com/mattn/go-sqlite3"
	uuid "github.com/satori/go.uuid"
)

func DisplayTables(db *sql.DB) error {
	rows, err := db.Query(`SELECT * FROM project`)
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		p := new(models.ProjectTable)
		err = rows.Scan(&p.UUID, &p.Name, &p.Alias, &p.Description)
		if err != nil {
			return err
		}
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
		err = rows.Scan(&t.UUID, &t.Label, &t.ProjectID)
		if err != nil {
			return err
		}
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
		err = rows.Scan(&cb.UUID, &cb.Alias, &cb.ProjectID)
		if err != nil {
			return err
		}
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
		err = rows.Scan(&cmd.UUID, &cmd.Cmd, &cmd.Position, &cmd.CBlockID)
		if err != nil {
			return err
		}
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
	_, err := db.Exec(createProjectTableSQL)
	if err != nil {
		return err
	}

	log.Println("Creating tag table...")
	_, err = db.Exec(createTagTableSQL)
	if err != nil {
		return err
	}

	log.Println("Creating commandBlock table...")
	_, err = db.Exec(createCommandBlockTableSQL)
	if err != nil {
		return err
	}

	log.Println("Creating command table...")
	_, err = db.Exec(createCommandTableSQL)
	if err != nil {
		return err
	}

	return nil
}

func AddProject(db *sql.DB, project models.Project) error {
	puuid, err := uuid.NewV4()
	if err != nil {
		return err
	}

	insertProjectSQL := `INSERT INTO project
		(uuid, name, alias, description)
		VALUES
		(?, ?, ?, ?)
	;`

	log.Println("Adding project data to database...")
	_, err = db.Exec(insertProjectSQL, puuid, project.Name, project.Alias, project.Description)
	if err != nil {
		return err
	}

	insertTagsSQL := `INSERT INTO tag
		(uuid, label, project_id)
		VALUES
		(?, ?, ?)
	`
	insertTagsStmt, err := db.Prepare(insertTagsSQL)
	if err != nil {
		return err
	}

	log.Println("Adding tag data to database...")
	for _, tag := range project.Tags {
		tuuid, err := uuid.NewV4()
		if err != nil {
			return err
		}
		_, err = insertTagsStmt.Exec(tuuid.String(), tag, puuid)
		if err != nil {
			return err
		}
	}
	defer insertTagsStmt.Close()

	insertCommandBlockSQL := `INSERT INTO command_block
		(uuid, alias, project_id)
		VALUES
		(?, ?, ?)
	`
	insertCommandBlockStmt, err := db.Prepare(insertCommandBlockSQL)
	if err != nil {
		return err
	}

	log.Println("Adding command block data to database...")

	insertCommandSQL := `INSERT INTO command
		(uuid, cmd, position, cblock_id)
		VALUES
		(?, ?, ?, ?)
	`
	insertCommandStmt, err := db.Prepare(insertCommandSQL)
	if err != nil {
		return err
	}

	log.Println("Adding commands data to database...")

	for _, cb := range project.CommandBlocks {
		cbuuid, err := uuid.NewV4()
		if err != nil {
			return err
		}
		insertCommandBlockStmt.Exec(cbuuid.String(), cb.Alias, puuid)
		for i, cmd := range cb.Commands {
			cuuid, err := uuid.NewV4()
			if err != nil {
				return err
			}
			insertCommandStmt.Exec(cuuid.String(), cmd, i+1, cbuuid)
		}
	}

	defer insertCommandBlockStmt.Close()
	defer insertCommandStmt.Close()

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
	deleteProjectSQL := `DELETE FROM project WHERE alias=?`
	_, err := db.Exec(deleteProjectSQL, alias)
	if err != nil {
		return err
	}
	return nil
}

func GetProjects(db *sql.DB) (map[string]string, error) {
	rows, err := db.Query(`SELECT * FROM project`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projectMap := make(map[string]string)
	for rows.Next() {
		p := models.ProjectTable{}
		err = rows.Scan(&p.UUID, &p.Name, &p.Alias, &p.Description)
		if err != nil {
			return nil, err
		}

		projectMap[p.Alias] = p.Name
	}

	return projectMap, nil
}

func GetCommandBlocks(db *sql.DB, projectAlias string) ([]string, error) {
	getCommandBlocksSQL := `SELECT command_block.alias
		FROM ((SELECT * FROM project where alias=?) project
		INNER JOIN command_block ON project.uuid = command_block.project_id)
	;`

	rows, err := db.Query(getCommandBlocksSQL, projectAlias)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var commandBlocks []string
	for rows.Next() {
		var cb string
		err = rows.Scan(&cb)
		if err != nil {
			return nil, err
		}

		commandBlocks = append(commandBlocks, cb)
	}
	return commandBlocks, err
}

func GetCommands(db *sql.DB, projectAlias string, cbAlias string) ([]string, error) {
	getCommandsSQL := `SELECT command.cmd
		FROM (((SELECT * FROM project where alias= ?) project
		INNER JOIN (SELECT * FROM command_block WHERE alias= ?) command_block ON project.uuid = command_block.project_id)
		INNER JOIN command ON command_block.uuid = command.cblock_id)
		ORDER BY command.position ASC
	;`

	rows, err := db.Query(getCommandsSQL, projectAlias, cbAlias)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var commands []string
	for rows.Next() {
		var cmd string
		err = rows.Scan(&cmd)
		if err != nil {
			return nil, err
		}

		commands = append(commands, cmd)
	}
	return commands, err
}

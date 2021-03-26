package models

type CommandTable struct {
	UUID     string `json:"uuid"`
	Cmd      string `json:"cmd"`
	Position int    `json:"position"`
	CBlockID string `json:"cblock_id"`
}

type CommandBlockTable struct {
	UUID      string `json:"uuid"`
	Alias     string `json:"alias"`
	ProjectID string `json:"project_id"`
}

type TagTable struct {
	UUID      string `json:"uuid"`
	Label     string `json:"label"`
	ProjectID string `json:"project_id"`
}

type ProjectTable struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	Description string `json:"description"`
}

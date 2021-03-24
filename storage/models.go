package models

type Command struct {
	UUID     string `json:"uuid"`
	CBlockID string `json:"cblock_id"`
	Cmd      string `json:"cmd"`
	Order    int    `json:"order"`
}

type CommandBlock struct {
	UUID      string `json:"uuid"`
	Alias     string `json:"alias"`
	ProjectID string `json:"project_id"`
}

type Tag struct {
	UUID      string `json:"uuid"`
	Label     string `json:"label"`
	ProjectID string `json:"project_id"`
}

type Project struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Alias       string `json:"alias"`
	Description string `json:"description"`
}

package models

type CommandBlock struct {
	Alias    string   `yaml:"alias"`
	Commands []string `yaml:"commands"`
}

type Project struct {
	Name          string         `yaml:"name"`
	Alias         string         `yaml:"alias"`
	Description   string         `yaml:"description"`
	Tags          []string       `yaml:"tags"`
	CommandBlocks []CommandBlock `yaml:"commands,flow"`
}

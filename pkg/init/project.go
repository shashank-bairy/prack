package init

type commandBlock struct {
	Alias    string
	Commands []string
}

type Project struct {
	Name          string
	Alias         string
	Description   string
	Tags          []string
	CommandBlocks []commandBlock
}

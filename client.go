package dnd

type Client struct {
	Rules    *rulesService
	Spells   *spellsService
	Monsters *monstersService
}

// NewClient creates a new instance
// of the DnD REST API client
func NewClient() *Client {
	return &Client{
		Rules:    &rulesService{},
		Spells:   &spellsService{},
		Monsters: &monstersService{},
	}
}

package dnd

type Client struct {
	Rules    *RulesService
	Spells   *SpellsService
	Monsters *MonstersService
}

// NewClient creates a new instance
// of the DnD REST API client
func NewClient() *Client {
	return &Client{
		Rules:    &RulesService{},
		Spells:   &SpellsService{},
		Monsters: &MonstersService{},
	}
}

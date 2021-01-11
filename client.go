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
		Rules:    NewRulesService(),
		Spells:   NewSpellsService(),
		Monsters: NewMonstersService(),
	}
}

// NewCustomClient creates a new instance
// of the DnD REST API client
func NewCustomClient(url string) *Client {
	r := NewCustomRulesService(url)
	s := NewCustomSpellsService(url, nil)
	m := NewCustomMonstersService(url, nil)
	return &Client{
		Rules:    r,
		Spells:   s,
		Monsters: m,
	}
}

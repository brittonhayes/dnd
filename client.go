package dnd

type Client struct {
	*RulesService
	*SpellsService
	*MonstersService
}

// NewClient creates a new instance
// of the DnD REST API client
func NewClient() *Client {
	return &Client{
		RulesService:    NewRulesService(),
		SpellsService:   NewSpellsService(),
		MonstersService: NewMonstersService(),
	}
}

// NewCustomClient creates a new instance
// of the DnD REST API client
func NewCustomClient(url string, sp *SpellParams, mp *MonstersParams) *Client {
	return &Client{
		RulesService:    NewCustomRulesService(url),
		SpellsService:   NewCustomSpellsService(url, sp),
		MonstersService: NewCustomMonstersService(url, mp),
	}
}

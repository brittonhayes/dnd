package dnd

type Client struct {
	Rules     *RulesService
	Spells    *SpellsService
	Monsters  *MonstersService
	Equipment *EquipmentService
}

// NewClient creates a new instance
// of the DnD REST API client
func NewClient() *Client {
	return &Client{
		Rules:     NewRulesService(),
		Spells:    NewSpellsService(),
		Monsters:  NewMonstersService(),
		Equipment: NewEquipmentService(),
	}
}

// NewCustomClient creates a new instance
// of the DnD REST API client
func NewCustomClient(url string, sp *SpellsParams, mp *MonstersParams) *Client {
	return &Client{
		Rules:     NewCustomRulesService(url),
		Spells:    NewCustomSpellsService(url, sp),
		Monsters:  NewCustomMonstersService(url, mp),
		Equipment: NewEquipmentService(),
	}
}

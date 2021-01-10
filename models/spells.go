package models

type Spells struct {
	Index         string   `json:"index"`
	Name          string   `json:"name"`
	Desc          []string `json:"desc"`
	HigherLevel   []string `json:"higher_level,omitempty"`
	Range         string   `json:"range"`
	Components    []string `json:"components"`
	Material      string   `json:"material"`
	Ritual        bool     `json:"ritual"`
	Duration      string   `json:"duration"`
	Concentration bool     `json:"concentration"`
	CastingTime   string   `json:"casting_time"`
	Level         int      `json:"level"`
	AttackType    string   `json:"attack_type"`
	Damage        struct {
		DamageType struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"damage_type"`
		DamageAtSlotLevel map[string]string `json:"damage_at_slot_level"`
	} `json:"damage"`
	School struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"school"`
	Classes []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"classes"`
	Subclasses []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"subclasses"`
	URL string `json:"url"`
}

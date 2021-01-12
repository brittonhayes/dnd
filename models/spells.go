package models

// Spells are the spells of D&D
type Spells struct {
	// Index of the spell for shorthand searching
	Index string `json:"index"`

	// Name is the name of the spell
	Name string `json:"name"`

	// Desc is a description of the spell
	Desc []string `json:"desc"`

	// HigherLevel is a description for casting
	// the spell at a higher level
	HigherLevel []string `json:"higher_level,omitempty"`

	// Range of the spell
	Range string `json:"range"`

	// Components are the required components for a spell, shorthanded
	// to V,S, and M which stand for verbal, somatic, and material
	Components []string `json:"components"`

	// Material component for the spell to be cast
	Material string `json:"material"`

	// Ritual determines if a spell can be cast in a 10-min(in-game) ritual
	Ritual bool `json:"ritual"`

	// Duration determines how long the spell effect lasts
	Duration string `json:"duration"`

	// Concentration determines if a spell needs concentration to persist
	Concentration bool `json:"concentration"`

	// CastingTime is How long it takes for the spell to activate
	CastingTime string `json:"casting_time"`

	// Level of the spell
	Level int `json:"level"`

	// AttackType is the attack type of the spell
	AttackType string `json:"attack_type"`

	// Damage is the damage type and the damage dice
	// for the spell at each spell slot level
	Damage struct {
		DamageType        DamageType        `json:"damage_type"`
		DamageAtSlotLevel map[string]string `json:"damage_at_slot_level"`
	} `json:"damage"`

	// School is the magic school this spell
	// belongs to
	School struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"school"`

	// Classes that are able to learn this spell
	Classes []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"classes"`

	// Subclasses that have access to this spell.
	Subclasses []struct {
		Index string `json:"index"`
		Name  string `json:"name"`
		URL   string `json:"url"`
	} `json:"subclasses"`

	// URL is the URL reference of this resource
	URL string `json:"url"`
}

// GetName returns the name of the spell
func (s *Spells) GetName() string {
	return s.Name
}

// GetURL returns the URL of the spell
func (s *Spells) GetURL() string {
	return s.URL
}

// GetIndex returns the index of the spell
func (s *Spells) GetIndex() string {
	return s.Index
}

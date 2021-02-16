package models

import "encoding/json"

var (
	_ Sheriff = &Conditions{}
	_ Sheriff = &DamageType{}
	_ Sheriff = &MagicSchool{}
)

// Conditions alter a creature’s capabilities in a variety of ways and can arise as a
// result of a spell, a class feature, a monster’s attack, or other effect.
// Most conditions, such as blinded, are impairments, but a few, such as invisible,
// can be advantageous.
type Conditions struct {
	// Index is the condition index for shorthand searching.
	Index string `json:"index"`

	// Name is the name for this condition resource.
	Name string `json:"name"`

	// Desc is a list of brief descriptions
	// of the condition.
	Desc []string `json:"desc"`

	// URL of the referenced resource.
	URL string `json:"url"`
}

func (c *Conditions) JSON(data []byte) error {
	return json.Unmarshal(data, &c)
}

// DamageType defines the unique identifiers
// for this kind of damage
type DamageType struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

func (d *DamageType) JSON(data []byte) error {
	return json.Unmarshal(data, &d)
}

// MagicSchool defines the unique identifiers
// for this school of magic
type MagicSchool struct {
	// Index is the magic school
	// index for shorthand searching.
	Index string `json:"index"`

	// Name is name for this magic
	// school resource.
	Name string `json:"name"`

	// URL of the referenced resource
	URL string `json:"url"`
}

func (m *MagicSchool) JSON(data []byte) error {
	return json.Unmarshal(data, &m)
}

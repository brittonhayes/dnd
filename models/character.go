package models

import "encoding/json"

var (
	_ Sheriff = &Skill{}
	_ Sheriff = &AbilityScore{}
	_ Sheriff = &Proficiency{}
	_ Sheriff = &Language{}
)

// Skill represents a skill that a character possesses
type Skill struct {
	// Index is the skill index for shorthand searching.
	Index string `json:"index"`

	// Name is the abbreviated name for this skill.
	Name string `json:"name"`

	// Desc is a brief description of this skill
	// and its uses.
	Desc []string `json:"desc"`

	// AbilityScore is the ability score associated
	// with this skill.
	AbilityScore AbilityScore `json:"ability_score"`

	// URL is the URL of the referenced resource
	URL string `json:"url"`
}

func (s *Skill) JSON(data []byte) error {
	return json.Unmarshal(data, &s)
}

// AbilityScore is a number that defines the magnitude of each creature’s abilities.
// An ability score is not just a measure of innate capabilities, but also encompasses a creature’s training
// and competence in activities related to that ability.
type AbilityScore struct {
	// Index is the ability score index for
	// shorthand searching.
	Index string `json:"index"`

	// Name is the abbreviated name for
	// this ability score.
	Name string `json:"name"`

	// FullName for this ability score.
	FullName string `json:"full_name"`

	// Desc is A brief description of this
	// ability score and its uses.
	Desc []string `json:"desc"`

	// Skills is a list of skills that
	// use this ability score.
	Skills []APIReference `json:"skills"`

	// The URL of the referenced resource
	URL string `json:"url"`
}

func (a *AbilityScore) JSON(data []byte) error {
	return json.Unmarshal(data, &a)
}

// Proficiency By virtue of your race, your character can speak, read, and write certain Proficiencies.
type Proficiency struct {
	// Index is the proficiency index for shorthand searching.
	Index string `json:"index"`

	// Type is the general category of the proficiency.
	Type string `json:"type"`

	// Name is the name of this proficiency resource
	Name string `json:"name"`

	// Classes that start with this proficiency.
	Classes []APIReference `json:"classes"`

	// Races that start with this proficiency.
	Races []APIReference `json:"races"`

	// URL of the referenced resource.
	URL string `json:"url"`

	// References is a list of references
	References []APIReference `json:"references"`
}

func (p *Proficiency) JSON(data []byte) error {
	return json.Unmarshal(data, &p)
}

// Language By virtue of your race, your character can speak, read, and write certain languages.
type Language struct {
	// Index is the language index for shorthand searching.
	Index string `json:"index"`

	// Name is the name of this language resource.
	Name string `json:"name"`

	// Type is whether the language is standard or exotic.
	Type string `json:"type"`

	// TypicalSpeakers are races that tend to speak this language.
	TypicalSpeakers []string `json:"typical_speakers"`

	// Script is the script used for writing in this language.
	Script string `json:"script"`

	// URL of the referenced resource.
	URL string `json:"url"`
}

func (l *Language) JSON(data []byte) error {
	return json.Unmarshal(data, &l)
}

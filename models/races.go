package models

// Race Each race grants your character ability and skill bonuses as well as racial traits.
type Race struct {

	// Index is the race index for shorthand searching.
	Index string `json:"index"`

	// Name is the name for this race resource.
	Name string `json:"name"`

	// Speed is the base move speed for this race (in feet per round).
	Speed int `json:"speed"`

	// AbilityBonuses are racial bonuses to ability scores.
	AbilityBonuses []AbilityBonus `json:"ability_bonuses"`

	// Alignment is a flavor description of likely alignments this race takes.
	Alignment string `json:"alignment"`

	// Age is the flavor description of possible ages for this race.
	Age string `json:"age"`

	// Size is the size class of this race.
	Size string `json:"size"`

	// SizeDescription is the flavor description
	// of height and weight for this race.
	SizeDescription string `json:"size_description"`

	// StartingProficiencies for all new characters of this race.
	StartingProficiencies []APIReference `json:"starting_proficiencies"`

	// StartingProficiencyOptions are the starting
	// proficiencies for all new characters of this race
	StartingProficiencyOptions []Choice `json:"starting_proficiency_options"`

	// Languages are starting languages for all new characters of this race.
	Languages []APIReference `json:"languages"`

	// LanguageDesc is a flavor description of the languages this race knows.
	LanguageDesc string `json:"language_desc"`

	// Traits are racial traits that provide benefits to its members.
	Traits []APIReference `json:"traits"`

	// SubRaces is list of related subraces
	SubRaces []APIReference `json:"subraces"`

	// The URL of the referenced resource
	URL string `json:"url"`
}

// SubRace contains the details for a parent race's associated
// subrace
type SubRace struct {
	// Index is the subrace index for shorthand searching.
	Index string `json:"index"`

	// Name is the name for this subrace resource.
	Name string `json:"name"`

	// Race is the parent race for this subrace.
	Race string `json:"race"`

	// Desc is a flavor description of this subrace
	Desc string `json:"name"`

	// AbilityBonuses are ability bonuses granted by this sub race.
	AbilityBonuses []AbilityBonus `json:"ability_bonuses"`

	// StartingProficiencies for all new characters of this subrace.
	StartingProficiencies []APIReference `json:"starting_proficiencies"`

	// Languages are starting languages for all new characters of this subrace.
	Languages []APIReference `json:"languages"`

	// Traits are racial traits that provide benefits to its members.
	Traits []APIReference `json:"traits"`

	// The URL of the referenced resource
	URL string `json:"url"`
}

// Traits for races in D&D
type Traits struct {
	// Index is the trait index for
	// shorthand searching.
	Index string `json:"index"`

	// Races is the list of races that have access
	// to this trait
	Races []APIReference `json:"races"`

	// SubRaces is list of related subraces
	SubRaces []APIReference `json:"subraces"`

	// Name is the name of the trait
	Name string `json:"name"`

	// Desc is the description of
	// the trait
	Desc string `json:"desc"`

	// Proficiencies are Proficiencies
	// this trait grants.
	Proficiencies []APIReference `json:"proficiencies"`

	// ProficiencyChoices are choices of proficiencies
	// this trait grants.
	ProficiencyChoices []Choice `json:"proficiency_choices"`
}

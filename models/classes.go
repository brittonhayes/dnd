package models

import "encoding/json"

var (
	_ Sheriff = &Class{}
	_ Sheriff = &Subclass{}
	_ Sheriff = &Features{}
	_ Sheriff = &StartingEquipment{}
	_ Sheriff = &StartingEquipmentItem{}
)

// Class is a fundamental part of the identity and nature of characters in the Dungeons & Dragons role-playing game.
// A character's capabilities, strengths, and weaknesses are largely defined by its class.
// A character's class affects a character's available skills and abilities.
type Class struct {
	// Index is the class index for shorthand searching.
	Index string `json:"index"`

	// Name is the name for this class resource.
	Name string `json:"name"`

	// HitDie is the hit die of the class. (ex: 12 == 1d12).
	HitDie int `json:"hit_die"`

	// ProficiencyChoices are starting proficiencies where
	// the player must choose a certain number from the given list of proficiencies.
	ProficiencyChoices []Choice `json:"proficiency_choices"`

	// Proficiencies are starting proficiencies all new characters of this class start with.
	Proficiencies []APIReference `json:"proficiencies"`

	// SavingThrows that the class is proficient in.
	SavingThrows []APIReference `json:"saving_throws"`

	// StartingEquipment is an object with the possible
	// choices of equipment for new characters of this class.
	StartingEquipment string `json:"starting_equipment"`

	// ClassLevels are all possible levels that this
	// class can obtain (excluding subclass-specific features)
	ClassLevels string `json:"class_levels"`

	// Spells is the URL of the class's spell Resource List.
	// Returns a list of all spells that can be learned or cast by the class.
	Spells []string `json:"spells"`

	// The URL reference of this resource
	URL string `json:"url"`
}

func (c *Class) JSON(data []byte) error {
	return json.Unmarshal(data, &c)
}

// Subclass reflects the different paths a class may take as levels are gained
type Subclass struct {
	// Index is the class index for shorthand searching.
	Index string `json:"index"`

	// Class is the parent class for this subclass
	Class string `json:"class"`

	// Name is the name for this class resource.
	Name string `json:"name"`

	// SubclassFlavor is the lore-friendly
	// flavor text for a classes respective subclass.
	SubclassFlavor string `json:"subclass_flavor"`

	// Desc is the description of the subclass resource.
	Desc string `json:"description"`

	// SubClassLevels is a resource url that shows the subclass level progression
	SubClassLevels APIReference `json:"sub_class_levels"`

	// Spells is the URL of the class's spell Resource List.
	// Returns a list of all spells that can be learned or cast by the class.
	Spells []string `json:"spells"`

	// The URL reference of this resource
	URL string `json:"url"`
}

func (c *Subclass) JSON(data []byte) error {
	return json.Unmarshal(data, &c)
}

// Features are the features of a class
type Features struct {
	// Index is the feature index for shorthand searching.
	Index string `json:"index"`

	// Name is the name for this feature resource.
	Name string `json:"name"`

	// Level is the level this feature
	// is gained
	Level int `json:"level"`

	// Class is the class that gains this feature.
	Class APIReference `json:"class"`

	// SubClass is the subclass that gains feature resource.
	SubClass APIReference `json:"sub_class"`

	// Desc is the description of the subclass resource.
	Desc string `json:"description"`

	// The URL reference of this resource
	URL string `json:"url"`
}

func (f *Features) JSON(data []byte) error {
	return json.Unmarshal(data, &f)
}

// StartingEquipment are the starting equipment of a class
type StartingEquipment struct {
	// Index is the class index for shorthand searching
	Index string `json:"index"`

	// Class is the class that gains this equipment.
	Class APIReference `json:"class"`

	// StartingEquipment is the starting equipment a character automatically gets.
	StartingEquipment []StartingEquipmentItem `json:"starting_equipment"`

	// StartingEquipmentOptions is the Starting equipment where the player must
	// choose a certain number from the given list of equipment.
	StartingEquipmentOptions []Choice `json:"starting_equipment_options"`

	// The URL reference of this resource
	URL string `json:"url"`
}

func (s *StartingEquipment) JSON(data []byte) error {
	return json.Unmarshal(data, &s)
}

// StartingEquipmentItem is a single item in
// a list of starting equipment
type StartingEquipmentItem struct {
	Equipment APIReference `json:"equipment"`
	Quantity  int          `json:"quantity"`
}

func (c *StartingEquipmentItem) JSON(data []byte) error {
	return json.Unmarshal(data, &c)
}

package models

type Resource struct {
	Count   int            `json:"count"`
	Results []APIReference `json:"results"`
}

type APIReference struct {
	// Index is the resource index for shorthand searching.
	Index string `json:"index"`

	// Name is the name of the referenced resource.
	Name string `json:"name"`

	// URL of the referenced resource.
	URL string `json:"url"`
}

type ClassAPIResource struct {
	// Index is the resource index for shorthand searching.
	Index string `json:"index"`

	// Class of whom the resource belongs to.
	Class string `json:"class"`

	// URL of the referenced resource.
	URL string `json:"url"`
}

type Choice struct {
	// Choose is the number of items to pick from the list.
	Choose int `json:"choose"`

	// Type of the resources to choose from.
	Type string `json:"type"`

	// From is a list of resources to choose from.
	From []APIReference `json:"from"`
}

type Cost struct {
	// Quantity is the numerical amount of coins.
	Quantity int `json:"quantity"`

	// Unit of coinage
	Unit string `json:"unit"`
}

type AbilityBonus struct {
	// Bonus amount for this ability score.
	Bonus int `json:"bonus"`

	// AbilityScore for this bonus.
	AbilityScore []APIReference `json:"ability_score"`
}

type DamageType struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

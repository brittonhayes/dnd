package models

// Results is used for any API endpoint without
// a resource index or name. It return a list of
// available resources for that API
type Resource struct {
	// Count is the total
	// number of results returned
	Count int `json:"count"`

	// Results is the list of API references
	// return by the request
	Results []APIReference `json:"results"`
}

// APIReference is the basic structure that defines
// where this resource resides and what it is called
type APIReference struct {
	// Index is the resource index for shorthand searching.
	Index string `json:"index"`

	// Name is the name of the referenced resource.
	Name string `json:"name"`

	// URL of the referenced resource.
	URL string `json:"url"`
}

// ClassAPIResource is very similar to APIReference but
// returns a class name rather than just a name
type ClassAPIResource struct {
	// Index is the resource index for shorthand searching.
	Index string `json:"index"`

	// Class of whom the resource belongs to.
	Class string `json:"class"`

	// URL of the referenced resource.
	URL string `json:"url"`
}

// Choice provides a set of options related to a resource
type Choice struct {
	// Choose is the number of items to pick from the list.
	Choose int `json:"choose"`

	// Type of the resources to choose from.
	Type string `json:"type"`

	// From is a list of resources to choose from.
	From []APIReference `json:"from"`
}

// Cost is the basic structure of defining
// an item's financial worth in D&D
type Cost struct {
	// Quantity is the numerical amount of coins.
	Quantity int `json:"quantity"`

	// Unit of coinage
	Unit string `json:"unit"`
}

// AbilityBonus is the bonus modifier for
// an ability and its associated references
type AbilityBonus struct {
	// Bonus amount for this ability score.
	Bonus int `json:"bonus"`

	// AbilityScore for this bonus.
	AbilityScore []APIReference `json:"ability_score"`
}


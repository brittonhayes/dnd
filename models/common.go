package models

type Resource struct {
	Count   int            `json:"count"`
	Results []APIReference `json:"results"`
}

type APIReference struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type ClassAPIResource struct {
	Index string `json:"index"`
	Class string `json:"class"`
	URL   string `json:"url"`
}

type Choice struct {
	Choose int            `json:"choose"`
	Type   string         `json:"type"`
	From   []APIReference `json:"from"`
}

type Cost struct {
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

type AbilityBonus struct {
	Bonus        int            `json:"bonus"`
	AbilityScore []APIReference `json:"ability_score"`
}

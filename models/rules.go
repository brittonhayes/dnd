package models

type Rules struct {
	Name        string            `json:"name"`
	Index       string            `json:"index"`
	Desc        string            `json:"desc"`
	Subsections []RulesSubsection `json:"subsections"`
	URL         string            `json:"url"`
}

type RulesSubsection struct {
	Name  string `json:"name"`
	Index string `json:"index"`
	URL   string `json:"url"`
}

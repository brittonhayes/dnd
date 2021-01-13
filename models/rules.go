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
	Desc  string `json:"desc,omitempty"`
}

// GetName returns the name of the rule
func (r *Rules) GetName() string {
	return r.Name
}

// GetName returns the index of the rule
func (r *Rules) GetIndex() string {
	return r.Index
}

// GetURL returns the URL of the rule
func (r *Rules) GetURL() string {
	return r.URL
}

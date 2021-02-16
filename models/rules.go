package models

import "encoding/json"

var (
	_ Sheriff = &Rules{}
	_ Sheriff = &RulesSubsection{}
)

type Rules struct {
	Name        string            `json:"name"`
	Index       string            `json:"index"`
	Desc        string            `json:"desc"`
	Subsections []RulesSubsection `json:"subsections"`
	URL         string            `json:"url"`
}

func (r *Rules) JSON(data []byte) error {
	return json.Unmarshal(data, &r)
}

type RulesSubsection struct {
	Name  string `json:"name"`
	Index string `json:"index"`
	URL   string `json:"url"`
	Desc  string `json:"desc,omitempty"`
}

func (r *RulesSubsection) JSON(data []byte) error {
	return json.Unmarshal(data, &r)
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

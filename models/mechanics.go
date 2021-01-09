package models

type Conditions struct {
	Index string   `json:"index"`
	Name  string   `json:"name"`
	Desc  []string `json:"desc"`
	URL   string   `json:"url"`
}

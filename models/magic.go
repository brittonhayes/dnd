package models

type MagicItems struct {
	Index             string            `json:"index"`
	Name              string            `json:"name"`
	EquipmentCategory EquipmentCategory `json:"equipment_category"`
	Desc              []string          `json:"desc"`
	URL               string            `json:"url"`
}

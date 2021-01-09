package models

type Equipment struct {
	Index             string            `json:"index"`
	Name              string            `json:"name"`
	EquipmentCategory EquipmentCategory `json:"equipment_category"`
	GearCategory      GearCategory      `json:"gear_category"`
	Cost              Cost              `json:"cost"`
	Weight            int               `json:"weight"`
	URL               string            `json:"url"`
}

type EquipmentCategory struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type GearCategory struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

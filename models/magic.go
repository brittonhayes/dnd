package models

// MagicItem is the structure for a magic item obtained
// in the dungeons you explore.
type MagicItem struct {
	// Index is the damage type index for shorthand searching.
	Index string `json:"index"`

	// Name is the name for this equipment resource
	Name string `json:"name"`

	// EquipmentCategory is the category of equipment this falls into
	EquipmentCategory APIReference `json:"equipment_category"`

	// Desc is a list of descriptors for this item
	Desc []string `json:"desc"`

	// URL is the URL reference of this resource
	URL string `json:"url"`
}

// GetName gets the name of the magic item
func (m *MagicItem) GetName() string {
	return m.Name
}

// GetIndex gets the index of the magic item
func (m *MagicItem) GetIndex() string {
	return m.Index
}

// GetURL gets the url of the magic item
func (m *MagicItem) GetURL() string {
	return m.URL
}

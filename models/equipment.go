package models

import "encoding/json"

var (
	_ Getter  = &AdventuringGear{}
	_ Getter  = &Armor{}
	_ Getter  = &EquipmentPack{}
	_ Getter  = &Weapon{}
	_ Sheriff = &AdventuringGear{}
	_ Sheriff = &ArmorClass{}
	_ Sheriff = &Armor{}
	_ Sheriff = &EquipmentPack{}
	_ Sheriff = &PackContents{}
	_ Sheriff = &WeaponProperties{}
	_ Sheriff = &Weapon{}
)

// Weapon is the structure for weapons obtained
// in the dungeons you explore.
type Weapon struct {
	// Index is the damage type index for shorthand searching.
	Index string `json:"index"`

	// Name is the name for this equipment resource
	Name string `json:"name"`

	// EquipmentCategory is the category of equipment this falls into
	EquipmentCategory APIReference `json:"equipment_category"`

	// WeaponCategory is the category of weapon this falls into
	WeaponCategory string `json:"weapon_category"`

	// WeaponRange is whether this is a Melee or Ranged weapon.
	WeaponRange string `json:"weapon_range"`

	// CategoryRange is a combination of weapon_category and weapon_range.
	CategoryRange string `json:"category_range"`

	// Cost is the financial worth of this equipment
	Cost Cost `json:"cost"`

	// Damage is the data on dice, bonus, and damage type
	Damage ActionDamage `json:"damage"`

	// TwoHandedDamage when wielded with two hands.
	// It includes data on dice, bonus, and damage type.
	TwoHandedDamage *ActionDamage `json:"two_handed_damage,omitempty"`

	// Range Includes the normal and long range for a weapon,
	// only pertains to Ranged weapons
	Range *Range `json:"range"`

	// Weight is the numerical weight of this item
	Weight int `json:"weight"`

	// Properties is a list of properties this weapon has
	Properties []APIReference `json:"properties"`

	// URL is the URL reference of this resource
	URL string `json:"url"`
}

func (w *Weapon) JSON(data []byte) error {
	return json.Unmarshal(data, &w)
}

type Range struct {
	Normal *int `json:"normal"`
	Long   *int `json:"long"`
}

// GetName returns the name of the weapon
func (w *Weapon) GetName() string {
	return w.Name
}

// GetName returns the index of the weapon
func (w *Weapon) GetIndex() string {
	return w.Index
}

// GetURL returns the URL of the weapon
func (w *Weapon) GetURL() string {
	return w.URL
}

// Armor is the structure for armor obtained
// in the dungeons you explore.
type Armor struct {
	// Index is the damage type index for shorthand searching.
	Index string `json:"index"`

	// Name is the name for this equipment resource
	Name string `json:"name"`

	// EquipmentCategory is the category of equipment this falls into
	EquipmentCategory APIReference `json:"equipment_category"`

	// ArmorCategory is the category of armor this falls into
	ArmorCategory string `json:"armor_category"`

	// ArmorCategory is the details on how to calculate armor class.
	ArmorClass ArmorClass `json:"armor_class"`

	// StrengthMinumum is Minimum STR required to use this armor.
	StrengthMinimum int `json:"str_minimum"`

	// StealthDisadvantage is Whether the armor gives disadvantage for Stealth.
	StealthDisadvantage bool `json:"stealth_disadvantage"`

	// Weight is the numerical weight of this item
	Weight int `json:"weight"`

	// Cost is the financial worth of this equipment
	Cost Cost `json:"cost"`

	// URL is the URL reference of this resource
	URL string `json:"url"`
}

func (a *Armor) JSON(data []byte) error {
	return json.Unmarshal(data, &a)
}

// GetName returns the name of the gear
func (a *Armor) GetName() string {
	return a.Name
}

// GetName returns the index of the gear
func (a *Armor) GetIndex() string {
	return a.Index
}

// GetURL returns the URL of the gear
func (a *Armor) GetURL() string {
	return a.URL
}

// ArmorClass is the details on how to calculate armor class.
type ArmorClass struct {
	Base     *int  `json:"base"`
	DexBonus *bool `json:"dex_bonus"`
	MaxBonus *bool `json:"max_bonus"`
}

func (a *ArmorClass) JSON(data []byte) error {
	return json.Unmarshal(data, &a)
}

// AdventuringGear is the structure for adventuring gear obtained
// in the dungeons you explore.
type AdventuringGear struct {
	// Index is the damage type index for shorthand searching.
	Index string `json:"index"`

	// Name is the name for this equipment resource
	Name string `json:"name"`

	// EquipmentCategory is the category of equipment this falls into
	EquipmentCategory APIReference `json:"equipment_category"`

	// GearCategory is the category of gear this falls into
	GearCategory APIReference `json:"gear_category"`

	// Cost is the financial worth of this equipment
	Cost Cost `json:"cost"`

	// Weight is the numerical weight of this item
	Weight int `json:"weight"`

	// URL is the URL reference of this resource
	URL string `json:"url"`
}

func (a *AdventuringGear) JSON(data []byte) error {
	return json.Unmarshal(data, &a)
}

// GetName returns the name of the gear
func (a *AdventuringGear) GetName() string {
	return a.Name
}

// GetName returns the index of the gear
func (a *AdventuringGear) GetIndex() string {
	return a.Index
}

// GetURL returns the URL of the gear
func (a *AdventuringGear) GetURL() string {
	return a.URL
}

// EquipmentPack is the structure for an equipment pack obtained
// in the dungeons you explore.
type EquipmentPack struct {
	// Index is the damage type index for shorthand searching.
	Index string `json:"index"`

	// Name is the name for this equipment resource
	Name string `json:"name"`

	// EquipmentCategory is the category of equipment this falls into
	EquipmentCategory APIReference `json:"equipment_category"`

	// GearCategory is the category of gear this falls into
	GearCategory APIReference `json:"gear_category"`

	// Contents is the list of items inside the pack
	Contents []PackContents `json:"contents"`

	// Cost is the financial worth of this equipment
	Cost Cost `json:"cost"`

	// URL is the URL reference of this resource
	URL string `json:"url"`
}

func (e *EquipmentPack) JSON(data []byte) error {
	return json.Unmarshal(data, &e)
}

type PackContents struct {
	// Item is the of item index details
	Item APIReference `json:"item"`

	// Quantity is the numerical amount of
	// this item present
	Quantity int `json:"quantity"`
}

func (p *PackContents) JSON(data []byte) error {
	return json.Unmarshal(data, &p)
}

// GetName returns the name of the equipment
func (e *EquipmentPack) GetName() string {
	return e.Name
}

// GetName returns the index of the equipment
func (e *EquipmentPack) GetIndex() string {
	return e.Index
}

// GetURL returns the URL of the equipment
func (e *EquipmentPack) GetURL() string {
	return e.URL
}

// WeaponProperties is the structure for a weapon's properties
type WeaponProperties struct {
	// Index is the damage type index for shorthand searching.
	Index string `json:"index"`

	// Name is the name for this equipment resource
	Name string `json:"name"`

	// Desc is a list of descriptors for this item
	Desc []string `json:"desc"`

	// URL is the URL reference of this resource
	URL string `json:"url"`
}

func (w *WeaponProperties) JSON(data []byte) error {
	return json.Unmarshal(data, &w)
}

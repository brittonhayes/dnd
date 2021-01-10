package models

type ActionDamage struct {
	DamageDice string        `json:"damage_dice"`
	DamageType *APIReference `json:"damage_type"`
}

type Action struct {
	AttackBonus int             `json:"attack_bonus"`
	Damage      []*ActionDamage `json:"damage"`
	Options     *Options        `json:"options,omitempty"`
	Desc        string          `json:"desc"`
	Name        string          `json:"name"`
}

type Options struct {
	Choose int             `json:"choose,omitempty"`
	From   [][]interface{} `json:"from,omitempty"`
}

type LegendaryAction struct {
	AttackBonus int           `json:"attack_bonus,omitempty"`
	Desc        string        `json:"desc"`
	Name        string        `json:"name"`
	Damage      []interface{} `json:"damage"`
}

type Proficiency struct {
	Proficiency *APIReference `json:"proficiency"`
	Value       int           `json:"value,omitempty"`
}

type Reaction struct {
	Desc string `json:"desc"`
	Name string `json:"name"`
}

type Sense struct {
	Blindsight        string `json:"blindsight,omitempty"`
	Darkvision        string `json:"darkvision,omitempty"`
	PassivePerception int    `json:"passive_perception,omitempty"`
	Tremorsense       string `json:"termorsense,omitempty"`
	Truesight         string `json:"truesight,omitempty"`
}

type SpecialAbility struct {
	Desc string `json:"desc"`
	Name string `json:"name"`
	DC   *DC    `json:"dc,omitempty"`
}

type DC struct {
	DcValue     int           `json:"dc_value"`
	SuccessType string        `json:"success_type"`
	DcType      *APIReference `json:"dc_type"`
}

type Speed struct {
	Burrow string `json:"burrow,omitempty"`
	Climb  string `json:"climb,omitempty"`
	Fly    string `json:"fly,omitempty"`
	Hover  bool   `json:"hover,omitempty"`
	Swim   string `json:"swim,omitempty"`
	Walk   string `json:"walk,omitempty"`
}

type Monster struct {
	Actions               []*Action          `json:"actions"`
	Alignment             string             `json:"alignment"`
	ArmorClass            int                `json:"armor_class"`
	ChallengeRating       int                `json:"challenge_rating"`
	Charisma              int                `json:"charisma"`
	ConditionImmunities   []*APIReference    `json:"condition_immunities"`
	Constitution          int                `json:"constitution"`
	DamageImmunities      []interface{}      `json:"damage_immunities"`
	DamageResistances     []interface{}      `json:"damage_resistances"`
	DamageVulnerabilities []interface{}      `json:"damage_vulnerabilities"`
	Dexterity             int                `json:"dexterity"`
	Forms                 []*APIReference    `json:"forms,omitempty"`
	HitDice               string             `json:"hit_dice"`
	HitPoints             int                `json:"hit_points"`
	Index                 string             `json:"index"`
	Intelligence          int                `json:"intelligence"`
	Languages             string             `json:"languages"`
	LegendaryActions      []*LegendaryAction `json:"legendary_actions"`
	Name                  string             `json:"name"`
	Proficiencies         []*Proficiency     `json:"proficiencies"`
	Reactions             []*Reaction        `json:"reactions,omitempty"`
	Senses                *Sense             `json:"senses"`
	Size                  string             `json:"size"`
	SpecialAbilities      []*SpecialAbility  `json:"special_abilities"`
	Speed                 *Speed             `json:"speed"`
	Strength              int                `json:"strength"`
	Subtype               interface{}        `json:"subtype"`
	Type                  string             `json:"type"`
	URL                   string             `json:"url"`
	Wisdom                int                `json:"wisdom"`
	XP                    int                `json:"xp"`
}

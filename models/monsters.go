package models

import "encoding/json"

var _ Sheriff = &Monster{}

type Monster struct {
	// Index is the monster index for shorthand searching.
	Index string `json:"index"`

	// Name is the name of the monster.
	Name string `json:"name"`

	// Size is the size of the monster ranging from Tiny to Gargantuan.
	Size string `json:"size"`

	// Type is the type of monster.
	Type string `json:"type"`

	// SubType is the subtype of monster.
	SubType interface{} `json:"subtype"`

	// Alignment is а creature's general moral and personal attitudes.
	Alignment string `json:"alignment"`

	// ArmorClass is difficulty for a player to successfully deal damage to a monster.
	ArmorClass int `json:"armor_class"`

	// HitPoints is hit points of a monster determine how much
	// damage it is able to take before it can be defeated
	HitPoints int `json:"hit_points"`

	// HitDice of a monster can be used to make a version of the same monster whose hit points are determined by the roll of the die.
	// For example: A monster with 2d6 would have its hit points determine by rolling a 6 sided die twice.
	HitDice string `json:"hit_dice"`

	// Forms applicable to Lycanthropes that have multiple forms. This links to the other related monster entries that are the same.
	Forms []*APIReference `json:"forms,omitempty"`

	// Speed for a monster determines how fast it can move per turn.
	Speed Speed `json:"speed"`

	// Strength of a monster. How hard a monster can hit a player.
	Strength int `json:"strength"`

	// Dexterity of a monster. The monster's ability for swift movement or stealth
	Dexterity int `json:"dexterity"`

	// Constitution of a monster. How sturdy a monster is.
	Constitution int `json:"constitution"`

	// Intelligence of a monster. The monster's ability to outsmart a player
	Intelligence int `json:"intelligence"`

	// Wisdom of a monster. A monster's ability to ascertain the player's plan.
	Wisdom int `json:"wisdom"`

	// Charisma of a monster. A monster's ability to charm or intimidate a player
	Charisma int `json:"charisma"`

	// Proficiencies is a list of proficiencies of a monster.
	Proficiencies []*MonsterProficiency `json:"proficiencies"`

	// DamageVulnerabilities is a list of damage types that a monster will take double damage from.
	DamageVulnerabilities []interface{} `json:"damage_vulnerabilities"`

	// DamageResistances is a list of damage types that a monster will take half damage from.
	DamageResistances []interface{} `json:"damage_resistances"`

	// DamageImmunities is a list of damage types that a monster will take zero damage from.
	DamageImmunities []interface{} `json:"damage_immunities"`

	// ConditionImmunities is a list of conditions that a monster is immune to.
	ConditionImmunities []*APIReference `json:"condition_immunities"`

	// Senses Monsters typically have a passive perception but they might also have other senses to detect players
	Senses *Sense `json:"senses"`

	// Languages are languages a monster is able to speak
	Languages string `json:"languages"`

	// ChallengeRating is monster's challenge rating is a guideline number that says when a monster
	// becomes an appropriate challenge against the party's average level.
	ChallengeRating int `json:"challenge_rating"`

	// XP is the experience points from defeating this monster
	XP int `json:"xp"`

	// SpecialAbilities is a list of the monster's special abilities.
	SpecialAbilities []*SpecialAbility `json:"special_abilities"`

	// Actions is a list of actions that is available to the monster to take during combat.
	Actions []interface{} `json:"actions"`

	// LegendaryActions is a list of actions that can be used outside of the monster's initiative.
	LegendaryActions []interface{} `json:"legendary_actions,omitempty"`

	// URL of the referenced resource
	URL string `json:"url"`
}

func (m *Monster) JSON(data []byte) error {
	return json.Unmarshal(data, &m)
}

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

type MonsterProficiency struct {
	Proficiency *APIReference `json:"proficiency"`
	Value       int           `json:"value,omitempty"`
}

type Reaction struct {
	Desc string `json:"desc"`
	Name string `json:"name"`
}

type Sense struct {
	BlindSight        string `json:"blindsight,omitempty"`
	DarkVision        string `json:"darkvision,omitempty"`
	PassivePerception int    `json:"passive_perception,omitempty"`
	TremorSense       string `json:"tremorsense,omitempty"`
	TrueSight         string `json:"truesight,omitempty"`
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

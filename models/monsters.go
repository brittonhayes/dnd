package models

type Monster struct {
	Index      string      `json:"index"`
	Name       string      `json:"name"`
	Size       string      `json:"size"`
	Type       string      `json:"type"`
	Subtype    interface{} `json:"subtype"`
	Alignment  string      `json:"alignment"`
	ArmorClass int         `json:"armor_class"`
	HitPoints  int         `json:"hit_points"`
	HitDice    string      `json:"hit_dice"`
	Speed      struct {
		Walk string `json:"walk"`
		Swim string `json:"swim"`
	} `json:"speed"`
	Strength      int `json:"strength"`
	Dexterity     int `json:"dexterity"`
	Constitution  int `json:"constitution"`
	Intelligence  int `json:"intelligence"`
	Wisdom        int `json:"wisdom"`
	Charisma      int `json:"charisma"`
	Proficiencies []struct {
		Value       int `json:"value"`
		Proficiency struct {
			Index string `json:"index"`
			Name  string `json:"name"`
			URL   string `json:"url"`
		} `json:"proficiency"`
	} `json:"proficiencies"`
	DamageVulnerabilities []interface{} `json:"damage_vulnerabilities"`
	DamageResistances     []interface{} `json:"damage_resistances"`
	DamageImmunities      []interface{} `json:"damage_immunities"`
	ConditionImmunities   []interface{} `json:"condition_immunities"`
	Senses                struct {
		Darkvision        string `json:"darkvision"`
		PassivePerception int    `json:"passive_perception"`
	} `json:"senses"`
	Languages        string `json:"languages"`
	ChallengeRating  int    `json:"challenge_rating"`
	Xp               int    `json:"xp"`
	SpecialAbilities []struct {
		Name string `json:"name"`
		Desc string `json:"desc"`
		Dc   struct {
			DcType struct {
				Index string `json:"index"`
				Name  string `json:"name"`
				URL   string `json:"url"`
			} `json:"dc_type"`
			DcValue     int    `json:"dc_value"`
			SuccessType string `json:"success_type"`
		} `json:"dc,omitempty"`
	} `json:"special_abilities"`
	Actions []struct {
		Name    string `json:"name"`
		Desc    string `json:"desc"`
		Options struct {
			Choose int `json:"choose"`
			From   [][]struct {
				Name  string `json:"name"`
				Count int    `json:"count"`
				Type  string `json:"type"`
			} `json:"from"`
		} `json:"options,omitempty"`
		Damage      []interface{} `json:"damage"`
		AttackBonus int           `json:"attack_bonus,omitempty"`
		Dc          struct {
			DcType struct {
				Index string `json:"index"`
				Name  string `json:"name"`
				URL   string `json:"url"`
			} `json:"dc_type"`
			DcValue     int    `json:"dc_value"`
			SuccessType string `json:"success_type"`
		} `json:"dc,omitempty"`
		Usage struct {
			Type  string `json:"type"`
			Times int    `json:"times"`
		} `json:"usage,omitempty"`
	} `json:"actions"`
	LegendaryActions []struct {
		Name        string `json:"name"`
		Desc        string `json:"desc"`
		AttackBonus int    `json:"attack_bonus,omitempty"`
		Damage      []struct {
			DamageType struct {
				Index string `json:"index"`
				Name  string `json:"name"`
				URL   string `json:"url"`
			} `json:"damage_type"`
			DamageDice string `json:"damage_dice"`
		} `json:"damage,omitempty"`
	} `json:"legendary_actions"`
	URL string `json:"url"`
}

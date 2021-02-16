package endpoints

type Endpoint string

const (
	BaseURL                Endpoint = "https://www.dnd5eapi.co/api"
	AbilityScoresURL       Endpoint = "/ability-scores"
	ClassesURL             Endpoint = "/classes"
	ConditionsURL          Endpoint = "/conditions"
	DamageTypesURL         Endpoint = "/damage-types"
	EquipmentCategoriesURL Endpoint = "/equipment-categories"
	EquipmentURL           Endpoint = "/equipment"
	FeaturesURL            Endpoint = "/features"
	LanguagesURL           Endpoint = "/languages"
	MagicItemsURL          Endpoint = "/magic-items"
	MonstersURL            Endpoint = "/monsters"
	MagicSchoolsURL        Endpoint = "/magic-schools"
	ProficienciesURL       Endpoint = "/proficiencies"
	RacesURL               Endpoint = "/races"
	RulesURL               Endpoint = "/rules"
	RuleSectionsURL        Endpoint = "/rule-sections"
	SkillsURL              Endpoint = "/skills"
	SpellcastingURL        Endpoint = "/spellcasting"
	SpellsURL              Endpoint = "/spells"
	StartingEquipmentURL   Endpoint = "/starting-equipment"
	SubClassesURL          Endpoint = "/subclasses"
	SubracesURL            Endpoint = "/subraces"
	TraitsURL              Endpoint = "/traits"
	WeaponPropertiesURL    Endpoint = "/weapon-properties"
)

func (e Endpoint) String() string {
	return string(e)
}

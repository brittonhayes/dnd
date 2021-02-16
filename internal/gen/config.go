package gen

import (
	"github.com/brittonhayes/dnd/endpoints"
	"github.com/brittonhayes/dnd/models"
)

type ServiceDefinition struct {
	Name    string
	Comment string
	Methods map[string]APIRequest
	Options Options
}

type Options map[string]string

type APIRequest struct {
	PathParams  map[string]string
	QueryParams map[string]string
	Result      interface{}
	Endpoint    endpoints.Endpoint
	Argument    string
	Comment     string
}

type EndpointDef struct {
	Name     string
	URL      string
	Response string `json:"response"`
}

var Endpoints = []EndpointDef{
	{Name: "AbilityScores", URL: endpoints.BaseURL.String() + endpoints.AbilityScoresURL.String()},
	{Name: "Classes", URL: endpoints.BaseURL.String() + endpoints.ClassesURL.String()},
	{Name: "Conditions", URL: endpoints.BaseURL.String() + endpoints.ConditionsURL.String()},
	{Name: "DamageTypes", URL: endpoints.BaseURL.String() + endpoints.DamageTypesURL.String()},
	{Name: "EquipmentWeapon", URL: endpoints.BaseURL.String() + endpoints.EquipmentURL.String() + "/club"},
	{Name: "EquipmentArmor", URL: endpoints.BaseURL.String() + endpoints.EquipmentURL.String() + "/padded"},
	{Name: "EquipmentAdventuringGear", URL: endpoints.BaseURL.String() + endpoints.EquipmentURL.String() + "/abacus"},
	{Name: "EquipmentPackBurglars", URL: endpoints.BaseURL.String() + endpoints.EquipmentURL.String() + "/burglars-pack"},
	{Name: "MonstersFindAboleth", URL: endpoints.BaseURL.String() + endpoints.MonstersURL.String() + "/aboleth"},
	{Name: "MonstersList", URL: endpoints.BaseURL.String() + endpoints.MonstersURL.String()},
	{Name: "RacesFindDwarf", URL: endpoints.BaseURL.String() + endpoints.RacesURL.String() + "/dwarf"},
	{Name: "RacesFindHuman", URL: endpoints.BaseURL.String() + endpoints.RacesURL.String() + "/human"},
	{Name: "RacesListRaces", URL: endpoints.BaseURL.String() + endpoints.RacesURL.String()},
	{Name: "RulesFindAdventuring", URL: endpoints.BaseURL.String() + endpoints.RulesURL.String() + "/adventuring"},
	{Name: "RulesList", URL: endpoints.BaseURL.String() + endpoints.RulesURL.String()},
	{Name: "RulesSectionAbilityChecks", URL: endpoints.BaseURL.String() + endpoints.RuleSectionsURL.String() + "/ability-checks"},
	{Name: "RulesSectionsList", URL: endpoints.BaseURL.String() + endpoints.RuleSectionsURL.String()},
	{Name: "SpellsList", URL: endpoints.BaseURL.String() + endpoints.SpellsURL.String()},
	{Name: "SpellsListFiltered", URL: endpoints.BaseURL.String() + endpoints.SpellsURL.String() + "?level=1&school=evo"},
	{Name: "SpellsFindAcidArrow", URL: endpoints.BaseURL.String() + endpoints.SpellsURL.String() + "/acid-arrow"},
	{Name: "SubracesFindHighElf", URL: endpoints.BaseURL.String() + endpoints.SubracesURL.String() + "/high-elf"},
	{Name: "SubracesList", URL: endpoints.BaseURL.String() + endpoints.SubracesURL.String()},
}

var Services = []ServiceDefinition{
	{
		Name:    "Monsters",
		Comment: "covers all endpoints related to Monsters",
		Options: map[string]string{
			"ChallengeRating": "challenge_rating",
		},
		Methods: map[string]APIRequest{
			"List": {
				Comment:  "lists the available monsters endpoints",
				Endpoint: endpoints.MonstersURL,
				Result:   models.Resource{},
			},
			"Find": {
				Comment:  "searches a specific monster by name",
				Endpoint: endpoints.MonstersURL,
				Result:   models.Monster{},
				Argument: "index",
			},
		},
	},
	{
		Name:    "Races",
		Comment: "covers all endpoints related to Races",
		Options: nil,
		Methods: map[string]APIRequest{
			"ListRaces": {
				Comment:  "lists out all races",
				Endpoint: endpoints.RacesURL,
				Result:   models.Resource{},
			},
			"ListSubRaces": {
				Comment:  "lists out all subraces",
				Endpoint: endpoints.SubracesURL,
				Result:   models.Resource{},
			},
			"FindRace": {
				Comment:  "find a race by its index",
				Endpoint: endpoints.RacesURL,
				Result:   models.Race{},
				Argument: "index",
			},
			"FindSubRace": {
				Comment:  "find a subrace by its index",
				Endpoint: endpoints.SubracesURL,
				Result:   models.SubRace{},
				Argument: "index",
			},
		},
	},
	{
		Name:    "Rules",
		Comment: "covers all endpoints related to Rules",
		Options: nil,
		Methods: map[string]APIRequest{
			"ListRules": {
				Comment:  "lists available rules endpoints",
				Endpoint: endpoints.RulesURL,
				Result:   models.Resource{},
			},
			"ListSections": {
				Comment:  "lists available rules sections endpoints",
				Endpoint: endpoints.RuleSectionsURL,
				Result:   models.Resource{},
			},
			"FindRule": {
				Comment:  "searches for specific rules based on their name",
				Endpoint: endpoints.RulesURL,
				Result:   models.Rules{},
				Argument: "index",
			},
			"FindSection": {
				Comment:  "searches for specific rules section based on their name",
				Endpoint: endpoints.RuleSectionsURL,
				Result:   models.RulesSubsection{},
				Argument: "index",
			},
		},
	},
	{
		Name:    "Spells",
		Comment: "covers all endpoints related to Spells",
		Options: map[string]string{
			"Name":   "name",
			"Level":  "level",
			"School": "school",
		},
		Methods: map[string]APIRequest{
			"List": {
				Comment:  "lists the available spells endpoints",
				Endpoint: endpoints.SpellsURL,
				Result:   models.Resource{},
			},
			"Find": {
				Comment:  "finds a spell by name",
				Endpoint: endpoints.SpellsURL,
				Result:   models.Spells{},
				Argument: "index",
			},
		},
	},
	{
		Name:    "Equipment",
		Comment: "covers all endpoints related to Equipment",
		Options: nil,
		Methods: map[string]APIRequest{
			"List": {
				Comment:  "lists the available equipment endpoints",
				Endpoint: endpoints.EquipmentURL,
				Result:   models.Resource{},
			},
			"FindWeapon": {
				Comment:  "finds a weapon's details by name",
				Endpoint: endpoints.EquipmentURL,
				Result:   models.Weapon{},
				Argument: "index",
			},
			"FindArmor": {
				Comment:  "finds an armor's details by name",
				Endpoint: endpoints.EquipmentURL,
				Result:   models.Armor{},
				Argument: "index",
			},
			"FindAdventuringGear": {
				Comment:  "finds a set of adventuring gear by name",
				Endpoint: endpoints.EquipmentURL,
				Result:   models.AdventuringGear{},
				Argument: "index",
			},
			"FindEquipmentPack": {
				Comment:  "finds an equipment pack by name",
				Endpoint: endpoints.EquipmentURL,
				Result:   models.EquipmentPack{},
				Argument: "index",
			},
		},
	},
}

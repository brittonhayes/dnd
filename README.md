<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# dnd

```go
import "github.com/brittonhayes/dnd"
```

package dnd is a Go Client for the DnD 5e REST API

### Installation

Install with the go get command

```
go get github.com/brittonhayes/dnd
```

### Documentation

View the full docs on pkg\.go\.dev https://pkg.go.dev/github.com/brittonhayes/dnd

### Usage

Using the package is as easy as create client\, pick the endpoint\, and run the method\. This applies across every data type so it is consistent across the board\. Here's a simple example of how to fetch a rule from the DnD 5e ruleset\.

```
func main() {
	// Create a dnd client
	c := dnd.NewClient()

	// Fetch DnD rules about adventuring
	r, err := c.Rules.FindRule("adventuring")
	if err != nil {
		panic(err)
	}

	// Print out the rule name
	fmt.Println("Name", r.Name)

	// Print out the rule description
	fmt.Println("Description", r.Desc)
}
```

Social image by Ashley Mcnamara https://twitter.com/ashleymcnamara 💖

### Development

If you'd like to contribute to DnD\, make sure you've got mage installed: https://magefile.org

```
# Download dependencies and run tests
mage download
mage test
```

## Index

- [Constants](<#constants>)
- [type Client](<#type-client>)
  - [func NewClient() *Client](<#func-newclient>)
  - [func NewCustomClient(url string, sp *SpellParams, mp *MonstersParams) *Client](<#func-newcustomclient>)
- [type Monsters](<#type-monsters>)
- [type MonstersParams](<#type-monstersparams>)
- [type MonstersService](<#type-monstersservice>)
  - [func NewCustomMonstersService(url string, params *MonstersParams) *MonstersService](<#func-newcustommonstersservice>)
  - [func NewMonstersService() *MonstersService](<#func-newmonstersservice>)
  - [func (s *MonstersService) FindMonster(name string) (*models.Monster, error)](<#func-monstersservice-findmonster>)
  - [func (s *MonstersService) ListMonsters() (*models.Resource, error)](<#func-monstersservice-listmonsters>)
- [type Races](<#type-races>)
- [type RacesService](<#type-racesservice>)
  - [func NewCustomRacesService(url string) *RacesService](<#func-newcustomracesservice>)
  - [func NewRacesService() *RacesService](<#func-newracesservice>)
  - [func (s *RacesService) FindRace(name string) (*models.Race, error)](<#func-racesservice-findrace>)
  - [func (s *RacesService) FindSubRace(name string) (*models.SubRace, error)](<#func-racesservice-findsubrace>)
  - [func (s *RacesService) ListRaces() (*models.Resource, error)](<#func-racesservice-listraces>)
  - [func (s *RacesService) ListSubRaces() (*models.Resource, error)](<#func-racesservice-listsubraces>)
- [type Rules](<#type-rules>)
- [type RulesService](<#type-rulesservice>)
  - [func NewCustomRulesService(url string) *RulesService](<#func-newcustomrulesservice>)
  - [func NewRulesService() *RulesService](<#func-newrulesservice>)
  - [func (s *RulesService) FindRule(name string) (*models.Rules, error)](<#func-rulesservice-findrule>)
  - [func (s *RulesService) FindSection(name string) (*models.RulesSubsection, error)](<#func-rulesservice-findsection>)
  - [func (s *RulesService) ListRules() (*models.Resource, error)](<#func-rulesservice-listrules>)
  - [func (s *RulesService) ListSections() (*models.Resource, error)](<#func-rulesservice-listsections>)
- [type SpellParams](<#type-spellparams>)
- [type Spells](<#type-spells>)
- [type SpellsService](<#type-spellsservice>)
  - [func NewCustomSpellsService(url string, params *SpellParams) *SpellsService](<#func-newcustomspellsservice>)
  - [func NewSpellsService() *SpellsService](<#func-newspellsservice>)
  - [func (s *SpellsService) FindSpell(name string) (*models.Spells, error)](<#func-spellsservice-findspell>)
  - [func (s *SpellsService) ListSpells() (*models.Resource, error)](<#func-spellsservice-listspells>)


## Constants

```go
const (
    BaseURL                = "https://www.dnd5eapi.co/api"
    AbilityScoresURL       = "/ability-scores"
    ClassesURL             = "/classes"
    ConditionsURL          = "/conditions"
    DamageTypesURL         = "/damage-types"
    EquipmentCategoriesURL = "/equipment-categories"
    EquipmentURL           = "/equipment"
    FeaturesURL            = "/features"
    LanguagesURL           = "/languages"
    MagicItemsURL          = "/magic-items"
    MonstersURL            = "/monsters"
    MagicSchoolsURL        = "/magic-schools"
    ProficienciesURL       = "/proficiencies"
    RacesURL               = "/races"
    RulesURL               = "/rules"
    RuleSectionsURL        = "/rule-sections"
    SkillsURL              = "/skills"
    SpellcastingURL        = "/spellcasting"
    SpellsURL              = "/spells"
    StartingEquipmentURL   = "/starting-equipment"
    SubClassesURL          = "/subclasses"
    SubracesURL            = "/subraces"
    TraitsURL              = "/traits"
    WeaponPropertiesURL    = "/weapon-properties"
)
```

## type Client

```go
type Client struct {
    Rules    *RulesService
    Spells   *SpellsService
    Monsters *MonstersService
}
```

### func NewClient

```go
func NewClient() *Client
```

NewClient creates a new instance of the DnD REST API client

### func NewCustomClient

```go
func NewCustomClient(url string, sp *SpellParams, mp *MonstersParams) *Client
```

NewCustomClient creates a new instance of the DnD REST API client

## type Monsters

The Monsters interface shows all of the available methods for the monsters endpoint

```go
type Monsters interface {
    ListMonsters() (*models.Resource, error)
    FindMonster(name string) (*models.Monster, error)
}
```

## type MonstersParams

```go
type MonstersParams struct {
    ChallengeRating string `url:"challenge_rating"`
}
```

## type MonstersService

```go
type MonstersService struct {
    // URL is the base URL of the service
    URL     string `default:"https://www.dnd5eapi.co/api"`
    Options *MonstersParams
}
```

### func NewCustomMonstersService

```go
func NewCustomMonstersService(url string, params *MonstersParams) *MonstersService
```

NewMonstersService creates a custom instance of the monsters service

<details><summary>Example</summary>
<p>

Create a new custom monsters service

```go
{
	s := NewCustomMonstersService(BaseURL, &MonstersParams{
		ChallengeRating: "5",
	})

	monsters, _ := s.ListMonsters()
	fmt.Println("Results: ", monsters.Results)
}
```

</p>
</details>

### func NewMonstersService

```go
func NewMonstersService() *MonstersService
```

NewMonstersService creates a new instance of the monsters service

<details><summary>Example</summary>
<p>

Create a new monsters service and apply custom query params

```go
{
	s := NewMonstersService()
	s.Options = &MonstersParams{
		ChallengeRating: "5",
	}
}
```

</p>
</details>

### func \(\*MonstersService\) FindMonster

```go
func (s *MonstersService) FindMonster(name string) (*models.Monster, error)
```

FindMonster fetches a monster's details by name

<details><summary>Example</summary>
<p>

Find a specific monster

```go
{
	c := NewClient()
	monster, _ := c.Monsters.FindMonster("aboleth")
	fmt.Printf("The monster %s has a challenge rating of %d", monster.Name, monster.ChallengeRating)
}
```

</p>
</details>

### func \(\*MonstersService\) ListMonsters

```go
func (s *MonstersService) ListMonsters() (*models.Resource, error)
```

ListMonsters available in the API

<details><summary>Example</summary>
<p>

Count the number of available monsters listed

```go
{
	s := NewMonstersService()
	s.Options = &MonstersParams{
		ChallengeRating: "3",
	}

	monsters, _ := s.ListMonsters()
	fmt.Printf("There are %d monsters available", monsters.Count)
}
```

</p>
</details>

## type Races

The Races interface shows all of the available methods for the monsters endpoint

```go
type Races interface {
    ListRaces() (*models.Resource, error)
    ListSubRaces() (*models.Resource, error)
    FindRace(name string) (*models.Race, error)
    FindSubRace(name string) (*models.SubRace, error)
}
```

## type RacesService

```go
type RacesService struct {
    // URL is the base URL of the service
    URL string `default:"https://www.dnd5eapi.co/api"`
}
```

<details><summary>Example (,ist Racers_ Count)</summary>
<p>

Count the number of available races listed

```go
{
	s := NewRacesService()

	races, _ := s.ListRaces()
	fmt.Printf("There are %d races available", races.Count)
}
```

</p>
</details>

### func NewCustomRacesService

```go
func NewCustomRacesService(url string) *RacesService
```

NewCustomRacesService creates a custom instance of the races service

<details><summary>Example</summary>
<p>

Create a new custom races service

```go
{
	s := NewCustomRacesService(BaseURL)

	races, _ := s.ListRaces()
	fmt.Println("Results: ", races.Results)
}
```

</p>
</details>

### func NewRacesService

```go
func NewRacesService() *RacesService
```

NewRacesService creates a new instance of the races service

### func \(\*RacesService\) FindRace

```go
func (s *RacesService) FindRace(name string) (*models.Race, error)
```

<details><summary>Example</summary>
<p>

Create a new races service and apply custom query params

```go
{
	s := NewRacesService()
	s.FindRace("dwarf")
}
```

</p>
</details>

### func \(\*RacesService\) FindSubRace

```go
func (s *RacesService) FindSubRace(name string) (*models.SubRace, error)
```

### func \(\*RacesService\) ListRaces

```go
func (s *RacesService) ListRaces() (*models.Resource, error)
```

### func \(\*RacesService\) ListSubRaces

```go
func (s *RacesService) ListSubRaces() (*models.Resource, error)
```

## type Rules

The Rules interface shows all of the available methods for the rules endpoint

```go
type Rules interface {
    ListRules() (*models.Resource, error)
    ListSections() (*models.Resource, error)
    FindRule(name string) (*models.Rules, error)
    FindSection(name string) (*models.RulesSubsection, error)
}
```

## type RulesService

```go
type RulesService struct {
    // URL is the base URL of the service
    URL string `default:"https://www.dnd5eapi.co/api"`
}
```

### func NewCustomRulesService

```go
func NewCustomRulesService(url string) *RulesService
```

NewCustomRulesService allows your to create a rules service with a custom base url

### func NewRulesService

```go
func NewRulesService() *RulesService
```

NewRulesService creates a default instance of the rules service

### func \(\*RulesService\) FindRule

```go
func (s *RulesService) FindRule(name string) (*models.Rules, error)
```

FindRule allows you to search for specific rules based on their name

<details><summary>Example</summary>
<p>

Basic example of printing a rule as JSON

```go
{

	c := NewClient()

	r, _ := c.Rules.FindRule("adventuring")

	j, _ := json.MarshalIndent(&r, "", "\t")
	fmt.Println(string(j))
}
```

</p>
</details>

### func \(\*RulesService\) FindSection

```go
func (s *RulesService) FindSection(name string) (*models.RulesSubsection, error)
```

FindSection allows you to search for a specific ruleset subsection based on its name

<details><summary>Example</summary>
<p>

Basic example of printing a rules section as JSON

```go
{

	c := NewClient()

	r, _ := c.Rules.FindSection("ability-checks")

	j, _ := json.MarshalIndent(&r, "", "\t")
	fmt.Println(string(j))
}
```

</p>
</details>

### func \(\*RulesService\) ListRules

```go
func (s *RulesService) ListRules() (*models.Resource, error)
```

ListRules lists the available DnD 5e rules in the API

### func \(\*RulesService\) ListSections

```go
func (s *RulesService) ListSections() (*models.Resource, error)
```

ListSections lists the available DnD 5e rule subsections in the API

## type SpellParams

```go
type SpellParams struct {
    Level  string `url:"level"`
    School string `url:"school"`
}
```

## type Spells

The Spells interface shows all of the available methods for the spells endpoint

```go
type Spells interface {
    ListSpells() (*models.Resource, error)
    FindSpell(name string) (*models.Spells, error)
}
```

## type SpellsService

```go
type SpellsService struct {
    // URL is the base URL of the service
    URL     string `default:"https://www.dnd5eapi.co/api"`
    Options *SpellParams
}
```

### func NewCustomSpellsService

```go
func NewCustomSpellsService(url string, params *SpellParams) *SpellsService
```

NewSpellsService creates a custom instance of the spells service

<details><summary>Example</summary>
<p>

Create a new custom spells service

```go
{
	s := NewCustomSpellsService(BaseURL, &SpellParams{
		Level:  "2",
		School: "",
	})

	spells, _ := s.ListSpells()
	fmt.Println("Results: ", spells.Results)
}
```

</p>
</details>

### func NewSpellsService

```go
func NewSpellsService() *SpellsService
```

NewSpellsService creates a default instance of the spells service

<details><summary>Example</summary>
<p>

Create a new spells service and apply custom query params

```go
{
	s := NewSpellsService()
	s.Options = &SpellParams{
		Level:  "5",
		School: "",
	}
}
```

</p>
</details>

### func \(\*SpellsService\) FindSpell

```go
func (s *SpellsService) FindSpell(name string) (*models.Spells, error)
```

FindSpell searches a specific spell by name

<details><summary>Example</summary>
<p>

Find a specific spell

```go
{
	c := NewClient()
	spell, _ := c.Spells.FindSpell("animate-objects")
	fmt.Printf("The spell %s has a range of %s", spell.Name, spell.Range)
}
```

</p>
</details>

### func \(\*SpellsService\) ListSpells

```go
func (s *SpellsService) ListSpells() (*models.Resource, error)
```

ListSpells lists the available spells endpoints

<details><summary>Example</summary>
<p>

Count the number of available spells listed

```go
{
	s := NewSpellsService()
	s.Options = &SpellParams{
		Level:  "5",
		School: "",
	}

	spells, _ := s.ListSpells()
	fmt.Printf("There are %d spells available", spells.Count)
}
```

</p>
</details>

## Stats

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/brittonhayes/dnd?color=blue&label=Latest%20Version&sort=semver)
[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/dnd)](https://goreportcard.com/report/github.com/brittonhayes/dnd)
![Test](https://github.com/brittonhayes/dnd/workflows/Test/badge.svg)
[![codecov](https://codecov.io/gh/brittonhayes/dnd/branch/main/graph/badge.svg?token=VN11FU4LBW)](https://codecov.io/gh/brittonhayes/dnd)


Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)

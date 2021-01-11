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

View the docs on pkg\.go\.dev https://pkg.go.dev/github.com/brittonhayes/dnd

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

## Index

- [Constants](<#constants>)
- [type Client](<#type-client>)
  - [func NewClient() *Client](<#func-newclient>)
  - [func NewCustomClient(url string) *Client](<#func-newcustomclient>)
- [type Monsters](<#type-monsters>)
- [type MonstersParams](<#type-monstersparams>)
- [type MonstersService](<#type-monstersservice>)
  - [func NewCustomMonstersService(url string, params *MonstersParams) *MonstersService](<#func-newcustommonstersservice>)
  - [func NewMonstersService() *MonstersService](<#func-newmonstersservice>)
  - [func (s *MonstersService) FindMonster(name string) (*models.Monster, error)](<#func-monstersservice-findmonster>)
  - [func (s *MonstersService) ListMonsters() (*models.APIReference, error)](<#func-monstersservice-listmonsters>)
- [type Rules](<#type-rules>)
- [type RulesService](<#type-rulesservice>)
  - [func NewCustomRulesService(url string) *RulesService](<#func-newcustomrulesservice>)
  - [func NewRulesService() *RulesService](<#func-newrulesservice>)
  - [func (r *RulesService) FindRule(name string) (*models.Rules, error)](<#func-rulesservice-findrule>)
  - [func (r *RulesService) FindSection(name string) (*models.RulesSubsection, error)](<#func-rulesservice-findsection>)
  - [func (r *RulesService) ListRules() (*models.APIReference, error)](<#func-rulesservice-listrules>)
  - [func (r *RulesService) ListSections() (*models.APIReference, error)](<#func-rulesservice-listsections>)
- [type SpellParams](<#type-spellparams>)
- [type Spells](<#type-spells>)
- [type SpellsService](<#type-spellsservice>)
  - [func NewCustomSpellsService(url string, params *SpellParams) *SpellsService](<#func-newcustomspellsservice>)
  - [func NewSpellsService() *SpellsService](<#func-newspellsservice>)
  - [func (s *SpellsService) FindSpell(name string) (*models.Spells, error)](<#func-spellsservice-findspell>)
  - [func (s *SpellsService) ListSpells() (*models.APIReference, error)](<#func-spellsservice-listspells>)


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
func NewCustomClient(url string) *Client
```

NewCustomClient creates a new instance of the DnD REST API client

## type Monsters

The Monsters interface shows all of the available methods for the monsters endpoint

```go
type Monsters interface {
    ListMonsters() (*models.APIReference, error)
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

### func NewMonstersService

```go
func NewMonstersService() *MonstersService
```

NewMonstersService creates a new instance of the monsters service

### func \(\*MonstersService\) FindMonster

```go
func (s *MonstersService) FindMonster(name string) (*models.Monster, error)
```

FindMonster fetches a monster's details by name

### func \(\*MonstersService\) ListMonsters

```go
func (s *MonstersService) ListMonsters() (*models.APIReference, error)
```

ListMonsters available in the API

## type Rules

The Rules interface shows all of the available methods for the rules endpoint

```go
type Rules interface {
    ListRules() (*models.APIReference, error)
    ListSections() (*models.APIReference, error)
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
func (r *RulesService) FindRule(name string) (*models.Rules, error)
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
func (r *RulesService) FindSection(name string) (*models.RulesSubsection, error)
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
func (r *RulesService) ListRules() (*models.APIReference, error)
```

ListRules lists the available DnD 5e rules in the API

### func \(\*RulesService\) ListSections

```go
func (r *RulesService) ListSections() (*models.APIReference, error)
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
    ListSpells() (*models.APIReference, error)
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

### func NewSpellsService

```go
func NewSpellsService() *SpellsService
```

NewSpellsService creates a default instance of the spells service

### func \(\*SpellsService\) FindSpell

```go
func (s *SpellsService) FindSpell(name string) (*models.Spells, error)
```

FindSpell searches a specific spell by name

### func \(\*SpellsService\) ListSpells

```go
func (s *SpellsService) ListSpells() (*models.APIReference, error)
```

ListSpells lists the available spells endpoints

## Stats

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/brittonhayes/dnd?color=blue&label=Latest%20Version&sort=semver)
[![Go Report Card](https://goreportcard.com/badge/github.com/brittonhayes/dnd)](https://goreportcard.com/report/github.com/brittonhayes/dnd)
![Test](https://github.com/brittonhayes/dnd/workflows/Test/badge.svg)
[![codecov](https://codecov.io/gh/brittonhayes/dnd/branch/main/graph/badge.svg?token=VN11FU4LBW)](https://codecov.io/gh/brittonhayes/dnd)


Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)

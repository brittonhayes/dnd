package gen

import (
	"fmt"
	"github.com/brittonhayes/dnd"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"
)

type endpoint struct {
	Name     string
	URL      string
	Response string `json:"response"`
}

type service struct {
	Resource string
	Model    string
}

func GenerateMocks() error {
	var tpl *template.Template
	m := []endpoint{
		{Name: "AbilityScores", URL: dnd.BaseURL + dnd.AbilityScoresURL},
		{Name: "Classes", URL: dnd.BaseURL + dnd.ClassesURL},
		{Name: "Conditions", URL: dnd.BaseURL + dnd.ConditionsURL},
		{Name: "DamageTypes", URL: dnd.BaseURL + dnd.DamageTypesURL},
		{Name: "EquipmentWeapon", URL: dnd.BaseURL + dnd.EquipmentURL + "/club"},
		{Name: "EquipmentArmor", URL: dnd.BaseURL + dnd.EquipmentURL + "/padded"},
		{Name: "EquipmentAdventuringGear", URL: dnd.BaseURL + dnd.EquipmentURL + "/abacus"},
		{Name: "EquipmentPackBurglars", URL: dnd.BaseURL + dnd.EquipmentURL + "/burglars-pack"},
		{Name: "MonstersFindAboleth", URL: dnd.BaseURL + dnd.MonstersURL + "/aboleth"},
		{Name: "MonstersList", URL: dnd.BaseURL + dnd.MonstersURL},
		{Name: "RacesFindDwarf", URL: dnd.BaseURL + dnd.RacesURL + "/dwarf"},
		{Name: "RacesFindHuman", URL: dnd.BaseURL + dnd.RacesURL + "/human"},
		{Name: "RacesListRaces", URL: dnd.BaseURL + dnd.RacesURL},
		{Name: "RulesFindAdventuring", URL: dnd.BaseURL + dnd.RulesURL + "/adventuring"},
		{Name: "RulesList", URL: dnd.BaseURL + dnd.RulesURL},
		{Name: "RulesSectionAbilityChecks", URL: dnd.BaseURL + dnd.RuleSectionsURL + "/ability-checks"},
		{Name: "RulesSectionsList", URL: dnd.BaseURL + dnd.RuleSectionsURL},
		{Name: "SpellsList", URL: dnd.BaseURL + dnd.SpellsURL},
		{Name: "SpellsListFiltered", URL: dnd.BaseURL + dnd.SpellsURL + "?level=1&school=evo"},
		{Name: "SpellsFindAcidArrow", URL: dnd.BaseURL + dnd.SpellsURL + "/acid-arrow"},
		{Name: "SubracesFindHighElf", URL: dnd.BaseURL + dnd.SubracesURL + "/high-elf"},
		{Name: "SubracesList", URL: dnd.BaseURL + dnd.SubracesURL},
	}

	for i, e := range m {
		res, err := http.Get(e.URL)
		if err != nil {
			return err
		}
		// close response body
		defer res.Body.Close()

		// read all response body
		data, _ := ioutil.ReadAll(res.Body)
		m[i] = endpoint{
			Name:     e.Name,
			URL:      e.URL,
			Response: string(data),
		}

		// Artificial rate limit to be a good
		// internet citizen
		time.Sleep(100 * time.Millisecond)
	}

	f, err := os.Create("mocks/gen.go")
	if err != nil {
		return err
	}
	defer f.Close()

	tpl = template.Must(template.ParseGlob("templates/mock.tmpl"))
	_ = tpl.ExecuteTemplate(f, "mock.tmpl", m)
	return nil
}

func GenerateService(resource string, model string) error {
	if resource == "" || model == "" {
		return fmt.Errorf("missing required arguments")
	}

	file := strings.ToLower(resource) + ".go"
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	s := &service{strings.Title(resource), strings.Title(model)}
	tpl := template.Must(template.ParseGlob("templates/service.tmpl"))
	return tpl.ExecuteTemplate(f, "service.tmpl", s)
}

func GenerateModel(model string) error {
	file := fmt.Sprintf("models/%s", strings.ToLower(model)+".go")
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	m := strings.Title(model)
	t, err := template.New("model").Parse(`package models
type {{.Model}} struct {}`)
	if err != nil {
		return err
	}

	return t.ExecuteTemplate(f, "model", &service{Model: m})
}

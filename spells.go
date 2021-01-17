package dnd

import (
	"encoding/json"
	"fmt"
	"github.com/brittonhayes/dnd/models"
	"github.com/google/go-querystring/query"
	"gopkg.in/mcuadros/go-defaults.v1"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ Spells = &SpellsService{}

// The Spells interface shows all of the
// available methods for the spells endpoint
type Spells interface {
	ListSpells() (*models.Resource, error)
	FindSpell(name string) (*models.Spells, error)
}

type SpellsService struct {
	// URL is the base URL of the service
	URL     string `default:"https://www.dnd5eapi.co/api"`
	Options *SpellParams
}

// NewSpellsService creates a custom instance of the
// spells service
func NewCustomSpellsService(url string, params *SpellParams) *SpellsService {
	return &SpellsService{URL: url, Options: params}
}

// NewSpellsService creates a default instance of the
// spells service
func NewSpellsService() *SpellsService {
	s := new(SpellsService)
	defaults.SetDefaults(s)
	return s
}

type SpellParams struct {
	Name   string `url:"name"`
	Level  string `url:"level"`
	School string `url:"school"`
}

// ListSpells lists the available spells endpoints
func (s *SpellsService) ListSpells() (*models.Resource, error) {
	q, _ := query.Values(s.Options)
	url := s.URL + SpellsURL
	if s.Options.Level != "" || s.Options.School != "" || s.Options.Name != "" {
		url = fmt.Sprintf("%s?%s", url, q.Encode())
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	spells := new(models.Resource)
	if err := json.Unmarshal(body, &spells); err != nil {
		return nil, err
	}

	return spells, nil
}

// FindSpell searches a specific spell by name
func (s *SpellsService) FindSpell(name string) (*models.Spells, error) {
	if name == "" {
		return nil, fmt.Errorf("missing name argument")
	}

	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.TrimSpace(name)

	url := s.URL + SpellsURL + fmt.Sprintf("/%s", strings.TrimPrefix(name, "/"))
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	spell := new(models.Spells)
	if err := json.Unmarshal(body, &spell); err != nil {
		return nil, err
	}

	return spell, nil
}

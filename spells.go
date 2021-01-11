package dnd

import (
	"encoding/json"
	"fmt"
	"github.com/brittonhayes/dnd/models"
	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"gopkg.in/mcuadros/go-defaults.v1"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ Spells = &SpellsService{}

// The Spells interface shows all of the
// available methods for the spells endpoint
type Spells interface {
	ListSpells() (*models.APIReference, error)
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
	Level  string `url:"level"`
	School string `url:"school"`
}

// ListSpells lists the available spells endpoints
func (s *SpellsService) ListSpells() (*models.APIReference, error) {

	q, _ := query.Values(s.Options)
	url := BaseURL + SpellsURL
	method := "GET"
	if s.Options.Level != "" || s.Options.School != "" {
		url = fmt.Sprintf("%s?%s", url, q.Encode())
	}
	logrus.Debug(url)
	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	spells := new(models.APIReference)
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

	n := strings.TrimSpace(name)
	url := BaseURL + SpellsURL + fmt.Sprintf("/%s", strings.TrimPrefix(n, "/"))
	method := "GET"

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)

	res, err := client.Do(req)
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

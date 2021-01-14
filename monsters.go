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

var _ Monsters = &MonstersService{}

// The Monsters interface shows all of the
// available methods for the monsters endpoint
type Monsters interface {
	ListMonsters() (*models.Resource, error)
	FindMonster(name string) (*models.Monster, error)
}

type MonstersService struct {
	// URL is the base URL of the service
	URL     string `default:"https://www.dnd5eapi.co/api"`
	Options *MonstersParams
}

// NewMonstersService creates a new instance of the
// monsters service
func NewMonstersService() *MonstersService {
	m := new(MonstersService)
	defaults.SetDefaults(m)
	return m
}

// NewMonstersService creates a custom instance of the
// monsters service
func NewCustomMonstersService(url string, params *MonstersParams) *MonstersService {
	return &MonstersService{URL: url, Options: params}
}

type MonstersParams struct {
	ChallengeRating string `url:"challenge_rating"`
}

// ListMonsters available in the API
func (s *MonstersService) ListMonsters() (*models.Resource, error) {

	q, _ := query.Values(s.Options)
	url := s.URL + MonstersURL
	url = fmt.Sprintf("%s?%s", url, q.Encode())

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	monsters := new(models.Resource)
	if err := json.Unmarshal(body, &monsters); err != nil {
		return nil, err
	}

	return monsters, nil
}

// FindMonster fetches a monster's details by name
func (s *MonstersService) FindMonster(name string) (*models.Monster, error) {
	if name == "" {
		return nil, fmt.Errorf("missing name argument")
	}

	n := strings.TrimSpace(name)
	q, _ := query.Values(s.Options)
	url := s.URL + MonstersURL + fmt.Sprintf("/%s", strings.TrimPrefix(n, "/"))
	url = fmt.Sprintf("%s?%s", url, q.Encode())

	if strings.Contains(name, " ") {
		strings.ToLower(name)
		strings.ReplaceAll(name, " ", "-")
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	monster := new(models.Monster)
	if err := json.Unmarshal(body, &monster); err != nil {
		return nil, err
	}

	return monster, nil
}

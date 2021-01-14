package dnd

import (
	"encoding/json"
	"fmt"
	"github.com/brittonhayes/dnd/models"
	"gopkg.in/mcuadros/go-defaults.v1"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ Races = &RacesService{}

// The Races interface shows all of the
// available methods for the monsters endpoint
type Races interface {
	ListRaces() (*models.Resource, error)
	ListSubRaces() (*models.Resource, error)
	FindRace(name string) (*models.Race, error)
	FindSubRace(name string) (*models.SubRace, error)
}

// NewRacesService creates a new instance of the
// races service
func NewRacesService() *RacesService {
	m := new(RacesService)
	defaults.SetDefaults(m)
	return m
}

// NewCustomRacesService creates a custom instance of the
// races service
func NewCustomRacesService(url string) *RacesService {
	return &RacesService{URL: url}
}

type RacesService struct {
	// URL is the base URL of the service
	URL string `default:"https://www.dnd5eapi.co/api"`
}

func (s *RacesService) FindSubRace(name string) (*models.SubRace, error) {
	if name == "" {
		return nil, fmt.Errorf("missing name argument")
	}

	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")

	url := s.URL + SubracesURL + fmt.Sprintf("/%s", strings.TrimPrefix(name, "/"))
	method := "GET"

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	subrace := new(models.SubRace)
	if err := json.Unmarshal(body, &subrace); err != nil {
		return nil, err
	}
	return subrace, nil
}

func (s *RacesService) ListRaces() (*models.Resource, error) {

	url := s.URL + RacesURL
	method := "GET"
	client := &http.Client{}

	req, _ := http.NewRequest(method, url, nil)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	races := new(models.Resource)
	if err := json.Unmarshal(body, &races); err != nil {
		return nil, err
	}

	return races, nil
}

func (s *RacesService) ListSubRaces() (*models.Resource, error) {

	url := s.URL + SubracesURL
	method := "GET"
	client := &http.Client{}

	req, _ := http.NewRequest(method, url, nil)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	sub := new(models.Resource)
	if err := json.Unmarshal(body, &sub); err != nil {
		return nil, err
	}

	return sub, nil
}

func (s *RacesService) FindRace(name string) (*models.Race, error) {
	if name == "" {
		return nil, fmt.Errorf("missing name argument")
	}

	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")

	n := strings.TrimSpace(name)
	url := s.URL + RacesURL + fmt.Sprintf("/%s", strings.TrimPrefix(n, "/"))
	method := "GET"

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	race := new(models.Race)
	if err := json.Unmarshal(body, &race); err != nil {
		return nil, err
	}

	return race, nil
}

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

var _ Equipment = &EquipmentService{}

// The Equipment interface shows all of the
// available methods for the Equipment endpoint
type Equipment interface {
	ListEquipment() (*models.Resource, error)
	FindWeapon(name string) (*models.Weapon, error)
	FindArmor(name string) (*models.Armor, error)
	FindAdventuringGear(name string) (*models.AdventuringGear, error)
	FindEquipmentPack(name string) (*models.EquipmentPack, error)
}

type EquipmentService struct {
	// URL is the base URL of the service
	URL     string `default:"https://www.dnd5eapi.co/api"`
	Options *EquipmentParams
}

type EquipmentParams struct {
	// TODO implement all parameters available for the Equipment service
}

// NewEquipmentService creates a new instance of the
// Equipment service
func NewEquipmentService() *EquipmentService {
	m := new(EquipmentService)
	defaults.SetDefaults(m)
	return m
}

// NewEquipmentService creates a custom instance of the
// Equipment service
func NewCustomEquipmentService(url string, params *EquipmentParams) *EquipmentService {
	return &EquipmentService{URL: url, Options: params}
}

// ListEquipment available in the API
func (s *EquipmentService) ListEquipment() (*models.Resource, error) {
	q, _ := query.Values(s.Options)
	url := s.URL + EquipmentURL
	url = fmt.Sprintf("%s?%s", url, q.Encode())

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	Equipment := new(models.Resource)
	if err := json.Unmarshal(body, &Equipment); err != nil {
		return nil, err
	}

	return Equipment, nil
}

// FindWeapon fetches a Weapon's details by name
func (s *EquipmentService) FindWeapon(name string) (*models.Weapon, error) {
	if name == "" {
		return nil, fmt.Errorf("missing name argument")
	}

	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")

	q, _ := query.Values(s.Options)
	url := s.URL + EquipmentURL + fmt.Sprintf("/%s", strings.TrimPrefix(name, "/"))
	url = fmt.Sprintf("%s?%s", url, q.Encode())

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	Equipment := new(models.Weapon)
	if err := json.Unmarshal(body, &Equipment); err != nil {
		return nil, err
	}

	return Equipment, nil
}

// FindArmor fetches Armor's details by name
func (s *EquipmentService) FindArmor(name string) (*models.Armor, error) {
	if name == "" {
		return nil, fmt.Errorf("missing name argument")
	}

	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")

	q, _ := query.Values(s.Options)
	url := s.URL + EquipmentURL + fmt.Sprintf("/%s", strings.TrimPrefix(name, "/"))
	url = fmt.Sprintf("%s?%s", url, q.Encode())

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	armor := new(models.Armor)
	if err := json.Unmarshal(body, &armor); err != nil {
		return nil, err
	}

	return armor, nil
}

// FindAdventuringGear fetches AdventuringGear's details by name
func (s *EquipmentService) FindAdventuringGear(name string) (*models.AdventuringGear, error) {
	if name == "" {
		return nil, fmt.Errorf("missing name argument")
	}

	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")

	q, _ := query.Values(s.Options)
	url := s.URL + EquipmentURL + fmt.Sprintf("/%s", strings.TrimPrefix(name, "/"))
	url = fmt.Sprintf("%s?%s", url, q.Encode())

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	ag := new(models.AdventuringGear)
	if err := json.Unmarshal(body, &ag); err != nil {
		return nil, err
	}

	return ag, nil
}

// FindEquipmentPack fetches EquipmentPack's details by name
func (s *EquipmentService) FindEquipmentPack(name string) (*models.EquipmentPack, error) {
	if name == "" {
		return nil, fmt.Errorf("missing name argument")
	}

	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")

	q, _ := query.Values(s.Options)
	url := s.URL + EquipmentURL + fmt.Sprintf("/%s", strings.TrimPrefix(name, "/"))
	url = fmt.Sprintf("%s?%s", url, q.Encode())

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	ep := new(models.EquipmentPack)
	if err := json.Unmarshal(body, &ep); err != nil {
		return nil, err
	}

	return ep, nil
}

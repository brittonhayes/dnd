package dnd

// Code generated by internal/gen; DO NOT EDIT.
import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/brittonhayes/dnd/internal/utils"
	"github.com/brittonhayes/dnd/models"
	"github.com/google/go-querystring/query"
	"gopkg.in/mcuadros/go-defaults.v1"
)

var _ Monsters = &MonstersService{}

// Monsters interface covers the methods
// available for the MonstersService
type Monsters interface {
	Find(index string) (*models.Monster, error)
	List() (*models.Resource, error)
}

type MonstersService struct {
	// URL is the base URL of the service
	URL     string `default:"https://www.dnd5eapi.co/api"`
	Options *MonstersParams
}

type MonstersParams struct {
	ChallengeRating string `url:"challenge_rating"`
}

// NewMonstersService creates a new instance of the
// Monsters service
func NewMonstersService() *MonstersService {
	s := new(MonstersService)
	defaults.SetDefaults(s)
	return s
}

// NewMonstersService creates a custom instance of the
// Monsters service
func NewCustomMonstersService(url string, params *MonstersParams) *MonstersService {
	return &MonstersService{URL: url, Options: params}
}

// Find searches a specific monster by name
func (s *MonstersService) Find(index string) (*models.Monster, error) {
	var model = &models.Monster{}

	if index == "" {
		return nil, fmt.Errorf("index not provided")
	}

	q, _ := query.Values(s.Options)
	url := utils.URLWithQuery(s.URL, "/monsters", index, q)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	if err := model.JSON(body); err != nil {
		return nil, err
	}
	return model, nil
}

// List lists the available monsters endpoints
func (s *MonstersService) List() (*models.Resource, error) {
	var model = &models.Resource{}

	q, _ := query.Values(s.Options)
	url := utils.URLWithQuery(s.URL, "/monsters", "", q)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	if err := model.JSON(body); err != nil {
		return nil, err
	}
	return model, nil
}

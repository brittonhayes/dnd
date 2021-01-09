package dnd

import (
	"encoding/json"
	"fmt"
	"github.com/brittonhayes/dnd/models"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ Monsters = &monstersService{}

// The Monsters interface shows all of the
// available methods for the monsters endpoint
type Monsters interface {
	ListMonsters() (*models.APIReference, error)
	FindMonster(name string) (*models.Monster, error)
}

type monstersService struct{}

// ListMonsters available in the API
func (s *monstersService) ListMonsters() (*models.APIReference, error) {
	url := BaseURL + MonstersURL
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	monsters := new(models.APIReference)
	if err := json.Unmarshal(body, &monsters); err != nil {
		return nil, err
	}

	return monsters, nil
}

// FindMonster fetches a monster's details by name
func (s *monstersService) FindMonster(name string) (*models.Monster, error) {
	if name == "" {
		return nil, fmt.Errorf("missing name argument")
	}

	n := strings.TrimSpace(name)
	url := BaseURL + MonstersURL + fmt.Sprintf("/%s", strings.TrimPrefix(n, "/"))
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	monster := new(models.Monster)
	if err := json.Unmarshal(body, &monster); err != nil {
		return nil, err
	}

	return monster, nil
}

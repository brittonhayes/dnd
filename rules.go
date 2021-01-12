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

var _ Rules = &RulesService{}

// The Rules interface shows all of the
// available methods for the rules endpoint
type Rules interface {
	ListRules() (*models.Resource, error)
	ListSections() (*models.Resource, error)
	FindRule(name string) (*models.Rules, error)
	FindSection(name string) (*models.RulesSubsection, error)
}

// NewCustomRulesService allows your to create a rules service
// with a custom base url
func NewCustomRulesService(url string) *RulesService {
	return &RulesService{URL: url}
}

// NewRulesService creates a default instance of the
// rules service
func NewRulesService() *RulesService {
	rs := new(RulesService)
	defaults.SetDefaults(rs)
	return rs
}

type RulesService struct {
	// URL is the base URL of the service
	URL string `default:"https://www.dnd5eapi.co/api"`
}

// ListRules lists the available DnD 5e rules in the API
func (s *RulesService) ListRules() (*models.Resource, error) {
	url := s.URL + RulesURL
	method := "GET"

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	rules := new(models.Resource)
	if err := json.Unmarshal(body, &rules); err != nil {
		return nil, err
	}

	return rules, nil
}

// ListSections lists the available DnD 5e rule subsections in the API
func (s *RulesService) ListSections() (*models.Resource, error) {
	url := s.URL + RuleSectionsURL
	method := "GET"

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	rules := new(models.Resource)
	if err := json.Unmarshal(body, &rules); err != nil {
		return nil, err
	}

	return rules, nil
}

// FindRule allows you to search for specific rules based on their name
func (s *RulesService) FindRule(name string) (*models.Rules, error) {
	if name == "" {
		return nil, fmt.Errorf("missing name argument")
	}

	n := strings.TrimSpace(name)
	url := s.URL + RulesURL + fmt.Sprintf("/%s", strings.TrimPrefix(n, "/"))
	method := "GET"

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	rules := new(models.Rules)
	if err := json.Unmarshal(body, &rules); err != nil {
		return nil, err
	}

	return rules, nil
}

// FindSection allows you to search for a specific ruleset subsection based on its name
func (s *RulesService) FindSection(name string) (*models.RulesSubsection, error) {
	if name == "" {
		return nil, fmt.Errorf("missing name argument")
	}

	n := strings.TrimSpace(name)
	url := s.URL + RuleSectionsURL + fmt.Sprintf("/%s", strings.TrimPrefix(n, "/"))
	method := "GET"

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	rules := new(models.RulesSubsection)
	if err := json.Unmarshal(body, &rules); err != nil {
		return nil, err
	}

	return rules, nil
}

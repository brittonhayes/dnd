package dnd

import (
	"encoding/json"
	"fmt"
	"github.com/brittonhayes/dnd/models"
	"io/ioutil"
	"net/http"
	"strings"
)

var _ Rules = &RulesService{}

// The Rules interface shows all of the
// available methods for the rules endpoint
type Rules interface {
	ListRules() (*models.APIReference, error)
	ListSections() (*models.APIReference, error)
	FindRule(name string) (*models.Rules, error)
	FindSection(name string) (*models.RulesSubsection, error)
}

type RulesService struct{}

// ListRules lists the available DnD 5e rules in the API
func (r *RulesService) ListRules() (*models.APIReference, error) {
	url := BaseURL + RulesURL
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

	rules := new(models.APIReference)
	if err := json.Unmarshal(body, &rules); err != nil {
		return nil, err
	}

	return rules, nil
}

// ListSections lists the available DnD 5e rule subsections in the API
func (r *RulesService) ListSections() (*models.APIReference, error) {
	url := BaseURL + RuleSectionsURL
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

	rules := new(models.APIReference)
	if err := json.Unmarshal(body, &rules); err != nil {
		return nil, err
	}

	return rules, nil
}

// FindRule allows you to search for specific rules based on their name
func (r *RulesService) FindRule(name string) (*models.Rules, error) {
	if name == "" {
		return nil, fmt.Errorf("missing name argument")
	}

	n := strings.TrimSpace(name)
	url := BaseURL + RulesURL + fmt.Sprintf("/%s", strings.TrimPrefix(n, "/"))
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

	rules := new(models.Rules)
	if err := json.Unmarshal(body, &rules); err != nil {
		return nil, err
	}

	return rules, nil
}

// FindSection allows you to search for a specific ruleset subsection based on its name
func (r *RulesService) FindSection(name string) (*models.RulesSubsection, error) {
	if name == "" {
		return nil, fmt.Errorf("missing name argument")
	}

	n := strings.TrimSpace(name)
	url := BaseURL + RuleSectionsURL + fmt.Sprintf("/%s", strings.TrimPrefix(n, "/"))
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

	rules := new(models.RulesSubsection)
	if err := json.Unmarshal(body, &rules); err != nil {
		return nil, err
	}

	return rules, nil
}
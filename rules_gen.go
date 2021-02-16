package dnd

// Code generated by internal/gen; DO NOT EDIT.
import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/brittonhayes/dnd/internal/utils"
	"github.com/brittonhayes/dnd/models"
	"gopkg.in/mcuadros/go-defaults.v1"
)

var _ Rules = &RulesService{}

// Rules interface covers the methods
// available for the RulesService
type Rules interface {
	FindRule(index string) (*models.Rules, error)
	FindSection(index string) (*models.RulesSubsection, error)
	ListRules() (*models.Resource, error)
	ListSections() (*models.Resource, error)
}

type RulesService struct {
	// URL is the base URL of the service
	URL string `default:"https://www.dnd5eapi.co/api"`
}

// NewRulesService creates a new instance of the
// Rules service
func NewRulesService() *RulesService {
	s := new(RulesService)
	defaults.SetDefaults(s)
	return s
}

// NewRulesService creates a custom instance of the
// Rules service
func NewCustomRulesService(url string) *RulesService {
	return &RulesService{URL: url}
}

// FindRule searches for specific rules based on their name
func (s *RulesService) FindRule(index string) (*models.Rules, error) {
	var model = &models.Rules{}

	if index == "" {
		return nil, fmt.Errorf("index not provided")
	}

	url := utils.URL(s.URL, "/rules", index)
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

// FindSection searches for specific rules section based on their name
func (s *RulesService) FindSection(index string) (*models.RulesSubsection, error) {
	var model = &models.RulesSubsection{}

	if index == "" {
		return nil, fmt.Errorf("index not provided")
	}

	url := utils.URL(s.URL, "/rule-sections", index)
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

// ListRules lists available rules endpoints
func (s *RulesService) ListRules() (*models.Resource, error) {
	var model = &models.Resource{}

	url := utils.URL(s.URL, "/rules", "")
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

// ListSections lists available rules sections endpoints
func (s *RulesService) ListSections() (*models.Resource, error) {
	var model = &models.Resource{}

	url := utils.URL(s.URL, "/rule-sections", "")
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
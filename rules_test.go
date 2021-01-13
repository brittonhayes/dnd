package dnd

import (
	"encoding/json"
	"fmt"
	"github.com/brittonhayes/dnd/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/kinbiko/jsonassert"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

func TestRulesService_FindRule(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		mock    mocks.Mock
		wantErr bool
	}{
		{name: "Find rule", args: args{name: "adventuring"}, mock: mocks.RulesMock},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", BaseURL+RulesURL, tt.args.name), httpmock.NewStringResponder(200, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			c := NewRulesService()
			got, err := c.FindRule(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindRule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			j, err := json.Marshal(&got)
			if err != nil {
				t.Errorf("failed to marshall json, %v", err)
			}

			ja := jsonassert.New(t)
			ja.Assertf(string(j), string(tt.mock))
		})
	}
	info := httpmock.GetCallCountInfo()
	logrus.Info(info)
}

func TestRulesService_FindSection(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		mock    mocks.Mock
		wantErr bool
	}{
		{name: "Find rules section", args: args{name: "ability-checks"}, mock: mocks.RulesSectionAbilityChecksMock, wantErr: false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", BaseURL+RuleSectionsURL, tt.args.name), httpmock.NewStringResponder(200, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			got, err := c.Rules.FindSection(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindSection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			j, err := json.Marshal(&got)
			if err != nil {
				t.Errorf("failed to marshall json, %v", err)
			}

			ja := jsonassert.New(t)
			ja.Assertf(string(j), string(tt.mock))
		})
	}
	info := httpmock.GetCallCountInfo()
	logrus.Info(info)
}

func TestRulesService_ListRules(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"List rules", fmt.Sprintf("%s%s", BaseURL, RulesURL), mocks.RulesListMock, 200, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		httpmock.RegisterResponder("GET", tt.url, httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			c := NewRulesService()
			got, err := c.ListRules()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListRules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			j, err := json.Marshal(&got)
			if err != nil {
				t.Errorf("failed to marshall json, %v", err)
			}

			ja := jsonassert.New(t)
			ja.Assertf(string(j), string(tt.mock))

		})
	}
	info := httpmock.GetCallCountInfo()
	logrus.Info(info)
}

func TestRulesService_ListSections(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"List rules sections", fmt.Sprintf("%s%s", BaseURL, RuleSectionsURL), mocks.RulesSectionsListMock, 200, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s", BaseURL, RuleSectionsURL), httpmock.NewStringResponder(200, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			c := NewClient()
			got, err := c.Rules.ListSections()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListSections() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			j, err := json.Marshal(&got)
			if err != nil {
				t.Errorf("failed to marshall json, %v", err)
			}

			ja := jsonassert.New(t)
			ja.Assertf(string(j), string(tt.mock))
		})
	}

	info := httpmock.GetCallCountInfo()
	logrus.Info(info)
}

func TestNewRulesService(t *testing.T) {
	tests := []struct {
		name string
		want *RulesService
	}{
		{"Create custom rules service", &RulesService{URL: BaseURL}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRulesService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRulesService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCustomRulesService(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want *RulesService
	}{
		{"Create custom rules service", args{url: BaseURL}, &RulesService{URL: BaseURL}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomRulesService(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomRulesService() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Basic example of printing a rule as JSON
func ExampleRulesService_FindRule() {
	// Create a client
	c := NewClient()

	// Search for a rule
	r, _ := c.Rules.FindRule("adventuring")

	// Read the results of that rule as JSON
	j, _ := json.MarshalIndent(&r, "", "\t")
	fmt.Println(string(j))
}

// Basic example of printing a rules section as JSON
func ExampleRulesService_FindSection() {
	// Create a client
	c := NewClient()

	// Search for a rule
	r, _ := c.Rules.FindSection("ability-checks")

	// Read the results of that rule section as JSON
	j, _ := json.MarshalIndent(&r, "", "\t")
	fmt.Println(string(j))
}

package dnd_test

import (
	"encoding/json"
	"fmt"
	pkg2 "github.com/brittonhayes/dnd"
	"github.com/brittonhayes/dnd/mocks"
	"github.com/brittonhayes/dnd/models"
	"github.com/jarcoal/httpmock"
	"github.com/kinbiko/jsonassert"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

// Basic example of printing a rule as JSON
func ExampleRulesService_FindRule() {
	// Create a client
	c := pkg2.NewClient()

	// Search for a rule
	r, _ := c.Rules.FindRule("adventuring")

	// Read the results of that rule as JSON
	j, _ := json.MarshalIndent(&r, "", "\t")
	fmt.Println(string(j))
}

func TestRulesService_FindRule(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		mock    mocks.Mock
		want    *models.Rules
		wantErr bool
	}{
		{name: "Find rule", args: args{name: "adventuring"}, mock: mocks.RulesMock, want: &models.Rules{Name: "Adventuring", Index: "adventuring", Desc: "# Adventuring\n", Subsections: []models.RulesSubsection{{Name: "Time", Index: "time", URL: "/api/rule-sections/time"}, {Name: "Movement", Index: "movement", URL: "/api/rule-sections/movement"}, {Name: "The Environment", Index: "the-environment", URL: "/api/rule-sections/the-environment"}, {Name: "Resting", Index: "resting", URL: "/api/rule-sections/resting"}, {Name: "Between Adventures", Index: "between-adventures", URL: "/api/rule-sections/between-adventures"}}, URL: "/api/rules/adventuring"}},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s/%s", BaseURL+RulesURL, tt.want.Index), httpmock.NewStringResponder(200, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			c := &pkg2.Client{}
			got, err := c.Rules.FindRule(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindRule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindRule() got = %v, want %v", got, tt.want)
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

// Basic example of printing a rules section as JSON
func ExampleRulesService_FindSection() {
	// Create a client
	c := pkg2.NewClient()

	// Search for a rule
	r, _ := c.Rules.FindSection("ability-checks")

	// Read the results of that rule section as JSON
	j, _ := json.MarshalIndent(&r, "", "\t")
	fmt.Println(string(j))
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
			c := &pkg2.Client{}
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
		mock    mocks.Mock
		want    *models.APIReference
		wantErr bool
	}{
		{"List rules", mocks.RulesListMock, &models.APIReference{}, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		httpmock.RegisterResponder("GET", BaseURL+RulesURL, httpmock.NewStringResponder(200, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			c := &pkg2.Client{}
			got, err := c.Rules.ListRules()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListRules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListRules() got = %v, want %v", got, tt.want)
			}
		})
	}
	info := httpmock.GetCallCountInfo()
	logrus.Info(info)
}

func TestRulesService_ListSections(t *testing.T) {
	tests := []struct {
		name    string
		mock    mocks.Mock
		want    *models.APIReference
		wantErr bool
	}{
		{"List rules", mocks.RulesSectionsListMock, &models.APIReference{}, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s", BaseURL, RuleSectionsURL), httpmock.NewStringResponder(200, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			c := &pkg2.Client{}
			got, err := c.Rules.ListSections()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListSections() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListSections() got = %v, want %v", got, tt.want)
			}
		})
	}

	info := httpmock.GetCallCountInfo()
	logrus.Info(info)
}

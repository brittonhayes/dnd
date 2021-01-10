package dnd_test

import (
	"encoding/json"
	"fmt"
	"github.com/brittonhayes/dnd"
	"github.com/brittonhayes/dnd/models"
	"reflect"
	"testing"
)

// Basic example of printing a rule as JSON
func ExampleRulesService_FindRule() {
	// Create a client
	c := dnd.NewClient()

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
		want    *models.Rules
		wantErr bool
	}{
		{name: "Find rule", args: args{name: "adventuring"}, want: &models.Rules{
			Name:  "Adventuring",
			Index: "adventuring",
			Desc:  "# Adventuring\n",
			Subsections: []models.RulesSubsection{
				{Name: "Time", Index: "time", URL: "/api/rule-sections/time"},
				{Name: "Movement", Index: "movement", URL: "/api/rule-sections/movement"},
				{Name: "The Environment", Index: "the-environment", URL: "/api/rule-sections/the-environment"},
				{Name: "Resting", Index: "resting", URL: "/api/rule-sections/resting"},
				{Name: "Between Adventures", Index: "between-adventures", URL: "/api/rule-sections/between-adventures"},
			},
			URL: "/api/rules/adventuring",
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &dnd.Client{}
			got, err := c.Rules.FindRule(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindRule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindRule() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// Basic example of printing a rules section as JSON
func ExampleRulesService_FindSection() {
	// Create a client
	c := dnd.NewClient()

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
		want    *models.RulesSubsection
		wantErr bool
	}{
		{"Find rules section", args{name: "ability-checks"}, &models.RulesSubsection{Name: "Ability Checks", Index: "ability-checks", URL: "/api/rule-sections/ability-checks"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &dnd.Client{}
			got, err := c.Rules.FindSection(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindSection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSection() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRulesService_ListRules(t *testing.T) {
	tests := []struct {
		name    string
		want    *models.APIReference
		wantErr bool
	}{
		{"List rules", &models.APIReference{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &dnd.Client{}
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
}

func TestRulesService_ListSections(t *testing.T) {
	tests := []struct {
		name    string
		want    *models.APIReference
		wantErr bool
	}{
		{"List rules", &models.APIReference{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &dnd.Client{}
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
}

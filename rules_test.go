package dnd_test

import (
	"github.com/brittonhayes/dnd"
	"github.com/brittonhayes/dnd/models"
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
		want    *models.Rules
		wantErr bool
	}{
		{"Find rule", args{name: "adventuring"}, &models.Rules{
			Name:  "Adventuring",
			Index: "adventuring",
			Desc:  "# Adventuring\n",
			Subsections: []models.RulesSubsection{
				{"Time", "time", "/api/rule-sections/time"},
				{"Movement", "movement", "/api/rule-sections/movement"},
				{"The Environment", "the-environment", "/api/rule-sections/the-environment"},
				{"Resting", "resting", "/api/rule-sections/resting"},
				{"Between Adventures", "between-adventures", "/api/rule-sections/between-adventures"},
			},
			URL: "/api/rules/adventuring",
		},
			false},
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

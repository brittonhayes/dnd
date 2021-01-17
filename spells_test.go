package dnd

import (
	"encoding/json"
	"fmt"
	"github.com/brittonhayes/dnd/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

func TestSpellsService_FindSpell(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		url     string
		params  *SpellParams
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"Find spells", args{"acid-arrow"}, fmt.Sprintf("%s%s/", BaseURL, SpellsURL), &SpellParams{"", ""}, mocks.SpellsFindAcidArrowMock, 200, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", tt.url+tt.args.name, httpmock.NewStringResponder(tt.status, string(tt.mock)))
			s := NewCustomSpellsService(BaseURL, tt.params)
			got, err := s.FindSpell(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindSpell() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			j, err := json.Marshal(&got)
			if err != nil {
				t.Errorf("failed to marshall json, %v", err)
			}

			if !reflect.DeepEqual(string(j), string(tt.mock)) {
				t.Errorf("\n%v \n\n%v", string(j), tt.mock)
			}

		})
	}

	info := httpmock.GetCallCountInfo()
	logrus.Info(info)
}

func TestSpellsService_ListSpells(t *testing.T) {

	tests := []struct {
		name    string
		params  *SpellParams
		url     string
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"List spells", &SpellParams{"", ""}, fmt.Sprintf("%s%s", BaseURL, SpellsURL), mocks.SpellsListMock, 200, false},
		{"List spells filtered", &SpellParams{"1", "evo"}, fmt.Sprintf("%s%s", BaseURL, SpellsURL), mocks.SpellsListFilteredMock, 200, false},
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		httpmock.RegisterResponder("GET", tt.url, httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := NewCustomSpellsService(BaseURL, tt.params)
			_, err := s.ListSpells()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindSpell() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
	info := httpmock.GetCallCountInfo()
	logrus.Info(info)
}

func TestNewSpellsService(t *testing.T) {
	tests := []struct {
		name string
		want *SpellsService
	}{
		{"Create spell service", &SpellsService{URL: BaseURL, Options: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSpellsService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpellsService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCustomSpellsService(t *testing.T) {
	type args struct {
		url    string
		params *SpellParams
	}
	tests := []struct {
		name string
		args args
		want *SpellsService
	}{
		{"Create custom spell service", args{url: BaseURL, params: &SpellParams{Level: "1", School: "example"}}, &SpellsService{URL: BaseURL, Options: &SpellParams{Level: "1", School: "example"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomSpellsService(tt.args.url, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomSpellsService() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Find a specific spell
func ExampleSpellsService_FindSpell() {
	c := NewClient()
	spell, _ := c.FindSpell("animate-objects")
	fmt.Printf("The spell %s has a range of %s", spell.Name, spell.Range)
}

// Count the number of available spells listed
func ExampleSpellsService_ListSpells_count() {
	s := NewSpellsService()
	s.Options = &SpellParams{
		Level:  "5",
		School: "",
	}

	spells, _ := s.ListSpells()
	fmt.Printf("There are %d spells available", spells.Count)
}

// Create a new spells service and apply custom query params
func ExampleNewSpellsService() {
	s := NewSpellsService()
	s.Options = &SpellParams{
		Level:  "5",
		School: "",
	}
}

// Create a new custom spells service
func ExampleNewCustomSpellsService() {
	s := NewCustomSpellsService(BaseURL, &SpellParams{
		Level:  "2",
		School: "",
	})

	spells, _ := s.ListSpells()
	fmt.Println("Results: ", spells.Results)
}

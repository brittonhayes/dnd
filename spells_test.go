package dnd

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/brittonhayes/dnd/endpoints"
	"github.com/brittonhayes/dnd/internal/mocks"
	"github.com/brittonhayes/dnd/internal/utils"
	"github.com/jarcoal/httpmock"
	"github.com/sirupsen/logrus"
)

func TestSpellsService_FindSpell(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		url     string
		params  *SpellsParams
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"Find spells",
			args{"acid-arrow"},
			fmt.Sprintf("%s%s/", endpoints.BaseURL, endpoints.SpellsURL),
			&SpellsParams{"", "", ""},
			mocks.SpellsFindAcidArrowMock,
			200,
			false,
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", tt.url+tt.args.name, httpmock.NewStringResponder(tt.status, string(tt.mock)))
			s := NewCustomSpellsService(endpoints.BaseURL.String(), tt.params)
			got, err := s.Find(tt.args.name)
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
		params  *SpellsParams
		url     string
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"List spells", &SpellsParams{"", "", ""}, utils.URL(endpoints.BaseURL.String(), endpoints.SpellsURL.String(), ""), mocks.SpellsListMock, 200, false},
		{"List spells filtered", &SpellsParams{"", "1", "evo"}, utils.URL(endpoints.BaseURL.String(), endpoints.SpellsURL.String(), ""), mocks.SpellsListFilteredMock, 200, false},
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		httpmock.RegisterResponder("GET", tt.url, httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := NewCustomSpellsService(endpoints.BaseURL.String(), tt.params)
			_, err := s.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindSpell() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestNewSpellsService(t *testing.T) {
	tests := []struct {
		name string
		want *SpellsService
	}{
		{"Create spell service", &SpellsService{URL: endpoints.BaseURL.String(), Options: nil}},
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
		params *SpellsParams
	}
	tests := []struct {
		name string
		args args
		want *SpellsService
	}{
		{"Create custom spell service", args{url: endpoints.BaseURL.String(), params: &SpellsParams{Level: "1", School: "example"}}, &SpellsService{URL: endpoints.BaseURL.String(), Options: &SpellsParams{Level: "1", School: "example"}}},
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
func ExampleSpellsService_Find() {
	c := NewClient()
	spell, _ := c.Spells.Find("animate-objects")
	fmt.Printf("The spell %s has a range of %s", spell.Name, spell.Range)
}

// Count the number of available spells listed
func ExampleSpellsService_List() {
	s := NewSpellsService()
	s.Options = &SpellsParams{
		Level:  "5",
		School: "",
	}

	spells, _ := s.List()
	fmt.Printf("There are %d spells available", spells.Count)
}

// Create a new spells service and apply custom query params
func ExampleNewSpellsService() {
	s := NewSpellsService()
	s.Options = &SpellsParams{
		Level:  "5",
		School: "",
	}
}

// Create a new custom spells service
func ExampleNewCustomSpellsService() {
	s := NewCustomSpellsService(endpoints.BaseURL.String(), &SpellsParams{
		Level:  "2",
		School: "",
	})

	spells, _ := s.List()
	fmt.Println("Results: ", spells.Results)
}

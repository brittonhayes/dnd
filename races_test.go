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
	"github.com/kinbiko/jsonassert"
)

func TestNewCustomRacesService(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want *RacesService
	}{
		{"New races service", args{url: endpoints.BaseURL.String()}, &RacesService{URL: endpoints.BaseURL.String()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomRacesService(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomRacesService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRacesService(t *testing.T) {
	tests := []struct {
		name string
		want *RacesService
	}{
		{"New races service", &RacesService{URL: endpoints.BaseURL.String()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRacesService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRacesService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRacesService_FindRace(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"Find race dwarf", args{"dwarf"}, mocks.RacesFindDwarfMock, 200, false},
		{"Find race human", args{"human"}, mocks.RacesFindHumanMock, 200, false},
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s/%s", endpoints.BaseURL.String(), endpoints.RacesURL, tt.args.name), httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := NewRacesService()
			got, err := s.FindRace(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindRace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.args.name != "" {
				j, err := json.Marshal(&got)
				if err != nil {
					t.Errorf("failed to marshall json, %v", err)
				}

				ja := jsonassert.New(t)
				ja.Assertf(string(j), string(tt.mock))
			}
		})
	}
}

func TestRacesService_FindSubRace(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"Find subrace high elf", args{"high-elf"}, mocks.SubracesFindHighElfMock, 200, false},
		{"Find subrace high elf with a space", args{"high-elf"}, mocks.SubracesFindHighElfMock, 200, false},
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s/%s", endpoints.BaseURL.String(), endpoints.SubracesURL, tt.args.name), httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := NewRacesService()
			got, err := s.FindSubRace(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindRace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.args.name != "" {
				j, err := json.Marshal(&got)
				if err != nil {
					t.Errorf("failed to marshall json, %v", err)
				}

				ja := jsonassert.New(t)
				ja.Assertf(string(j), string(tt.mock))
			}
		})
	}
}

func TestRacesService_ListSubRaces(t *testing.T) {
	tests := []struct {
		name    string
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"List subraces", mocks.SubracesListMock, 200, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		httpmock.RegisterResponder("GET", utils.URL(endpoints.BaseURL.String(), endpoints.SubracesURL.String(), ""), httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {

			s := NewRacesService()
			got, err := s.ListSubRaces()
			if (err != nil) != tt.wantErr {
				t.Errorf("FindRace() error = %v, wantErr %v", err, tt.wantErr)
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
}

func TestRacesService_ListRaces(t *testing.T) {
	tests := []struct {
		name    string
		URL     string
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"List races", utils.URL(endpoints.BaseURL.String(), endpoints.RacesURL.String(), ""), mocks.RacesListRacesMock, 200, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		httpmock.RegisterResponder("GET", utils.URL(endpoints.BaseURL.String(), endpoints.RacesURL.String(), ""), httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := NewRacesService()
			got, err := s.ListRaces()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListRaces() error = %v, wantErr %v", err, tt.wantErr)
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
}

// Count the number of available races listed
func ExampleRacesService_ListRaces_count() {
	s := NewRacesService()

	races, _ := s.ListRaces()
	fmt.Printf("There are %d races available", races.Count)
}

// Create a new races service and finds a race
func ExampleRacesService_FindRace() {
	s := NewRacesService()
	r, err := s.FindRace("dwarf")
	if err != nil {
		panic(err)
	}
	fmt.Println("Race traits: ", r.Traits)
}

// Create a new custom races service
func ExampleNewCustomRacesService() {
	s := NewCustomRacesService(endpoints.BaseURL.String())

	races, _ := s.ListRaces()
	fmt.Println("Results: ", races.Results)
}

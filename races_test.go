package dnd

import (
	"encoding/json"
	"fmt"
	"github.com/brittonhayes/dnd/mocks"
	"github.com/jarcoal/httpmock"
	"github.com/kinbiko/jsonassert"
	"reflect"
	"testing"
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
		{"New races service", args{url: BaseURL}, &RacesService{URL: BaseURL}},
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
		{"New races service", &RacesService{URL: BaseURL}},
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
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s/%s", BaseURL, RacesURL, tt.args.name), httpmock.NewStringResponder(tt.status, string(tt.mock)))
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
	}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s/%s", BaseURL, SubracesURL, tt.args.name), httpmock.NewStringResponder(tt.status, string(tt.mock)))
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
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s", BaseURL, SubracesURL), httpmock.NewStringResponder(tt.status, string(tt.mock)))
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
		{"List races", fmt.Sprintf("%s%s", BaseURL, RacesURL), mocks.RacesListRacesMock, 200, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s", BaseURL, RacesURL), httpmock.NewStringResponder(tt.status, string(tt.mock)))
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
func ExampleRacesService_ListRacers_Count() {
	s := NewRacesService()

	races, _ := s.ListRaces()
	fmt.Printf("There are %d races available", races.Count)
}

// Create a new races service and apply custom query params
func ExampleRacesService_FindRace() {
	s := NewRacesService()
	s.FindRace("dwarf")
}

// Create a new custom races service
func ExampleNewCustomRacesService() {
	s := NewCustomRacesService(BaseURL)

	races, _ := s.ListRaces()
	fmt.Println("Results: ", races.Results)
}

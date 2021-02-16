package dnd

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/brittonhayes/dnd/endpoints"
	"github.com/brittonhayes/dnd/internal/mocks"
	"github.com/google/go-querystring/query"
	"github.com/jarcoal/httpmock"
	"github.com/kinbiko/jsonassert"
)

func TestMonstersService_ListMonsters(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		params  *MonstersParams
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"List monsters", fmt.Sprintf("%s%s/?challenge_rating=", endpoints.BaseURL, endpoints.MonstersURL), &MonstersParams{""}, mocks.MonstersListMock, 200, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		httpmock.RegisterResponder("GET", tt.url, httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := NewCustomMonstersService(endpoints.BaseURL.String(), tt.params)
			_, err := s.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMonsters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestMonstersService_FindMonster(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		url     string
		params  *MonstersParams
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"Find monster", args{"aboleth"}, fmt.Sprintf("%s%s", endpoints.BaseURL, endpoints.MonstersURL), &MonstersParams{""}, mocks.MonstersFindAbolethMock, 200, false},
		{"Missing name", args{""}, fmt.Sprintf("%s%s", endpoints.BaseURL, endpoints.MonstersURL), &MonstersParams{""}, mocks.MonstersFindAbolethMock, 200, true},
		{"Bad URL name", args{""}, fmt.Sprintf("%s%s/%s", endpoints.BaseURL, endpoints.MonstersURL, "thisdoesntexist"), &MonstersParams{""}, mocks.MonstersFindAbolethMock, 200, true},
		{"Bad query param", args{""}, fmt.Sprintf("%s%s", endpoints.BaseURL, endpoints.MonstersURL), &MonstersParams{"4"}, mocks.MonstersFindAbolethMock, 200, true},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s/%s", endpoints.BaseURL, endpoints.MonstersURL, tt.args.name), httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := NewCustomMonstersService(endpoints.BaseURL.String(), tt.params)
			got, err := s.Find(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if s.Options.ChallengeRating != "" {
				q, _ := query.Values(s.Options)
				u := fmt.Sprintf("%s/%s?challenge_rating=%s", tt.url, tt.args.name, tt.params.ChallengeRating)
				p := tt.url + "/" + tt.args.name + "?" + q.Encode()
				if u != p {
					t.Errorf("FindMonster() query params are not valid, %s != %s", q, p)
				}
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

func TestNewCustomMonstersService(t *testing.T) {
	type args struct {
		url    string
		params *MonstersParams
	}
	tests := []struct {
		name string
		args args
		want *MonstersService
	}{
		{"Create monster service with default url", args{url: endpoints.BaseURL.String(), params: nil},
			&MonstersService{
				URL:     endpoints.BaseURL.String(),
				Options: nil,
			},
		},
		{"Create monster service with custom url", args{url: "exampleurl", params: &MonstersParams{ChallengeRating: "5"}},
			&MonstersService{
				URL:     "exampleurl",
				Options: &MonstersParams{ChallengeRating: "5"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomMonstersService(tt.args.url, tt.args.params); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomMonstersService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMonstersService(t *testing.T) {
	tests := []struct {
		name string
		want *MonstersService
	}{
		{"Create monster service with default url", &MonstersService{
			URL:     endpoints.BaseURL.String(),
			Options: nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMonstersService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMonstersService() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Find a specific monster
func ExampleMonstersService_Find() {
	c := NewClient()
	monster, _ := c.Monsters.Find("aboleth")
	fmt.Printf("The monster %s has a challenge rating of %d", monster.Name, monster.ChallengeRating)
}

// Count the number of available monsters listed
func ExampleMonstersService_List() {
	s := NewMonstersService()
	s.Options = &MonstersParams{
		ChallengeRating: "3",
	}

	monsters, _ := s.List()
	fmt.Printf("There are %d monsters available", monsters.Count)
}

// Create a new monsters service and apply custom query params
func ExampleNewMonstersService() {
	s := NewMonstersService()
	s.Options = &MonstersParams{
		ChallengeRating: "5",
	}
}

// Create a new custom monsters service
func ExampleNewCustomMonstersService() {
	s := NewCustomMonstersService(endpoints.BaseURL.String(), &MonstersParams{
		ChallengeRating: "5",
	})

	monsters, _ := s.List()
	fmt.Println("Results: ", monsters.Count)
}

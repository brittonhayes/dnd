package dnd

import (
	"encoding/json"
	"fmt"
	"github.com/brittonhayes/dnd/mocks"
	"github.com/google/go-querystring/query"
	"github.com/jarcoal/httpmock"
	"github.com/kinbiko/jsonassert"
	"github.com/sirupsen/logrus"
	"testing"
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
		{
			"List monsters",
			fmt.Sprintf("%s%s", BaseURL, MonstersURL),
			&MonstersParams{""},
			mocks.MonstersListMock,
			200,
			false,
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		httpmock.RegisterResponder("GET", tt.url, httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := &MonstersService{
				Options: tt.params,
			}
			_, err := s.ListMonsters()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListMonsters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}

	logrus.Info(httpmock.GetCallCountInfo())
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
		{"Find monster", args{"aboleth"}, fmt.Sprintf("%s%s", BaseURL, MonstersURL), &MonstersParams{""}, mocks.MonstersFindAbolethMock, 200, false},
		{"Missing name", args{""}, fmt.Sprintf("%s%s", BaseURL, MonstersURL), &MonstersParams{""}, mocks.MonstersFindAbolethMock, 200, true},
		{"Bad URL name", args{""}, fmt.Sprintf("%s%s", BaseURL, "thisdoesntexist"), &MonstersParams{""}, mocks.MonstersFindAbolethMock, 200, true},
		{"Bad query param", args{""}, fmt.Sprintf("%s%s", BaseURL, MonstersURL), &MonstersParams{"4"}, mocks.MonstersFindAbolethMock, 200, true},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		httpmock.RegisterResponder("GET", tt.url+"/"+tt.args.name, httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := &MonstersService{
				Options: tt.params,
			}
			got, err := s.FindMonster(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if s.Options.ChallengeRating != ""  {
				q, _ := query.Values(s.Options)
				u := fmt.Sprintf("%s/%s?challenge_rating=%s", tt.url, tt.args.name, tt.params.ChallengeRating)
				p := tt.url+"/"+tt.args.name+"?"+q.Encode()
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

	logrus.Info(httpmock.GetCallCountInfo())
}

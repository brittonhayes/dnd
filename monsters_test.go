package dnd

import (
	"fmt"
	"github.com/brittonhayes/dnd/mocks"
	"github.com/jarcoal/httpmock"
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
		{
			"Find monster",
			args{"aboleth"},
			fmt.Sprintf("%s%s", BaseURL, MonstersURL),
			&MonstersParams{""},
			mocks.MonstersFindAbolethMock,
			200,
			false,
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		httpmock.RegisterResponder("GET", tt.url+"/"+tt.args.name, httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := &MonstersService{
				Options: tt.params,
			}
			_, err := s.FindMonster(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindMonster() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//
			//j, err := json.Marshal(&got)
			//if err != nil {
			//	t.Errorf("failed to marshall json, %v", err)
			//}
			//
			//ja := jsonassert.New(t)
			//ja.Assertf(string(j), string(tt.mock))
		})
	}

	logrus.Info(httpmock.GetCallCountInfo())
}

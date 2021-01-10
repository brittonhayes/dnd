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
		params  SpellParams
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"Find spells", args{"acid-arrow"}, fmt.Sprintf("%s%s/", BaseURL, SpellsURL), SpellParams{"", ""}, mocks.SpellsAcidArrowMock, 200, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", tt.url+tt.args.name, httpmock.NewStringResponder(tt.status, string(tt.mock)))
			s := &SpellsService{Options: tt.params}
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
		params  SpellParams
		url     string
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"List spells", SpellParams{"", ""}, fmt.Sprintf("%s%s", BaseURL, SpellsURL), mocks.SpellsListMock, 200, false},
		{"List spells filtered", SpellParams{"1", "evo"}, fmt.Sprintf("%s%s", BaseURL, SpellsURL), mocks.SpellsListFilterMock, 200, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		httpmock.RegisterResponder("GET", tt.url, httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := &SpellsService{Options: tt.params}
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

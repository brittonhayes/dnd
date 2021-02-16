package dnd

import (
	"reflect"
	"testing"

	"github.com/brittonhayes/dnd/endpoints"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
		want *Client
	}{
		{"Create client", &Client{
			Spells:    NewSpellsService(),
			Monsters:  NewMonstersService(),
			Rules:     NewRulesService(),
			Equipment: NewEquipmentService(),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCustomClient(t *testing.T) {
	type args struct {
		url string
		sp  *SpellsParams
		mp  *MonstersParams
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{"Create custom client and custom services", args{url: endpoints.BaseURL.String()}, &Client{
			Spells:    NewCustomSpellsService(endpoints.BaseURL.String(), nil),
			Monsters:  NewCustomMonstersService(endpoints.BaseURL.String(), nil),
			Rules:     NewCustomRulesService(endpoints.BaseURL.String()),
			Equipment: NewCustomEquipmentService(endpoints.BaseURL.String()),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomClient(tt.args.url, tt.args.sp, tt.args.mp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

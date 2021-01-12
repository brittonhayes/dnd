package dnd

import (
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
		want *Client
	}{
		{"Create client", &Client{
			Spells:   NewSpellsService(),
			Monsters: NewMonstersService(),
			Rules:    NewRulesService(),
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
		sp  *SpellParams
		mp  *MonstersParams
	}
	tests := []struct {
		name string
		args args
		want *Client
	}{
		{"Create custom client", args{url: BaseURL}, &Client{
			Spells:   NewSpellsService(),
			Monsters: NewMonstersService(),
			Rules:    NewRulesService(),
		}},
		{"Create custom client and custom services", args{url: BaseURL}, &Client{
			Spells:   NewCustomSpellsService(BaseURL, nil),
			Monsters: NewCustomMonstersService(BaseURL, nil),
			Rules:    NewCustomRulesService(BaseURL),
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

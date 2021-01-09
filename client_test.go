package dnd_test

import (
	"github.com/brittonhayes/dnd"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name string
		want *dnd.Client
	}{
		{"Create client", &dnd.Client{
			Spells:   &dnd.SpellsService{},
			Monsters: &dnd.MonstersService{},
			Rules:    &dnd.RulesService{},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dnd.NewClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

package models

import "testing"

func TestSpells_GetIndex(t *testing.T) {
	type fields struct {
		Index string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Get spell index", fields{Index: "example"}, "example"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Spells{
				Index: tt.fields.Index,
			}
			if got := s.GetIndex(); got != tt.want {
				t.Errorf("GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpells_GetName(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Get spell name", fields{Name: "example"}, "example"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Spells{
				Name: tt.fields.Name,
			}
			if got := s.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpells_GetURL(t *testing.T) {
	type fields struct {
		URL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Get spell url", fields{URL: "example"}, "example"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Spells{
				URL: tt.fields.URL,
			}
			if got := s.GetURL(); got != tt.want {
				t.Errorf("GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Get the index address of a spell
func ExampleSpells_GetIndex() {
	s := Spells{Index: "example-index"}
	s.GetIndex()
}

// Get the name of a spell
func ExampleSpells_GetName() {
	s := Spells{Name: "example spell name"}
	s.GetName()
}

// Get the URL of a spell
func ExampleSpells_GetURL() {
	s := Spells{Name: "example spell URL"}
	s.GetURL()
}

package models

import "testing"

func TestMagicItem_GetIndex(t *testing.T) {
	type fields struct {
		Index string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Get index", fields{Index: "example-index"}, "example-index"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MagicItem{
				Index: tt.fields.Index,
			}
			if got := m.GetIndex(); got != tt.want {
				t.Errorf("GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMagicItem_GetName(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Get name", fields{Name: "example name"}, "example name"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MagicItem{
				Name: tt.fields.Name,
			}
			if got := m.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMagicItem_GetURL(t *testing.T) {
	type fields struct {
		URL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Get URL", fields{URL: "example.url"}, "example.url"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MagicItem{
				URL: tt.fields.URL,
			}
			if got := m.GetURL(); got != tt.want {
				t.Errorf("GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

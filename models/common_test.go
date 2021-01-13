package models

import "testing"

func TestAPIReference_GetIndex(t *testing.T) {
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
			a := &APIReference{
				Index: tt.fields.Index,
			}
			if got := a.GetIndex(); got != tt.want {
				t.Errorf("GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPIReference_GetName(t *testing.T) {
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
			a := &APIReference{
				Name: tt.fields.Name,
			}
			if got := a.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPIReference_GetURL(t *testing.T) {
	type fields struct {
		URL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Get url", fields{URL: "example.url"}, "example.url"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &APIReference{
				URL: tt.fields.URL,
			}
			if got := a.GetURL(); got != tt.want {
				t.Errorf("GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

package models

import (
	"reflect"
	"testing"
)

func TestAPIReference_GetIndex(t *testing.T) {
	type fields struct {
		Index string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Get index", fields{Index: "_example-index"}, "_example-index"},
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
		{"Get name", fields{Name: "_example name"}, "_example name"},
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
		{"Get url", fields{URL: "_example.url"}, "_example.url"},
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

func TestResource_ResultsNames(t *testing.T) {
	type fields struct {
		Count   int
		Results []APIReference
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{"Get resource result names", fields{Results: []APIReference{{Name: "first"}, {Name: "second"}}}, []string{"first", "second"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Resource{
				Count:   tt.fields.Count,
				Results: tt.fields.Results,
			}
			if got := r.ResultsNames(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ResultsNames() = %v, want %v", got, tt.want)
			}
		})
	}
}

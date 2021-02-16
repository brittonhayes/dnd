package models

import "testing"

func TestAdventuringGear_GetIndex(t *testing.T) {
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
			a := &AdventuringGear{
				Index: tt.fields.Index,
			}
			if got := a.GetIndex(); got != tt.want {
				t.Errorf("GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdventuringGear_GetName(t *testing.T) {
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
			a := &AdventuringGear{
				Name: tt.fields.Name,
			}
			if got := a.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdventuringGear_GetURL(t *testing.T) {
	type fields struct {
		URL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Get URL", fields{URL: "_example.url"}, "_example.url"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &AdventuringGear{
				URL: tt.fields.URL,
			}
			if got := a.GetURL(); got != tt.want {
				t.Errorf("GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArmor_GetIndex(t *testing.T) {
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
			a := &Armor{
				Index: tt.fields.Index,
			}
			if got := a.GetIndex(); got != tt.want {
				t.Errorf("GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArmor_GetName(t *testing.T) {
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
			a := &Armor{
				Name: tt.fields.Name,
			}
			if got := a.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArmor_GetURL(t *testing.T) {
	type fields struct {
		URL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Get URL", fields{URL: "_example.url"}, "_example.url"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Armor{
				URL: tt.fields.URL,
			}
			if got := a.GetURL(); got != tt.want {
				t.Errorf("GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEquipmentPack_GetIndex(t *testing.T) {
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
			e := &EquipmentPack{
				Index: tt.fields.Index,
			}
			if got := e.GetIndex(); got != tt.want {
				t.Errorf("GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEquipmentPack_GetName(t *testing.T) {
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
			e := &EquipmentPack{
				Name: tt.fields.Name,
			}
			if got := e.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEquipmentPack_GetURL(t *testing.T) {
	type fields struct {
		URL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Get URL", fields{URL: "_example.url"}, "_example.url"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EquipmentPack{
				URL: tt.fields.URL,
			}
			if got := e.GetURL(); got != tt.want {
				t.Errorf("GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeapon_GetIndex(t *testing.T) {
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
			w := &Weapon{
				Index: tt.fields.Index,
			}
			if got := w.GetIndex(); got != tt.want {
				t.Errorf("GetIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeapon_GetName(t *testing.T) {
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
			w := &Weapon{
				Name: tt.fields.Name,
			}
			if got := w.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWeapon_GetURL(t *testing.T) {
	type fields struct {
		URL string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Get URL", fields{URL: "_example.url"}, "_example.url"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Weapon{
				URL: tt.fields.URL,
			}
			if got := w.GetURL(); got != tt.want {
				t.Errorf("GetURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

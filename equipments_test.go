package dnd

import (
	"encoding/json"
	"fmt"
	"github.com/brittonhayes/dnd/mocks"
	"github.com/brittonhayes/dnd/models"
	"github.com/jarcoal/httpmock"
	"github.com/kinbiko/jsonassert"
	"github.com/sirupsen/logrus"
	"reflect"
	"testing"
)

func TestEquipmentService_FindAdventuringGear(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		url     string
		params  *EquipmentParams
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"Find abacus", args{"abacus"}, fmt.Sprintf("%s%s", BaseURL, EquipmentURL), &EquipmentParams{}, mocks.EquipmentAdventuringGearMock, 200, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s/%s", BaseURL, EquipmentURL, tt.args.name), httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := NewCustomEquipmentService(BaseURL, tt.params)
			got, err := s.FindAdventuringGear(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindAdventuringGear() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.args.name != "" {
				j, err := json.Marshal(&got)
				if err != nil {
					t.Errorf("failed to marshall json, %v", err)
				}

				ja := jsonassert.New(t)
				ja.Assertf(string(j), string(tt.mock))
			}
		})
	}
	logrus.Info(httpmock.GetCallCountInfo())
}

func TestEquipmentService_FindArmor(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		url     string
		params  *EquipmentParams
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"Find padded armor", args{"padded"}, fmt.Sprintf("%s%s", BaseURL, EquipmentURL), &EquipmentParams{}, mocks.EquipmentArmorMock, 200, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s/%s", BaseURL, EquipmentURL, tt.args.name), httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := NewCustomEquipmentService(BaseURL, tt.params)
			got, err := s.FindArmor(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindArmor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.args.name != "" {
				j, err := json.Marshal(&got)
				if err != nil {
					t.Errorf("failed to marshall json, %v", err)
				}

				ja := jsonassert.New(t)
				ja.Assertf(string(j), string(tt.mock))
			}
		})
	}
	logrus.Info(httpmock.GetCallCountInfo())
}

func TestEquipmentService_FindEquipmentPack(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		url     string
		params  *EquipmentParams
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"Find burglar's pack", args{"padded"}, fmt.Sprintf("%s%s", BaseURL, EquipmentURL), &EquipmentParams{}, mocks.EquipmentPackBurglarsMock, 200, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s/%s", BaseURL, EquipmentURL, tt.args.name), httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := NewCustomEquipmentService(BaseURL, tt.params)
			got, err := s.FindEquipmentPack(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindArmor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.args.name != "" {
				j, err := json.Marshal(&got)
				if err != nil {
					t.Errorf("failed to marshall json, %v", err)
				}

				ja := jsonassert.New(t)
				ja.Assertf(string(j), string(tt.mock))
			}
		})
	}
	logrus.Info(httpmock.GetCallCountInfo())
}

func TestEquipmentService_FindWeapon(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		url     string
		params  *EquipmentParams
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"Find weapon club", args{"padded"}, fmt.Sprintf("%s%s", BaseURL, EquipmentURL), &EquipmentParams{}, mocks.EquipmentWeaponMock, 200, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s/%s", BaseURL, EquipmentURL, tt.args.name), httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := NewCustomEquipmentService(BaseURL, tt.params)
			got, err := s.FindWeapon(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindArmor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.args.name != "" {
				j, err := json.Marshal(&got)
				if err != nil {
					t.Errorf("failed to marshall json, %v", err)
				}

				ja := jsonassert.New(t)
				ja.Assertf(string(j), string(tt.mock))
			}
		})
	}
	logrus.Info(httpmock.GetCallCountInfo())
}

func TestEquipmentService_ListEquipment(t *testing.T) {
	type fields struct {
		URL     string
		Options *EquipmentParams
	}
	tests := []struct {
		name    string
		fields  fields
		want    *models.Resource
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EquipmentService{
				URL:     tt.fields.URL,
				Options: tt.fields.Options,
			}
			got, err := s.ListEquipment()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListEquipment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListEquipment() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCustomEquipmentService(t *testing.T) {
	tests := []struct {
		name string
		want *EquipmentService
	}{
		{"Create custom equipment service with default url", &EquipmentService{
			URL:     BaseURL,
			Options: nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomEquipmentService(tt.want.URL, tt.want.Options); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomEquipmentService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewEquipmentService(t *testing.T) {
	tests := []struct {
		name string
		want *EquipmentService
	}{
		{"Create equipment service with default url", &EquipmentService{
			URL:     BaseURL,
			Options: nil,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEquipmentService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEquipmentService() = %v, want %v", got, tt.want)
			}
		})
	}
}

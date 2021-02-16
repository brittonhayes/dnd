package dnd

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/brittonhayes/dnd/endpoints"
	"github.com/brittonhayes/dnd/internal/mocks"
	"github.com/brittonhayes/dnd/internal/utils"
	"github.com/jarcoal/httpmock"
	"github.com/kinbiko/jsonassert"
	"github.com/sirupsen/logrus"
)

func TestEquipmentService_FindAdventuringGear(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		url     string
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"Find abacus", args{"abacus"}, fmt.Sprintf("%s%s", endpoints.BaseURL, endpoints.EquipmentURL), mocks.EquipmentAdventuringGearMock, 200, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s/%s", endpoints.BaseURL, endpoints.EquipmentURL, tt.args.name), httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := NewCustomEquipmentService(endpoints.BaseURL.String())
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
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"Find padded armor", args{"padded"}, fmt.Sprintf("%s%s", endpoints.BaseURL, endpoints.EquipmentURL), mocks.EquipmentArmorMock, 200, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s/%s", endpoints.BaseURL, endpoints.EquipmentURL, tt.args.name), httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := NewCustomEquipmentService(endpoints.BaseURL.String())
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
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"Find burglar's pack", args{"padded"}, fmt.Sprintf("%s%s", endpoints.BaseURL, endpoints.EquipmentURL), mocks.EquipmentPackBurglarsMock, 200, false},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		httpmock.RegisterResponder("GET", fmt.Sprintf("%s%s/%s", endpoints.BaseURL, endpoints.EquipmentURL, tt.args.name), httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := NewCustomEquipmentService(endpoints.BaseURL.String())
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
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"Find weapon padded", args{"club"}, utils.URL(endpoints.BaseURL.String(), endpoints.EquipmentURL.String(), "club"), mocks.EquipmentWeaponMock, 200, false},
		{"Find weapon missing arg", args{""}, utils.URL(endpoints.BaseURL.String(), endpoints.EquipmentURL.String(), "club"), mocks.EquipmentWeaponMock, 200, true},
		{"Bad URL", args{""}, utils.URL("local.lan", endpoints.EquipmentURL.String(), ""), mocks.EquipmentWeaponMock, 200, true},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	for _, tt := range tests {
		httpmock.RegisterResponder("GET", tt.url, httpmock.NewStringResponder(tt.status, string(tt.mock)))
		t.Run(tt.name, func(t *testing.T) {
			s := NewCustomEquipmentService(endpoints.BaseURL.String())
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
		URL string
	}
	tests := []struct {
		name    string
		fields  fields
		mock    mocks.Mock
		status  int
		wantErr bool
	}{
		{"List equipment", fields{URL: utils.URL(endpoints.BaseURL.String(), endpoints.EquipmentURL.String(), "")}, mocks.EquipmentArmorMock, 200, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &EquipmentService{
				URL: tt.fields.URL,
			}
			_, err := s.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("ListEquipment() error = %v, wantErr %v", err, tt.wantErr)
				return
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
			URL: endpoints.BaseURL.String(),
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomEquipmentService(tt.want.URL); !reflect.DeepEqual(got, tt.want) {
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
			URL: endpoints.BaseURL.String(),
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

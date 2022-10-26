package client

import (
	"net/url"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/consts"
)

func TestConfig_Validate(t *testing.T) {
	type fields struct {
		BaseUrl             *url.URL
		Token               string
		ServiceAccountEmail string
		OrganizationID      string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"empty token", fields{}, true},
		{"empty service account id", fields{Token: "a"}, true},
		{"empty org id", fields{Token: "a", ServiceAccountEmail: consts.SCHWARZ_ORGANIZATION_ID}, true},
		{"all ok", fields{Token: "a", ServiceAccountEmail: consts.SCHWARZ_ORGANIZATION_ID, OrganizationID: consts.SCHWARZ_ORGANIZATION_ID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BaseUrl:             tt.fields.BaseUrl,
				ServiceAccountToken: tt.fields.Token,
				ServiceAccountEmail: tt.fields.ServiceAccountEmail,
				OrganizationID:      tt.fields.OrganizationID,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Config.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfig_SetURL(t *testing.T) {
	type fields struct {
		BaseUrl             *url.URL
		Token               string
		ServiceAccountEmail string
		OrganizationID      string
	}
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"bad url", fields{}, args{"a@b!://#!@$!^&"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BaseUrl:             tt.fields.BaseUrl,
				ServiceAccountToken: tt.fields.Token,
				ServiceAccountEmail: tt.fields.ServiceAccountEmail,
				OrganizationID:      tt.fields.OrganizationID,
			}
			if err := c.SetURL(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Config.SetURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthConfig_Validate(t *testing.T) {
	type fields struct {
		BaseUrl      *url.URL
		ClientID     string
		ClientSecret string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"empty client ID", fields{}, true},
		{"empty client secret", fields{ClientID: consts.SCHWARZ_ORGANIZATION_ID}, true},
		{"all ok", fields{ClientID: consts.SCHWARZ_ORGANIZATION_ID, ClientSecret: "something"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &AuthConfig{
				BaseUrl:      tt.fields.BaseUrl,
				ClientID:     tt.fields.ClientID,
				ClientSecret: tt.fields.ClientSecret,
			}
			if err := c.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("AuthConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthConfig_SetURL(t *testing.T) {
	type fields struct {
		BaseUrl      *url.URL
		ClientID     string
		ClientSecret string
	}
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"bad url", fields{}, args{"a@b!://#!@$!^&"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &AuthConfig{
				BaseUrl:      tt.fields.BaseUrl,
				ClientID:     tt.fields.ClientID,
				ClientSecret: tt.fields.ClientSecret,
			}
			if err := c.SetURL(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("AuthConfig.SetURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

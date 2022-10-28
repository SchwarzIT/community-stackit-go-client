package client

import (
	"net/url"
	"testing"
)

func TestConfig_Validate(t *testing.T) {
	type fields struct {
		BaseUrl             *url.URL
		Token               string
		ServiceAccountEmail string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"empty token", fields{}, true},
		{"empty service account id", fields{Token: "a"}, true},
		{"all ok", fields{Token: "a", ServiceAccountEmail: "b"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				BaseUrl:             tt.fields.BaseUrl,
				ServiceAccountToken: tt.fields.Token,
				ServiceAccountEmail: tt.fields.ServiceAccountEmail,
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
			}
			if err := c.SetURL(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Config.SetURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

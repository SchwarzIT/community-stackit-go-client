package stackit

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

package credentialsgroup

import "testing"

func TestValidateCredentialsGroupID(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		wantErr bool
	}{
		{"all ok", "3969597f-d2f3-4c07-8533-1a4bf8159c0e", false},
		{"err", "123", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateCredentialsGroupID(tt.id); (err != nil) != tt.wantErr {
				t.Errorf("ValidateCredentialsGroupID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

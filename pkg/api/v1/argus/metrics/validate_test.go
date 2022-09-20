package metrics

import "testing"

func TestRetentions_Validate(t *testing.T) {
	type fields struct {
		MetricsRetentionTimeRaw string
		MetricsRetentionTime5m  string
		MetricsRetentionTime1h  string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"all ok", fields{"13months", "720h", "24h"}, false},
		{"too big raw", fields{"14months", "720h", "24h"}, true},
		{"invalid string 1", fields{"abc", "720h", "24h"}, true},
		{"invalid string 2", fields{"720h", "abc", "24h"}, true},
		{"invalid string 3", fields{"720h", "720h", "abc"}, true},
		{"raw < 5m", fields{"720h", "9360h", "24h"}, true},
		{"5m < 1h", fields{"720h", "24h", "9360h"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Config{
				MetricsRetentionTimeRaw: tt.fields.MetricsRetentionTimeRaw,
				MetricsRetentionTime5m:  tt.fields.MetricsRetentionTime5m,
				MetricsRetentionTime1h:  tt.fields.MetricsRetentionTime1h,
			}
			if err := r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Retentions.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

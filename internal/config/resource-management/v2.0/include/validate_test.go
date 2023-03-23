package resourcemanagement

import (
	"testing"

	resourcemanagement "github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0"
)

func TestRole(t *testing.T) {
	type args struct {
		role resourcemanagement.ProjectMemberRole
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"validate admin", args{resourcemanagement.PROJECT_ADMIN}, false},
		{"validate owner", args{resourcemanagement.PROJECT_OWNER}, false},
		{"validate auditor", args{resourcemanagement.PROJECT_AUDITOR}, false},
		{"validate member", args{resourcemanagement.PROJECT_MEMBER}, false},
		{"error", args{"something"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateRole(tt.args.role); (err != nil) != tt.wantErr {
				t.Errorf("Role() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

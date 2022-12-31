package projects

import (
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/services/resource-management/v2.0/generated/projects"
)

func TestRole(t *testing.T) {
	type args struct {
		role projects.ProjectMemberRole
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"validate admin", args{projects.PROJECT_ADMIN}, false},
		{"validate owner", args{projects.PROJECT_OWNER}, false},
		{"validate auditor", args{projects.PROJECT_AUDITOR}, false},
		{"validate member", args{projects.PROJECT_MEMBER}, false},
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

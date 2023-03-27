package clients

import (
	"context"
	"os"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	"github.com/stretchr/testify/assert"
)

func TestTokenFlow_processConfig(t *testing.T) {
	// test env variable loading
	a := os.Getenv(ServiceAccountEmail)
	b := os.Getenv(ServiceAccountToken)
	c := os.Getenv(Environment)

	os.Setenv(ServiceAccountEmail, "test 1")
	os.Setenv(ServiceAccountToken, "test 2")
	os.Setenv(Environment, "dev")

	tf := &TokenFlow{}
	tf.processConfig()

	want := TokenFlowConfig{
		ServiceAccountEmail: "test 1",
		ServiceAccountToken: "test 2",
		Environment:         env.Parse("dev"),
	}
	assert.EqualValues(t, want, *tf.config)

	// revert
	os.Setenv(ServiceAccountEmail, a)
	os.Setenv(ServiceAccountToken, b)
	os.Setenv(Environment, c)

	// Test manual configuration
	type args struct {
		cfg []TokenFlowConfig
	}
	tests := []struct {
		name string
		args args
	}{
		{"test manual 1", args{[]TokenFlowConfig{
			{ServiceAccountEmail: "test 1"},
			{ServiceAccountToken: "test 2"},
			{Environment: "dev"},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TokenFlow{}
			c.processConfig(tt.args.cfg...)
			assert.Equal(t, want, c.GetConfig())
		})
	}
}

func TestTokenFlow_Init(t *testing.T) {
	type args struct {
		ctx context.Context
		cfg []TokenFlowConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"ok", args{context.Background(), []TokenFlowConfig{
			{
				ServiceAccountEmail: "abc",
				ServiceAccountToken: "efg",
			},
		}}, false},
		{"error 1", args{context.Background(), []TokenFlowConfig{
			{
				ServiceAccountEmail: "",
				ServiceAccountToken: "",
			},
		}}, true},
		{"error 2", args{context.Background(), []TokenFlowConfig{
			{
				ServiceAccountEmail: "",
				ServiceAccountToken: "efg",
			},
		}}, true},
	}
	a := os.Getenv(ServiceAccountEmail)
	b := os.Getenv(ServiceAccountToken)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &TokenFlow{}
			assert.EqualValues(t, "", c.GetEnvironment())

			os.Setenv(ServiceAccountEmail, "")
			os.Setenv(ServiceAccountToken, "")
			if err := c.Init(tt.args.ctx, tt.args.cfg...); (err != nil) != tt.wantErr {
				t.Errorf("TokenFlow.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			os.Setenv(ServiceAccountEmail, a)
			os.Setenv(ServiceAccountToken, b)
			if c.config == nil {
				t.Error("config is nil")
			}
			assert.EqualValues(t, "prod", c.GetEnvironment())
			assert.EqualValues(t, c.config.ServiceAccountEmail, c.GetServiceAccountEmail())
		})
	}
}
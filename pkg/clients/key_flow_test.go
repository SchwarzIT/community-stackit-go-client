package clients

import (
	"os"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	"github.com/stretchr/testify/assert"
)

func TestKeyFlow_processConfig(t *testing.T) {

	// test env variable loading
	a := os.Getenv(PrivateKeyPath)
	b := os.Getenv(ServiceAccountKeyPath)
	c := os.Getenv(Environment)

	os.Setenv(PrivateKeyPath, "test 1")
	os.Setenv(ServiceAccountKeyPath, "test 2")
	os.Setenv(Environment, "dev")

	kf := &KeyFlow{}
	kf.processConfig()

	want := KeyFlowConfig{
		PrivateKeyPath:        "test 1",
		PrivateKey:            []byte{},
		ServiceAccountKeyPath: "test 2",
		ServiceAccountKey:     []byte{},
		Environment:           env.Parse("dev"),
	}
	assert.EqualValues(t, want, *kf.config)

	// revert
	os.Setenv(PrivateKeyPath, a)
	os.Setenv(ServiceAccountKeyPath, b)
	os.Setenv(Environment, c)

	type args struct {
		cfg []KeyFlowConfig
	}
	tests := []struct {
		name string
		args args
	}{
		{"test manual 1", args{[]KeyFlowConfig{
			{PrivateKeyPath: "test 1"},
			{ServiceAccountKeyPath: "test 2"},
			{Environment: "dev"},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &KeyFlow{}
			c.processConfig(tt.args.cfg...)
			assert.Equal(t, want, c.GetConfig())
		})
	}
}

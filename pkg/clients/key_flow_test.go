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
	d := os.Getenv(PrivateKey)
	e := os.Getenv(ServiceAccountKey)

	os.Setenv(PrivateKeyPath, "test 1")
	os.Setenv(ServiceAccountKeyPath, "test 2")
	os.Setenv(PrivateKey, "test 3")
	os.Setenv(ServiceAccountKey, "test 4")
	os.Setenv(Environment, "dev")

	kf := &KeyFlow{}
	kf.processConfig()

	want := KeyFlowConfig{
		PrivateKeyPath:        "test 1",
		ServiceAccountKeyPath: "test 2",
		PrivateKey:            []byte("test 3"),
		ServiceAccountKey:     []byte("test 4"),
		Environment:           env.Parse("dev"),
	}
	assert.EqualValues(t, want, *kf.config)

	// revert
	os.Setenv(PrivateKeyPath, a)
	os.Setenv(ServiceAccountKeyPath, b)
	os.Setenv(Environment, c)
	os.Setenv(PrivateKey, d)
	os.Setenv(ServiceAccountKey, e)

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
			{PrivateKey: []byte("test 3")},
			{ServiceAccountKey: []byte("test 4")},
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

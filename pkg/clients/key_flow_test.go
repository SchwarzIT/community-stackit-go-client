package clients

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/SchwarzIT/community-stackit-go-client/pkg/env"
	"github.com/google/uuid"
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
	os.Setenv(Environment, "qa")

	kf := &KeyFlow{}
	kf.processConfig()

	want := KeyFlowConfig{
		PrivateKeyPath:        "test 1",
		ServiceAccountKeyPath: "test 2",
		PrivateKey:            []byte("test 3"),
		ServiceAccountKey:     []byte("test 4"),
		Environment:           env.Parse("qa"),
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
			{Environment: "qa"},
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

func TestKeyFlow_validateConfig(t *testing.T) {
	type fields struct {
		config *KeyFlowConfig
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"ok", fields{&KeyFlowConfig{ServiceAccountKey: []byte("a"), PrivateKey: []byte("b")}}, false},
		{"fail 1", fields{&KeyFlowConfig{ServiceAccountKey: []byte("a")}}, true},
		{"fail 2", fields{&KeyFlowConfig{PrivateKey: []byte("b")}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &KeyFlow{
				config: tt.fields.config,
			}
			if err := c.validateConfig(); (err != nil) != tt.wantErr {
				t.Errorf("KeyFlow.validateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

const saKeyStrPattern = `{
	"active": true,
	"createdAt": "2023-03-23T18:26:20.335Z",
	"credentials": {
	  "aud": "https://stackit-service-account-prod.apps.01.cf.eu01.stackit.cloud",
	  "iss": "stackit@sa.stackit.cloud",
	  "kid": "%s",
	  "sub": "%s"
	},
	"id": "%s",
	"keyAlgorithm": "RSA_2048",
	"keyOrigin": "USER_PROVIDED",
	"keyType": "USER_MANAGED",
	"publicKey": "...",
	"validUntil": "2024-03-22T18:05:41Z"
}`

var saKey = fmt.Sprintf(saKeyStrPattern, uuid.New().String(), uuid.New().String(), uuid.New().String())

func TestKeyFlow_Init(t *testing.T) {
	// Create a temporary file with a random name in the default temporary directory
	tmpfile1, err := ioutil.TempFile("", "sakey")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile1.Name())

	// Create a temporary file with a random name in the default temporary directory
	tmpfile2, err := ioutil.TempFile("", "pkey")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpfile2.Name())

	// Generate a new RSA key pair with a size of 2048 bits
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Encode the private key in PEM format
	privKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	}

	// Print the private and public keys
	pkp := pem.EncodeToMemory(privKeyPEM)

	a := os.Getenv(PrivateKeyPath)
	b := os.Getenv(ServiceAccountKeyPath)
	c := os.Getenv(Environment)
	d := os.Getenv(PrivateKey)
	e := os.Getenv(ServiceAccountKey)

	os.Setenv(PrivateKeyPath, "")
	os.Setenv(ServiceAccountKeyPath, "")
	os.Setenv(PrivateKey, "")
	os.Setenv(ServiceAccountKey, "")
	os.Setenv(Environment, "")

	type args struct {
		cfg []KeyFlowConfig
	}
	tests := []struct {
		name    string
		config  []KeyFlowConfig
		wantErr bool
	}{
		{
			"ok 1",
			[]KeyFlowConfig{
				{PrivateKey: pkp},
				{ServiceAccountKey: []byte(saKey)},
			},
			false,
		},
		{
			"ok 2",
			[]KeyFlowConfig{
				{PrivateKeyPath: tmpfile2.Name()},
				{ServiceAccountKeyPath: tmpfile1.Name()},
			},
			false,
		},
		{
			"bad config 1",
			[]KeyFlowConfig{
				{PrivateKey: pkp},
			},
			true,
		},
		{
			"bad config 2",
			[]KeyFlowConfig{
				{PrivateKey: []byte("somekey")},
				{ServiceAccountKey: []byte(saKey)},
			},
			true,
		},
		{
			"bad config 3",
			[]KeyFlowConfig{},
			true,
		},
		{
			"bad files 1",
			[]KeyFlowConfig{
				{PrivateKeyPath: "somepath"},
				{ServiceAccountKey: []byte(saKey)},
			},
			true,
		},
		{
			"bad files 2",
			[]KeyFlowConfig{
				{PrivateKey: pkp},
				{ServiceAccountKeyPath: "somepath"},
			},
			true,
		},
		{
			"bad file 1",
			[]KeyFlowConfig{
				{PrivateKey: pkp},
				{ServiceAccountKeyPath: tmpfile1.Name()},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &KeyFlow{}
			switch tt.name {
			case "bad file 1":
				if _, err := tmpfile1.Write([]byte("something")); err != nil {
					t.Error(err)
					return
				}
			case "ok 2":
				if _, err := tmpfile1.Write([]byte(saKey)); err != nil {
					t.Error(err)
					return
				}
				if _, err := tmpfile2.Write(pkp); err != nil {
					t.Error(err)
					return
				}
			case "bad config 3":
				if !reflect.DeepEqual(c.GetConfig(), KeyFlowConfig{}) {
					t.Error("config doesn't match")
				}
				if c.GetEnvironment() != "" {
					t.Error("env should be empty before Init")
				}
				if c.GetServiceAccountEmail() != "" {
					t.Error("sa email should be empty before Init")
				}
			}
			if err := c.Init(context.Background(), tt.config...); (err != nil) != tt.wantErr {
				t.Errorf("KeyFlow.Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			if c.GetEnvironment() != "prod" {
				t.Errorf("c.GetEnvironment() = %s != prod", c.GetEnvironment())
			}
			if tt.name == "ok 2" {
				if c.GetServiceAccountEmail() != "stackit@sa.stackit.cloud" {
					t.Errorf("c.GetServiceAccountEmail() = %s != stackit@sa.stackit.cloud", c.GetServiceAccountEmail())
				}
			}
		})
	}
	// revert
	os.Setenv(PrivateKeyPath, a)
	os.Setenv(ServiceAccountKeyPath, b)
	os.Setenv(Environment, c)
	os.Setenv(PrivateKey, d)
	os.Setenv(ServiceAccountKey, e)
}

type errorReader struct{}

func (r *errorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("read error")
}

func (r *errorReader) Close() error {
	return nil
}

const randomTokenResp = `{"access_token":"impEdfCJyeIpDtUOSwMy","expires_in":2726,"refresh_token":"MIaQYdeONTZTwRNgZfxk","scope":"some_scope","token_type":"Bearer"}`

func TestKeyFlow_parseTokenResponse(t *testing.T) {
	r := &errorReader{}
	ok := ioutil.NopCloser(strings.NewReader(randomTokenResp))

	tests := []struct {
		name    string
		res     *http.Response
		wantErr bool
	}{
		{"nil res", nil, true},
		{"bad status", &http.Response{StatusCode: http.StatusBadRequest}, true},
		{"bad body", &http.Response{StatusCode: http.StatusOK, Body: r}, true},
		{"ok", &http.Response{StatusCode: http.StatusOK, Body: ok}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &KeyFlow{}
			if err := c.parseTokenResponse(tt.res); (err != nil) != tt.wantErr {
				t.Errorf("KeyFlow.parseTokenResponse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKeyFlow_validateToken(t *testing.T) {
	// Generate a random private key
	privateKey := make([]byte, 32)
	if _, err := rand.Read(privateKey); err != nil {
		t.Fatal(err)
	}
	type args struct {
		token string
	}
	tests := []struct {
		name    string
		token   string
		want    bool
		wantErr bool
	}{
		{"no token", "", false, false},
		{"fail", "bad token", false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &KeyFlow{
				config: &KeyFlowConfig{
					PrivateKey: privateKey,
				},
			}
			got, err := c.validateToken(tt.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyFlow.validateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("KeyFlow.validateToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

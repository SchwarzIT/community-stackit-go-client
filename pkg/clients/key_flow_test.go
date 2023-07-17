package clients

import (
	"bytes"
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

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestKeyFlow_processConfig(t *testing.T) {

	// test env variable loading
	a := os.Getenv(PrivateKeyPath)
	b := os.Getenv(ServiceAccountKeyPath)
	d := os.Getenv(PrivateKey)
	e := os.Getenv(ServiceAccountKey)

	os.Setenv(PrivateKeyPath, "test 1")
	os.Setenv(ServiceAccountKeyPath, "test 2")
	os.Setenv(PrivateKey, "test 3")
	os.Setenv(ServiceAccountKey, "test 4")

	kf := &KeyFlow{}
	kf.processConfig()

	want := KeyFlowConfig{
		PrivateKeyPath:        "test 1",
		ServiceAccountKeyPath: "test 2",
		PrivateKey:            []byte("test 3"),
		ServiceAccountKey:     []byte("test 4"),
	}
	assert.EqualValues(t, want, *kf.config)

	// revert
	os.Setenv(PrivateKeyPath, a)
	os.Setenv(ServiceAccountKeyPath, b)
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
	d := os.Getenv(PrivateKey)
	e := os.Getenv(ServiceAccountKey)

	os.Setenv(PrivateKeyPath, "")
	os.Setenv(ServiceAccountKeyPath, "")
	os.Setenv(PrivateKey, "")
	os.Setenv(ServiceAccountKey, "")

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
				if c.GetServiceAccountEmail() != "" {
					t.Error("sa email should be empty before Init")
				}
			}
			if err := c.Init(context.Background(), tt.config...); (err != nil) != tt.wantErr {
				t.Errorf("KeyFlow.Init() error = %v, wantErr %v", err, tt.wantErr)
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

func TestClone(t *testing.T) {
	c := &KeyFlow{
		client: &http.Client{},
		config: &KeyFlowConfig{},
		key:    &ServiceAccountKeyPrivateResponse{},
		token:  &TokenResponseBody{},
	}

	clone := c.Clone().(*KeyFlow)

	if !reflect.DeepEqual(c, clone) {
		t.Errorf("Clone() = %v, want %v", clone, c)
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
		{"bad token - ask to recreate", "bad token", false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &KeyFlow{
				config: &KeyFlowConfig{
					PrivateKey: privateKey,
				},
				doer: do,
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

type MockDoer struct {
	mock.Mock
}

func (m *MockDoer) Do(client *http.Client, req *http.Request, cfg *RetryConfig) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestGetJwksJSON(t *testing.T) {
	testCases := []struct {
		name           string
		token          string
		mockResponse   *http.Response
		mockError      error
		expectedResult []byte
		expectedError  error
	}{
		{
			name:  "Success",
			token: "test_token",
			mockResponse: &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"key": "value"}`))),
			},
			mockError:      nil,
			expectedResult: []byte(`{"key": "value"}`),
			expectedError:  nil,
		},
		{
			name:           "Error",
			token:          "test_token",
			mockResponse:   nil,
			mockError:      fmt.Errorf("some error"),
			expectedResult: nil,
			expectedError:  fmt.Errorf("some error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDoer := new(MockDoer)
			mockDoer.On("Do", mock.Anything).Return(tc.mockResponse, tc.mockError)

			c := &KeyFlow{
				config: &KeyFlowConfig{ClientRetry: NewRetryConfig()},
				doer:   mockDoer.Do,
			}

			result, err := c.getJwksJSON(tc.token)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.expectedResult, result)
		})
	}
}

func TestRequestToken(t *testing.T) {
	testCases := []struct {
		name          string
		grant         string
		assertion     string
		mockResponse  *http.Response
		mockError     error
		expectedError error
	}{
		{
			name:      "Success",
			grant:     "test_grant",
			assertion: "test_assertion",
			mockResponse: &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(strings.NewReader(`{"access_token": "test_token"}`)),
			},
			mockError:     nil,
			expectedError: nil,
		},
		{
			name:          "Error",
			grant:         "test_grant",
			assertion:     "test_assertion",
			mockResponse:  nil,
			mockError:     errors.New("request error"),
			expectedError: errors.New("request error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockDoer := new(MockDoer)
			mockDoer.On("Do", mock.Anything).Return(tc.mockResponse, tc.mockError)

			c := &KeyFlow{
				config: &KeyFlowConfig{ClientRetry: NewRetryConfig()},
				doer:   mockDoer.Do,
			}

			res, err := c.requestToken(tc.grant, tc.assertion)

			if tc.expectedError != nil {
				assert.Error(t, err)
				assert.Equal(t, tc.expectedError.Error(), err.Error())
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tc.mockResponse, res)
		})
	}
}

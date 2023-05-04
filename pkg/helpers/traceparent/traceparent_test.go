package traceparent

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	traceID := "0af7651916cd43dd8448eb211c80319c"
	spanID := "b7ad6b7169203331"

	tp := New(traceID, spanID)

	// Check if the Traceparent has the correct trace ID, span ID, and config
	assert.Equal(t, traceID, tp.TraceID)
	assert.Equal(t, spanID, tp.SpanID)
	assert.Equal(t, DefaultConfig.Flag, tp.Config.Flag)
	assert.Equal(t, DefaultConfig.Version, tp.Config.Version)
}

func TestNewCustom(t *testing.T) {
	traceID := "0af7651916cd43dd8448eb211c80319c"
	spanID := "b7ad6b7169203331"

	customConfig := Config{
		Flag:    DoNotRecordFlag,
		Version: "01",
	}

	tp := NewCustom(traceID, spanID, customConfig)

	// Check if the Traceparent has the correct trace ID, span ID, and custom config
	assert.Equal(t, traceID, tp.TraceID)
	assert.Equal(t, spanID, tp.SpanID)
	assert.Equal(t, customConfig.Flag, tp.Config.Flag)
	assert.Equal(t, customConfig.Version, tp.Config.Version)
}

func TestGenerate(t *testing.T) {
	tp, err := Generate()
	assert.NoError(t, err)

	// traceparent should have 55 characters: 2 for version, 32 for traceID, 16 for spanID, 2 for flag, and 3 hyphens
	assert.Equal(t, 55, len(tp.String()))

	// since Generate uses RecordFlag and CurrentVersion, we expect those to be in the traceparent
	assert.Contains(t, tp.Version, string(CurrentVersion))
	assert.Contains(t, tp.Flag, string(RecordFlag))
}

func TestGenerateCustom(t *testing.T) {
	tp, err := GenerateCustom(Config{DoNotRecordFlag, "01"})
	assert.NoError(t, err)

	// traceparent should have 55 characters: 2 for version, 32 for traceID, 16 for spanID, 2 for flag, and 3 hyphens
	assert.Equal(t, 55, len(tp.String()))

	// since GenerateCustom uses the provided flag and version, we expect those to be in the traceparent
	assert.Contains(t, tp.Flag, string(DoNotRecordFlag))
	assert.Contains(t, tp.Version, "01")
}

func TestSetHeader(t *testing.T) {
	tp, _ := Generate()

	// Create a new http.Request
	req, err := http.NewRequest("GET", "http://test.io", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Call SetHeader with the Traceparent
	tp.SetHeader(req)

	// Check if the "Traceparent" header was set correctly
	assert.Equal(t, tp.String(), req.Header.Get("Traceparent"))

	// Test with a nil http.Request
	var nilReq *http.Request
	tp.SetHeader(nilReq)
	assert.Nil(t, nilReq)
}

package traceparent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	tp, err := Generate()
	assert.NoError(t, err)

	// traceparent should have 55 characters: 2 for version, 32 for traceID, 16 for spanID, 2 for flag, and 3 hyphens
	assert.Equal(t, 55, len(tp))

	// since Generate uses RecordFlag and CurrentVersion, we expect those to be in the traceparent
	assert.Contains(t, tp, string(CurrentVersion))
	assert.Contains(t, tp, string(RecordFlag))
}

func TestGenerateCustom(t *testing.T) {
	tp, err := GenerateCustom(Config{DoNotRecordFlag, "01"})
	assert.NoError(t, err)

	// traceparent should have 55 characters: 2 for version, 32 for traceID, 16 for spanID, 2 for flag, and 3 hyphens
	assert.Equal(t, 55, len(tp))

	// since GenerateCustom uses the provided flag and version, we expect those to be in the traceparent
	assert.Contains(t, tp, string(DoNotRecordFlag))
	assert.Contains(t, tp, "01")
}

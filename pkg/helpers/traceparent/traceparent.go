package traceparent

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

type Flag string
type Version string

const (
	DoNotRecordFlag Flag    = "00"
	RecordFlag      Flag    = "01"
	CurrentVersion  Version = "00"
)

type Config struct {
	Flag    Flag
	Version Version
}
type Traceparent struct {
	Config
	TraceID string
	SpanID  string
}

func (t *Traceparent) String() string {
	return fmt.Sprintf("%s-%s-%s-%s", t.Version, t.TraceID, t.SpanID, t.Flag)
}

func (t *Traceparent) SetHeader(req *http.Request) {
	if req == nil {
		return
	}
	req.Header.Set("Traceparent", t.String())
}

func GenerateCustom(c Config) (*Traceparent, error) {
	// Generate a random 16 byte trace ID
	traceID, err := generateRandomHex(16)
	if err != nil {
		return nil, errors.Wrap(err, "error generating traceID")
	}

	// Generate a random 8 byte span ID
	spanID, err := generateRandomHex(8)
	if err != nil {
		return nil, errors.Wrap(err, "error generating spanID")
	}

	return &Traceparent{
		c,
		traceID,
		spanID,
	}, nil
}

func generateRandomHex(size int) (string, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func Generate() (*Traceparent, error) {
	return GenerateCustom(Config{
		RecordFlag,
		CurrentVersion,
	})
}

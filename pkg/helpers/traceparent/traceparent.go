package traceparent

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

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

var DefaultConfig = Config{
	RecordFlag,
	CurrentVersion,
}

func New(traceID, spanID string) *Traceparent {
	return &Traceparent{
		DefaultConfig,
		traceID,
		spanID,
	}
}

func NewCustom(traceID, spanID string, c Config) *Traceparent {
	return &Traceparent{
		c,
		traceID,
		spanID,
	}
}

func Parse(s string) (*Traceparent, error) {
	sl := strings.Split(s, "-")
	if len(sl) != 4 {
		return nil, errors.New("unexpected traceparent structure")
	}

	var f Flag
	switch sl[0] {
	case string(RecordFlag):
		f = RecordFlag
	case string(DoNotRecordFlag):
		f = DoNotRecordFlag
	default:
		return nil, fmt.Errorf("unknown flag '%s'", sl[0])
	}

	var v Version
	switch sl[3] {
	case string(CurrentVersion):
		v = CurrentVersion
	default:
		return nil, fmt.Errorf("unknown version '%s'", sl[3])
	}

	if sl[1] == "" {
		return nil, errors.New("traceID can't be empty")
	}

	if sl[2] == "" {
		return nil, errors.New("spanID can't be empty")
	}
	return NewCustom(sl[1], sl[2], Config{f, v}), nil
}

func (t *Traceparent) String() string {
	return fmt.Sprintf("%s-%s-%s-%s", t.Version, t.TraceID, t.SpanID, t.Flag)
}

func (t *Traceparent) Pretty() string {
	return fmt.Sprintf("traceparent:\n- version:  %s\n- trace id: %s\n- span id:  %s\n- flag:     %s\n", t.Version, t.TraceID, t.SpanID, t.Flag)
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
	return GenerateCustom(DefaultConfig)
}

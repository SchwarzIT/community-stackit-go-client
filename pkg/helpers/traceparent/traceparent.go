package traceparent

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

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

func GenerateCustom(c Config) (string, error) {
	// Generate a random 16 byte trace ID
	traceID, err := generateRandomHex(16)
	if err != nil {
		return "", errors.Wrap(err, "error generating traceID")
	}

	// Generate a random 8 byte span ID
	spanID, err := generateRandomHex(8)
	if err != nil {
		return "", errors.Wrap(err, "error generating spanID")
	}

	// Construct the traceparent header
	return fmt.Sprintf("%s-%s-%s-%s", c.Version, traceID, spanID, c.Flag), nil
}

func generateRandomHex(size int) (string, error) {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func Generate() (string, error) {
	return GenerateCustom(Config{
		RecordFlag,
		CurrentVersion,
	})
}

package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"opsie/types"

	"github.com/sony/sonyflake"
)

var flake = sonyflake.NewSonyflake(sonyflake.Settings{})

const maxInt64 = 1<<63 - 1

func GenerateID() types.ID {
	for {
		id, err := flake.NextID()
		if err != nil {
			panic(err)
		}

		if id <= uint64(maxInt64) {
			return types.ID(id)
		}
		// Retry if overflow
	}
}


// GenerateSessionKey generates a cryptographically secure random session key
func GenerateSessionKey() (string, error) {
	b := make([]byte, 32) // 32 bytes = 256 bits
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate session key: %w", err)
	}
	return hex.EncodeToString(b), nil // 64-character hex string
}
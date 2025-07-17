package gen

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func GenerateHash(str string) string {
	input := str + time.Now().Format(time.RFC3339Nano)
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}
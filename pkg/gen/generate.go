package gen

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHash(str string) string {
	hash := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hash[:])
}
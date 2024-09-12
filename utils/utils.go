package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"time"
)

func GenerateID(task string) string {
	hasher := sha1.New()
	hasher.Write([]byte(task + time.Now().String()))
	return hex.EncodeToString(hasher.Sum(nil))
}

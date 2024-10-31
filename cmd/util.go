package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateUniqueID() (string, error) {
	timestamp := fmt.Sprintf("%d", time.Now().UnixNano())
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	randomPart := hex.EncodeToString(bytes)
	return timestamp + randomPart, nil
}

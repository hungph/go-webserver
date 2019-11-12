package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"github.com/google/uuid"
	"time"
)

func EncryptString(originString string) string {
	hasher := sha256.New()
	hasher.Write([]byte(originString))
	encryptedString := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return encryptedString
}

func RandomString() string {
	uuid, err := uuid.NewUUID()

	if err != nil {
		return ""
	}

	return uuid.String()
}

func TimeDiff(startTime int64) int64 {
	return (time.Now().UnixNano() / 1000000) - startTime
}

package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// SHA256Sum ...
func SHA256Sum(str string) string {
	buf := sha256.Sum256([]byte(str))
	return hex.EncodeToString(buf[:])
}

// RandBytes ...
func RandBytes(size int) []byte {
	b := make([]byte, size)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return b
}

// IsHashString ...
func IsHashString(str string) bool {
	res, err := hex.DecodeString(str)
	if err != nil {
		return false
	}
	return len(res) == 32
}

// Time returns the time with millisecond precision.
func Time(t time.Time) time.Time {
	return time.Unix(0, t.UnixNano()/1e6*1e6)
}

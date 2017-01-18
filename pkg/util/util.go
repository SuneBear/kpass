package util

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"

	"github.com/google/uuid"
)

// IsUUID ...
func IsUUID(id string) bool {
	if _, err := uuid.Parse(id); err == nil {
		return true
	}
	return false
}

// NewUUID return a UUID with given space
func NewUUID(space string) uuid.UUID {
	id, err := uuid.Parse(space)
	if err != nil {
		id = uuid.NewSHA1(uuid.NameSpaceOID, []byte(space))
	}
	return uuid.NewSHA1(id, RandBytes(16))
}

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

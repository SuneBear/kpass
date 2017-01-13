package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"sync"

	"golang.org/x/crypto/pbkdf2"
)

var globalCrypto = New([]byte("KPass"))

// Global ...
func Global() *Crypto {
	return globalCrypto
}

// Reset ...
func Reset(salt []byte) {
	globalCrypto = New(salt)
}

// SHA256Sum ...
func SHA256Sum(str string) string {
	buf := sha256.Sum256([]byte(str))
	return hex.EncodeToString(buf[:])
}

// IsHashString ...
func IsHashString(str string) bool {
	res, err := hex.DecodeString(str)
	if err != nil {
		return false
	}
	return len(res) == 32
}

// Crypto ...
type Crypto struct {
	salt []byte
	hash hash.Hash
	mu   sync.Mutex
}

// New ...
func New(salt []byte) *Crypto {
	return &Crypto{salt: salt, hash: hmac.New(sha256.New, salt)}
}

// AESKey ...
func (c *Crypto) AESKey(key, pass string) string {
	buf := c.hmacSum(fmt.Sprintf("%s.%s", key, pass))
	return base64.StdEncoding.EncodeToString(buf)
}

// EncryptUserPass ...
func (c *Crypto) EncryptUserPass(userID, userPass string) string {
	b := pbkdf2.Key([]byte(userPass), c.salt, 1024, 32, func() hash.Hash {
		return hmac.New(sha256.New, []byte(userID))
	})
	return hex.EncodeToString(b)
}

// ValidateUserPass ...
func (c *Crypto) ValidateUserPass(userID, userPass, dbUserPass string) bool {
	return c.EncryptUserPass(userID, userPass) == dbUserPass
}

// EncryptData ...
func (c *Crypto) EncryptData(key, plainData string) (string, error) {
	block, err := aes.NewCipher(c.hmacSum(key))
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()

	cipherData := make([]byte, blockSize+len(plainData))
	iv := cipherData[:blockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(cipherData[blockSize:], []byte(plainData))
	return base64.StdEncoding.EncodeToString(cipherData), nil
}

// DecryptData ...
func (c *Crypto) DecryptData(key, cipherData string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(cipherData)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(c.hmacSum(key))
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	iv := data[:blockSize]

	plainData := make([]byte, len(data)-blockSize)
	stream := cipher.NewCTR(block, iv)
	stream.XORKeyStream(plainData, data[blockSize:])
	return string(plainData), nil
}

func (c *Crypto) hmacSum(str string) []byte {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.hash.Reset()
	c.hash.Write([]byte(str))
	return c.hash.Sum(nil)
}

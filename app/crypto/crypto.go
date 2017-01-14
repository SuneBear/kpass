package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
	"hash"
	"io"
	"sync"

	"golang.org/x/crypto/pbkdf2"
)

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
func (c *Crypto) AESKey(userPass, dbPass string) string {
	buf := c.hmacSum(userPass + dbPass)
	return base64.RawURLEncoding.EncodeToString(buf)
}

// EncryptUserPass ...
func (c *Crypto) EncryptUserPass(userID, userPass string) string {
	iv := RandBytes(16)
	b := c.encryptUserPass(iv, []byte(userPass+userID))
	return base64.RawURLEncoding.EncodeToString(b)
}

func (c *Crypto) encryptUserPass(iv, pass []byte) []byte {
	b := pbkdf2.Key(pass, c.salt, 1025, 32, func() hash.Hash {
		return hmac.New(sha256.New, iv)
	})
	return append(b, iv...)
}

// ValidateUserPass ...
func (c *Crypto) ValidateUserPass(userID, userPass, dbPass string) bool {
	a, err := base64.RawURLEncoding.DecodeString(dbPass)
	if err != nil {
		return false
	}
	b := c.encryptUserPass(a[32:], []byte(userPass+userID))
	return subtle.ConstantTimeCompare(a, b) == 1
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
	return base64.RawURLEncoding.EncodeToString(cipherData), nil
}

// DecryptData ...
func (c *Crypto) DecryptData(key, cipherData string) (string, error) {
	data, err := base64.RawURLEncoding.DecodeString(cipherData)
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

// SHA256Sum ...
func SHA256Sum(str string) string {
	buf := sha256.Sum256([]byte(str))
	return hex.EncodeToString(buf[:])
}

// RandBytes ...
func RandBytes(size int) []byte {
	b := make([]byte, size)
	rand.Read(b)
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

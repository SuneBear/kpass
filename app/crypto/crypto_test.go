package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrypto(t *testing.T) {
	c := New([]byte("KPass"))

	t.Run("AESKey", func(t *testing.T) {
		assert := assert.New(t)

		pass := SHA256Sum("test pass")
		k := c.AESKey("admin", pass)
		assert.True(len(k) == 44)
	})

	t.Run("EncryptUserPass and ValidateUserPass", func(t *testing.T) {
		assert := assert.New(t)

		pass := SHA256Sum("test pass")
		epass := c.EncryptUserPass("admin", pass)
		assert.True(c.ValidateUserPass("admin", pass, epass))
	})

	t.Run("EncryptData and DecryptData", func(t *testing.T) {
		assert := assert.New(t)

		pass := SHA256Sum("test pass")
		key := c.AESKey("admin", pass)

		edata, _ := c.EncryptData(key, "Hello! 中国")
		data, _ := c.DecryptData(key, edata)
		assert.Equal("Hello! 中国", data)
	})
}

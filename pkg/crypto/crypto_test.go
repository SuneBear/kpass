package crypto

import (
	"encoding/base64"
	"testing"

	"github.com/seccom/kpass/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestCrypto(t *testing.T) {
	c := New([]byte("KPass"))

	t.Run("AESKey", func(t *testing.T) {
		assert := assert.New(t)

		pass := util.SHA256Sum("test pass")
		k := c.AESKey("admin", pass)
		b, _ := base64.RawURLEncoding.DecodeString(k)
		assert.True(len(b) == 32)
	})

	t.Run("EncryptUserPass and ValidateUserPass", func(t *testing.T) {
		assert := assert.New(t)

		pass := util.SHA256Sum("test pass")
		epass := c.EncryptUserPass("admin", pass)
		assert.True(c.ValidateUserPass("admin", pass, epass))
	})

	t.Run("EncryptData and DecryptData", func(t *testing.T) {
		assert := assert.New(t)

		pass := util.SHA256Sum("test pass")
		key := c.AESKey("admin", pass)

		edata, _ := c.EncryptData(key, "Hello! 中国")
		data, _ := c.DecryptData(key, edata)
		assert.Equal("Hello! 中国", data)
	})
}

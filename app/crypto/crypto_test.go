package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCrypto(t *testing.T) {
	t.Run("AESKey", func(t *testing.T) {
		assert := assert.New(t)

		pass := SHA256Sum("test pass")
		k := Global().AESKey("admin", pass)
		assert.True(len(k) == 44)
	})

	t.Run("EncryptUserPass and ValidateUserPass", func(t *testing.T) {
		assert := assert.New(t)

		pass := SHA256Sum("test pass")
		epass := Global().EncryptUserPass("admin", pass)
		assert.True(Global().ValidateUserPass("admin", pass, epass))
	})

	t.Run("EncryptData and DecryptData", func(t *testing.T) {
		assert := assert.New(t)

		pass := SHA256Sum("test pass")
		key := Global().AESKey("admin", pass)

		edata, _ := Global().EncryptData(key, "Hello! 中国")
		data, _ := Global().DecryptData(key, edata)
		assert.Equal("Hello! 中国", data)
	})
}

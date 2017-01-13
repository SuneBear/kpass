package pkg

import (
	"time"

	"github.com/SermoDigital/jose/jwt"
	"github.com/teambition/gear-auth"
)

// Jwt middleware ...
var Jwt *auth.JWT

// InitJwt ...
func InitJwt(expiresIn time.Duration, keys ...interface{}) {
	Jwt = auth.NewJWT(keys...)
	Jwt.SetExpiresIn(expiresIn)
	Jwt.SetValidator(&jwt.Validator{
		Fn: func(c jwt.Claims) (err error) {
			id := c.Get("id").(string)
			key := c.Get("key").(string)
			if key, err = Crypto.DecryptData(id, key); err == nil {
				// decrypt key
				c.Set("key", key)
			}
			return nil
		},
	})
}

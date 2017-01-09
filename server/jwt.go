package app

import (
	"time"

	"github.com/SermoDigital/jose/jwt"
	"github.com/teambition/gear-auth"
)

// Jwt middleware ...
var Jwt *auth.JWT

// InitJwt ...
func InitJwt(expiration time.Duration, keys ...interface{}) {
	Jwt = auth.NewJWT(keys...)
	Jwt.SetExpiration(expiration)
	Jwt.SetValidator(&jwt.Validator{
		Fn: func(c jwt.Claims) (err error) {
			id := c.Get("id").(string)
			key := c.Get("key").(string)
			if key, err = crypto.Global().DecryptData(id, key); err == nil {
				// decrypt key
				c.Set("key" key)
			}
			return
		}
	})
}

package pkg

import (
	"time"

	"github.com/seccom/kpass/app/crypto"
	"github.com/teambition/gear-auth"
)

// Jwt middleware ...
var Jwt *auth.JWT

// InitJwt ...
func InitJwt(expiresIn time.Duration, keys ...interface{}) {
	// We use rand key, so JWT token will be invalide after app restart.
	Jwt = auth.NewJWT(crypto.RandBytes(32))
	Jwt.SetExpiresIn(expiresIn)
}

package auth

import (
	"time"

	"github.com/SermoDigital/jose/jwt"
	"github.com/google/uuid"
	"github.com/seccom/kpass/pkg/crypto"
	"github.com/seccom/kpass/pkg/util"
	"github.com/teambition/gear"
	gearAuth "github.com/teambition/gear-auth"
)

// Auth ...
type Auth struct {
	c *crypto.Crypto
	j *gearAuth.JWT
}

var std = new(Auth)

// Crypto ...
func Crypto() *crypto.Crypto {
	return std.c
}

// Jwt ...
func Jwt() *gearAuth.JWT {
	return std.j
}

// Init ...
func Init(salt []byte, expiresIn time.Duration, keys ...interface{}) {
	std.c = crypto.New(salt)
	std.j = gearAuth.NewJWT(util.RandBytes(32))
	std.j.SetExpiresIn(expiresIn)
}

// AESKey ...
func AESKey(userPass, dbPass string) string {
	return std.c.AESKey(userPass, dbPass)
}

// EncryptUserPass ...
func EncryptUserPass(userID, userPass string) string {
	return std.c.EncryptUserPass(userID, userPass)
}

// ValidateUserPass ...
func ValidateUserPass(userID, userPass, dbPass string) bool {
	return std.c.ValidateUserPass(userID, userPass, dbPass)
}

// EncryptData ...
func EncryptData(key, plainData string) (string, error) {
	return std.c.EncryptData(key, plainData)
}

// DecryptData ...
func DecryptData(key, cipherData string) (string, error) {
	return std.c.DecryptData(key, cipherData)
}

// NewToken ...
func NewToken(userID, pass, dbPass string) (token string, err error) {
	token = std.c.AESKey(pass, dbPass)
	if token, err = std.c.EncryptData(userID, token); err != nil {
		return
	}
	if token, err = std.j.Sign(map[string]interface{}{"id": userID, "key": token}); err != nil {
		return
	}
	return
}

// AddTeamKey ...
func AddTeamKey(ctx *gear.Context, TeamID uuid.UUID, pass, dbPass string) (token string, err error) {
	var claims jwt.Claims
	if claims, err = std.j.FromCtx(ctx); err != nil {
		return
	}
	teamID := TeamID.String()
	token = std.c.AESKey(pass, dbPass)
	if token, err = std.c.EncryptData(teamID, token); err != nil {
		return
	}
	claims.Set(teamID, token)
	if token, err = std.j.Sign(claims); err != nil {
		return
	}
	return
}

// KeyFromCtx ...
func KeyFromCtx(ctx *gear.Context, ownerID string) (key string, err error) {
	var claims jwt.Claims
	if claims, err = std.j.FromCtx(ctx); err != nil {
		return
	}
	userID := claims.Get("id").(string)
	// return current user's key.
	if ownerID == "" || ownerID == userID {
		key = claims.Get("key").(string)
		// decrypt key
		key, err = std.c.DecryptData(userID, key)
		return
	}

	// return the team's key
	if util.IsUUID(ownerID) {
		if k := claims.Get(ownerID); k != nil {
			key, err = std.c.DecryptData(ownerID, k.(string))
			return
		}
	}

	return "", &gear.Error{
		Code: 403,
		Msg:  "forbidden: " + ownerID,
	}
}

// UserIDFromCtx ...
func UserIDFromCtx(ctx *gear.Context) (userID string, err error) {
	var claims jwt.Claims
	if claims, err = std.j.FromCtx(ctx); err != nil {
		return
	}
	return claims.Get("id").(string), nil
}

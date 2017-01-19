package auth

import (
	"time"

	josejwt "github.com/SermoDigital/jose/jwt"
	"github.com/google/uuid"
	"github.com/seccom/kpass/pkg/util"
	"github.com/teambition/gear"
	au "github.com/teambition/gear-auth"
	"github.com/teambition/gear-auth/crypto"
	"github.com/teambition/gear-auth/jwt"
)

// Auth ...
type Auth struct {
	*au.Auth
}

var std = &Auth{au.New(util.RandBytes(32))} // use a rand key
var Middleware = std.Serve

// Default returns the global Auth instance
func Default() *Auth {
	return std
}

// Crypto ...
func Crypto() *crypto.Crypto {
	return std.Crypto()
}

// JWT ...
func JWT() *jwt.JWT {
	return std.JWT()
}

// Init ...
func Init(salt []byte, expiresIn time.Duration) {
	std.SetCrypto(crypto.New(salt))
	std.JWT().SetExpiresIn(expiresIn)
}

// AESKey ...
func AESKey(userPass, dbPass string) string {
	return std.Crypto().AESKey(userPass, dbPass)
}

// EncryptUserPass ...
func EncryptUserPass(userID, userPass string) string {
	return std.Crypto().EncryptUserPass(userID, userPass)
}

// ValidateUserPass ...
func ValidateUserPass(userID, userPass, dbPass string) bool {
	return std.Crypto().ValidateUserPass(userID, userPass, dbPass)
}

// EncryptData ...
func EncryptData(key, plainData string) (string, error) {
	return std.Crypto().EncryptData(key, plainData)
}

// DecryptData ...
func DecryptData(key, cipherData string) (string, error) {
	return std.Crypto().DecryptData(key, cipherData)
}

// Sign ...
func Sign(c map[string]interface{}) (string, error) {
	return std.JWT().Sign(c)
}

// NewToken ...
func NewToken(userID, pass, dbPass string) (token string, err error) {
	token = AESKey(pass, dbPass)
	if token, err = EncryptData(userID, token); err != nil {
		return
	}
	if token, err = Sign(map[string]interface{}{"id": userID, "key": token}); err != nil {
		return
	}
	return
}

// AddTeamKey ...
func AddTeamKey(ctx *gear.Context, TeamID uuid.UUID, pass, dbPass string) (token string, err error) {
	var claims josejwt.Claims
	if claims, err = FromCtx(ctx); err != nil {
		return
	}
	teamID := TeamID.String()
	token = AESKey(pass, dbPass)
	if token, err = EncryptData(teamID, token); err != nil {
		return
	}
	claims.Set(teamID, token)
	if token, err = Sign(claims); err != nil {
		return
	}
	return
}

// FromCtx ...
func FromCtx(ctx *gear.Context) (josejwt.Claims, error) {
	return std.FromCtx(ctx)
}

// KeyFromCtx ...
func KeyFromCtx(ctx *gear.Context, ownerID string) (key string, err error) {
	var claims josejwt.Claims
	if claims, err = FromCtx(ctx); err != nil {
		return
	}
	userID := claims.Get("id").(string)
	// return current user's key.
	if ownerID == "" || ownerID == userID {
		key = claims.Get("key").(string)
		// decrypt key
		key, err = DecryptData(userID, key)
		return
	}

	// return the team's key
	if util.IsUUID(ownerID) {
		if k := claims.Get(ownerID); k != nil {
			key, err = DecryptData(ownerID, k.(string))
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
	var claims josejwt.Claims
	if claims, err = FromCtx(ctx); err != nil {
		return
	}
	return claims.Get("id").(string), nil
}

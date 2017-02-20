package auth

import (
	"time"

	josejwt "github.com/SermoDigital/jose/jwt"
	"github.com/seccom/kpass/src/util"
	"github.com/teambition/gear"
	gearauth "github.com/teambition/gear-auth"
	"github.com/teambition/gear-auth/crypto"
	"github.com/teambition/gear-auth/jwt"
)

// Auth ...
type Auth struct {
	*gearauth.Auth
}

// use a rand key for JWT token, means that tokens will be invalid after service restart.
var std = &Auth{gearauth.New(util.RandBytes(32))}

// Middleware use to ...
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

// SignPass ...
func SignPass(userID, userPass string) string {
	return std.Crypto().SignPass(userID, userPass)
}

// VerifyPass ...
func VerifyPass(userID, userPass, dbPass string) bool {
	return std.Crypto().VerifyPass(userID, userPass, dbPass)
}

// Encrypt ...
func Encrypt(key, plainData []byte) ([]byte, error) {
	return std.Crypto().Encrypt(key, plainData)
}

// Decrypt ...
func Decrypt(key, cipherData []byte) ([]byte, error) {
	return std.Crypto().Decrypt(key, cipherData)
}

// EncryptText ...
func EncryptText(key, plainData string) (string, error) {
	return std.Crypto().EncryptText(key, plainData)
}

// DecryptText ...
func DecryptText(key, cipherData string) (string, error) {
	return std.Crypto().DecryptText(key, cipherData)
}

// Sign ...
func Sign(c map[string]interface{}, expiresIn ...time.Duration) (string, error) {
	return std.JWT().Sign(c, expiresIn...)
}

// Verify ...
func Verify(token string) (josejwt.Claims, error) {
	return std.JWT().Verify(token)
}

// NewToken ...
func NewToken(userID string) (string, error) {
	return Sign(map[string]interface{}{"id": userID})
}

// AddTeamKey ...
func AddTeamKey(ctx *gear.Context, TeamID util.OID, pass, checkPass string) (token string, err error) {
	var claims josejwt.Claims
	if claims, err = FromCtx(ctx); err != nil {
		return
	}

	key := AESKey(pass, checkPass)
	userID := claims.Get("id").(string)
	if key, err = EncryptText(userID, key); err != nil {
		return
	}
	claims.Set("team"+TeamID.String(), key)
	return Sign(claims)
}

// AddShareKey ...
func AddShareKey(ctx *gear.Context, ShareID util.OID, pass, key string) (token string, err error) {
	var claims josejwt.Claims
	if claims, err = FromCtx(ctx); err != nil {
		return
	}

	userID := claims.Get("id").(string)
	if key, err = DecryptText(SignPass(userID, pass), key); err != nil {
		return
	}
	if key, err = EncryptText(userID, key); err != nil {
		return
	}
	claims.Set("share"+ShareID.String(), key)
	return Sign(claims)
}

// FromCtx ...
func FromCtx(ctx *gear.Context) (josejwt.Claims, error) {
	return std.FromCtx(ctx)
}

// KeyFromCtx ...
func KeyFromCtx(ctx *gear.Context, ID util.OID, keyType string) (key string, err error) {
	var claims josejwt.Claims
	if claims, err = FromCtx(ctx); err != nil {
		return
	}

	id := ID.String()
	userID := claims.Get("id").(string)
	switch keyType {
	case "team", "share":
		// return the team's key or share's key
		if k := claims.Get(keyType + id); k != nil {
			return DecryptText(userID, k.(string))
		}
	}

	return "", &gear.Error{
		Code: 403,
		Msg:  "forbidden: " + id,
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

// SignedFileKey ...
func SignedFileKey(FileID util.OID, key string) (signed string, err error) {
	if signed, err = Sign(map[string]interface{}{"key": key}, time.Minute); err != nil {
		return
	}
	signed, err = EncryptText(FileID.String(), signed)
	return
}

// FileKeyFromSigned ...
func FileKeyFromSigned(FileID util.OID, signed string) (key string, err error) {
	if signed, err = DecryptText(FileID.String(), signed); err != nil {
		return
	}
	var claims josejwt.Claims
	if claims, err = Verify(signed); err == nil {
		key = claims.Get("key").(string)
	}
	return
}

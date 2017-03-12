package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"time"

	josejwt "github.com/SermoDigital/jose/jwt"
	"github.com/seccom/kpass/src/util"
	"github.com/teambition/crypto-go"
	"github.com/teambition/gear"
	gearauth "github.com/teambition/gear-auth"
	"github.com/teambition/gear-auth/jwt"
)

// Auth ...
type Auth struct {
	*gearauth.Auth
	salt []byte
}

// use a rand key for JWT token, means that tokens will be invalid after service restart.
var std = &Auth{gearauth.New(util.RandBytes(32)), nil}

// Middleware use to ...
var Middleware = std.Serve

// Default returns the global Auth instance
func Default() *Auth {
	return std
}

// JWT ...
func JWT() *jwt.JWT {
	return std.JWT()
}

// Init ...
func Init(salt []byte, expiresIn time.Duration) {
	std.salt = salt
	std.JWT().SetExpiresIn(expiresIn)
}

// HmacSum ...
func HmacSum(str string) string {
	return hex.EncodeToString(crypto.HmacSum(sha256.New, std.salt, []byte(str)))
}

// AESKey ...
func AESKey(userPass, dbPass string) string {
	return base64.URLEncoding.EncodeToString(crypto.SHA256Hmac([]byte(userPass), []byte(dbPass)))
}

// SignPass ...
func SignPass(userID, userPass string) string {
	return crypto.SignPass(std.salt, userID, userPass)
}

// VerifyPass ...
func VerifyPass(userID, userPass, dbPass string) bool {
	return crypto.VerifyPass(std.salt, userID, userPass, dbPass)
}

// Encrypt ...
func Encrypt(key, plainData []byte) ([]byte, error) {
	return crypto.AESEncrypt(std.salt, key, plainData)
}

// Decrypt ...
func Decrypt(key, cipherData []byte) ([]byte, error) {
	return crypto.AESDecrypt(std.salt, key, cipherData)
}

// EncryptStr ...
func EncryptStr(key, plainData string) (string, error) {
	return crypto.AESEncryptStr(std.salt, key, plainData)
}

// DecryptStr ...
func DecryptStr(key, cipherData string) (string, error) {
	return crypto.AESDecryptStr(std.salt, key, cipherData)
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
func NewToken(userID, pass, checkPass string) (string, error) {
	key := AESKey(pass, checkPass)
	key, err := EncryptStr(HmacSum(userID), key)
	if err != nil {
		return "", err
	}
	return Sign(map[string]interface{}{"id": userID, "key": key})
}

// KeyFromCtx ...
func KeyFromCtx(ctx *gear.Context) (key string, err error) {
	var claims josejwt.Claims
	if claims, err = FromCtx(ctx); err != nil {
		return
	}

	userID := claims.Get("id").(string)
	key = claims.Get("key").(string)
	if key, err = DecryptStr(HmacSum(userID), key); err != nil {
		err = &gear.Error{Code: 403, Msg: err.Error()}
	}
	return
}

// UserIDFromCtx ...
func UserIDFromCtx(ctx *gear.Context) (userID string, err error) {
	var claims josejwt.Claims
	if claims, err = FromCtx(ctx); err != nil {
		return
	}
	return claims.Get("id").(string), nil
}

// FromCtx ...
func FromCtx(ctx *gear.Context) (josejwt.Claims, error) {
	return std.FromCtx(ctx)
}

// SignedFileKey ...
func SignedFileKey(FileID util.OID, key string) (signed string, err error) {
	if signed, err = Sign(map[string]interface{}{"key": key}, time.Minute); err != nil {
		return
	}
	signed, err = EncryptStr(FileID.String(), signed)
	return
}

// FileKeyFromSigned ...
func FileKeyFromSigned(FileID util.OID, signed string) (key string, err error) {
	if signed, err = DecryptStr(FileID.String(), signed); err != nil {
		return
	}
	var claims josejwt.Claims
	if claims, err = Verify(signed); err == nil {
		key = claims.Get("key").(string)
	}
	return
}

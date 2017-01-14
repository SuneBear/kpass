package pkg

import (
	"github.com/SermoDigital/jose/jwt"
	"github.com/google/uuid"
	"github.com/seccom/kpass/app/crypto"
	"github.com/teambition/gear"
)

// Crypto ...
var Crypto *cryptoJ

// InitCrypto ...
func InitCrypto(salt []byte) {
	Crypto = &cryptoJ{crypto.New(salt)}
}

type cryptoJ struct {
	*crypto.Crypto
}

func (c *cryptoJ) NewToken(id, userPass, dbPass string) (token string, err error) {
	token = c.AESKey(userPass, dbPass)
	if token, err = c.EncryptData(id, token); err != nil {
		return
	}
	if token, err = Jwt.Sign(map[string]interface{}{"id": id, "key": token}); err != nil {
		return
	}
	return
}

func (c *cryptoJ) AddTeamKey(claims jwt.Claims, id uuid.UUID, userPass, dbPass string) (token string, err error) {
	teamID := id.String()
	token = c.AESKey(userPass, dbPass)
	if token, err = c.EncryptData(teamID, token); err != nil {
		return
	}
	claims.Set(teamID, token)
	if token, err = Jwt.Sign(claims); err != nil {
		return
	}
	return
}

func (c *cryptoJ) UserKeyFromCtx(ctx *gear.Context) (key string, err error) {
	var claims jwt.Claims
	if claims, err = Jwt.FromCtx(ctx); err != nil {
		return
	}
	id := claims.Get("id").(string)
	key = claims.Get("key").(string)
	// decrypt key
	key, err = c.DecryptData(id, key)
	return
}

func (c *cryptoJ) TeamKeyFromCtx(ctx *gear.Context, id uuid.UUID) (key string, err error) {
	var claims jwt.Claims
	teamID := id.String()
	if claims, err = Jwt.FromCtx(ctx); err != nil {
		return
	}
	if k := claims.Get(teamID); k != nil {
		key, err = c.DecryptData(teamID, k.(string))
		return
	}
	return "", &gear.Error{
		Code: 403,
		Msg:  "forbidden: " + teamID,
	}
}

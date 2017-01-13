package pkg

import (
	"github.com/SermoDigital/jose/jwt"
	"github.com/google/uuid"
	"github.com/seccom/kpass/app/crypto"
	"github.com/teambition/gear"
)

// Crypto middleware ...
var Crypto *cryptoJ

// InitCrypto ...
func InitCrypto(salt []byte) {
	Crypto = &cryptoJ{crypto.New(salt)}
}

type cryptoJ struct {
	*crypto.Crypto
}

func (c *cryptoJ) NewToken(id, pass string) (token string, err error) {
	token = c.AESKey(id, pass)
	// encrypt key
	if token, err = c.EncryptData(id, token); err != nil {
		return
	}
	if token, err = Jwt.Sign(map[string]interface{}{"id": id, "key": token}); err != nil {
		return
	}
	return
}

func (c *cryptoJ) AddTeamKey(claims jwt.Claims, id uuid.UUID, pass string) (token string, err error) {
	strID := id.String()
	token = c.AESKey(strID, pass)
	if token, err = c.EncryptData(strID, token); err != nil {
		return "", err
	}
	claims.Set(strID, token)
	if token, err = Jwt.Sign(claims); err != nil {
		return "", err
	}
	return
}

func (c *cryptoJ) UserKeyFromCtx(ctx *gear.Context) (key string, err error) {
	claims, _ := Jwt.FromCtx(ctx) // It was validated by pre-middleware
	id := claims.Get("id").(string)
	key = claims.Get("key").(string)
	// decrypt key
	key, err = c.DecryptData(id, key)
	return
}

func (c *cryptoJ) TeamKeyFromCtx(ctx *gear.Context, id uuid.UUID) (key string, err error) {
	strID := id.String()
	claims, _ := Jwt.FromCtx(ctx) // It was validated by pre-middleware
	if k := claims.Get(strID); k != nil {
		key, err = c.DecryptData(strID, k.(string))
		return
	}
	return "", &gear.Error{
		Code: 403,
		Msg:  "invalid team id: " + strID,
	}
}

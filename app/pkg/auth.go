package pkg

import (
	"github.com/SermoDigital/jose/jwt"
	"github.com/google/uuid"
	"github.com/seccom/kpass/app/crypto"
	"github.com/teambition/gear"
)

type authT struct {
	*crypto.Crypto
}

// Auth ...
var Auth *authT

// InitAuth ...
func InitAuth(salt []byte) {
	Auth = &authT{crypto.New(salt)}
}

func (a *authT) NewToken(id, userPass, dbPass string) (token string, err error) {
	token = a.AESKey(userPass, dbPass)
	if token, err = a.EncryptData(id, token); err != nil {
		return
	}
	if token, err = Jwt.Sign(map[string]interface{}{"id": id, "key": token}); err != nil {
		return
	}
	return
}

func (a *authT) AddTeamKey(claims jwt.Claims, id uuid.UUID, userPass, dbPass string) (token string, err error) {
	teamID := id.String()
	token = a.AESKey(userPass, dbPass)
	if token, err = a.EncryptData(teamID, token); err != nil {
		return
	}
	claims.Set(teamID, token)
	if token, err = Jwt.Sign(claims); err != nil {
		return
	}
	return
}

func (a *authT) KeyFromCtx(ctx *gear.Context, ownerID string) (key string, err error) {
	var claims jwt.Claims
	if claims, err = Jwt.FromCtx(ctx); err != nil {
		return
	}
	userID := claims.Get("id").(string)
	// return current user's key.
	if ownerID == "" || ownerID == userID {
		key = claims.Get("key").(string)
		// decrypt key
		key, err = a.DecryptData(userID, key)
		return
	}

	// return the team's key
	if IsUUID(ownerID) {
		if k := claims.Get(ownerID); k != nil {
			key, err = a.DecryptData(ownerID, k.(string))
			return
		}
	}

	return "", &gear.Error{
		Code: 403,
		Msg:  "forbidden: " + ownerID,
	}
}

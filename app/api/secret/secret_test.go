package secretAPI_test

import (
	"encoding/hex"
	"testing"

	"github.com/DavidCai1993/request"
	"github.com/google/uuid"
	"github.com/seccom/kpass/app"
	"github.com/seccom/kpass/app/crypto"
	"github.com/seccom/kpass/app/dao"
	"github.com/stretchr/testify/assert"
	"github.com/teambition/gear"
)

func TestSecretAPI(t *testing.T) {
	srv := app.New("", "test").Start()
	defer srv.Close()

	host := "http://" + srv.Addr().String()
	_, _, accessToken := func() (id, pass, accessToken string) {
		id = "test" + hex.EncodeToString(crypto.RandBytes(8))
		pass = crypto.SHA256Sum(crypto.SHA256Sum(crypto.RandPass(8, 2, 2)))
		_, err := request.Post(host+"/join").
			Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
			Send(map[string]interface{}{"id": id, "pass": pass}).
			End()

		if err != nil {
			panic(err)
		}

		res, err := request.Post(host+"/login").
			Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
			Send(map[string]interface{}{"username": id, "password": pass, "grant_type": "password"}).
			JSON()
		if err != nil {
			panic(err)
		}
		accessToken = "Bearer " + (*res.(*map[string]interface{}))["access_token"].(string)
		return
	}()

	entry := new(dao.EntrySum)
	_, err := request.Post(host+"/entries").
		Set(gear.HeaderAuthorization, accessToken).
		Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
		Send(map[string]interface{}{"name": "test"}).
		JSON(entry)
	assert.Nil(t, err)

	var secretID uuid.UUID
	t.Run("Add a secret", func(t *testing.T) {
		assert := assert.New(t)
		res := new(dao.SecretResult)

		_, err := request.Post(host+"/entries/"+entry.ID.String()+"/secrets").
			Set(gear.HeaderAuthorization, accessToken).
			Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
			Send(map[string]interface{}{"name": "test secret", "url": "test.com", "password": "123456"}).
			JSON(res)
		assert.Nil(err)

		secretID = res.ID
		assert.Equal("test secret", res.Name)
		assert.Equal("test.com", res.URL)
		assert.Equal("123456", res.Pass)
	})

	t.Run("Update a secret", func(t *testing.T) {
		assert := assert.New(t)
		res := new(dao.SecretResult)

		_, err := request.Put(host+"/entries/"+entry.ID.String()+"/secrets/"+secretID.String()).
			Set(gear.HeaderAuthorization, accessToken).
			Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
			Send(map[string]interface{}{"url": "www.test.com", "password": "abcdefg", "note": "note"}).
			JSON(res)
		assert.Nil(err)

		secretID = res.ID
		assert.Equal("test secret", res.Name)
		assert.Equal("www.test.com", res.URL)
		assert.Equal("abcdefg", res.Pass)
		assert.Equal("note", res.Note)
	})

	t.Run("Delete a secret", func(t *testing.T) {
		assert := assert.New(t)

		res, err := request.Delete(host+"/entries/"+entry.ID.String()+"/secrets/"+secretID.String()).
			Set(gear.HeaderAuthorization, accessToken).End()
		assert.Nil(err)
		assert.Equal(204, res.StatusCode)

		res2 := new(dao.EntryResult)
		_, err = request.Get(host+"/entries/"+entry.ID.String()).
			Set(gear.HeaderAuthorization, accessToken).
			JSON(res)
		assert.Nil(err)
		assert.True(len(res2.Secrets) == 0)
	})
}

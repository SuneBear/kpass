package api_test

import (
	"encoding/hex"
	"testing"

	"github.com/DavidCai1993/request"
	"github.com/google/uuid"
	"github.com/seccom/kpass/pkg"
	"github.com/seccom/kpass/pkg/schema"
	"github.com/seccom/kpass/pkg/util"
	"github.com/stretchr/testify/assert"
	"github.com/teambition/gear"
)

func TestSecretAPI(t *testing.T) {
	app := pkg.New("", "test")
	srv := app.Start()
	defer srv.Close()

	host := "http://" + srv.Addr().String()
	_, _, accessToken := func() (id, pass, accessToken string) {
		id = "test" + hex.EncodeToString(util.RandBytes(8))
		pass = util.SHA256Sum(util.SHA256Sum(util.RandPass(8, 2, 2)))
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

	entry := new(schema.EntrySum)
	_, err := request.Post(host+"/entries").
		Set(gear.HeaderAuthorization, accessToken).
		Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
		Send(map[string]interface{}{"name": "test"}).
		JSON(entry)
	assert.Nil(t, err)

	var secretID uuid.UUID
	t.Run("Add a secret", func(t *testing.T) {
		assert := assert.New(t)
		res := new(schema.SecretResult)

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
		res := new(schema.SecretResult)

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

		res2 := new(schema.EntryResult)
		_, err = request.Get(host+"/entries/"+entry.ID.String()).
			Set(gear.HeaderAuthorization, accessToken).
			JSON(res)
		assert.Nil(err)
		assert.True(len(res2.Secrets) == 0)
	})
}

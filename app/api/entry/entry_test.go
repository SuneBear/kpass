package entryAPI_test

import (
	"encoding/hex"
	"testing"

	"strings"

	"github.com/DavidCai1993/request"
	"github.com/google/uuid"
	"github.com/seccom/kpass/app"
	"github.com/seccom/kpass/app/crypto"
	"github.com/seccom/kpass/app/dao"
	"github.com/stretchr/testify/assert"
	"github.com/teambition/gear"
)

func TestEntryAPI(t *testing.T) {
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

	t.Run("Find with no content", func(t *testing.T) {
		assert := assert.New(t)
		res := new([]dao.EntrySum)

		_, err := request.Get(host+"/entries").
			Set(gear.HeaderAuthorization, accessToken).
			JSON(res)
		assert.Nil(err)
		assert.True(len(*res) == 0)
	})

	var entryID uuid.UUID
	t.Run("Create a entry", func(t *testing.T) {
		assert := assert.New(t)
		res := new(dao.EntrySum)

		_, err := request.Post(host+"/entries").
			Set(gear.HeaderAuthorization, accessToken).
			Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
			Send(map[string]interface{}{"name": "test"}).
			JSON(res)
		assert.Nil(err)

		entryID = res.ID
		assert.Equal("test", res.Name)
		assert.Equal("", res.Category)
		assert.Equal(0, res.Priority)
	})

	t.Run("Find a entry", func(t *testing.T) {
		assert := assert.New(t)
		res := new(dao.EntryResult)

		_, err := request.Get(host+"/entries/"+entryID.String()).
			Set(gear.HeaderAuthorization, accessToken).
			JSON(res)
		assert.Nil(err)

		assert.Equal(entryID, res.ID)
		assert.Equal("test", res.Name)
		assert.Equal("", res.Category)
		assert.Equal(0, res.Priority)
		assert.True(strings.Contains(res.String(), `"secrets":[]`))
		assert.True(strings.Contains(res.String(), `"shares":[]`))
	})

	var secretID uuid.UUID
	t.Run("Add a secret", func(t *testing.T) {
		assert := assert.New(t)
		res := new(dao.SecretResult)

		_, err := request.Post(host+"/entries/"+entryID.String()+"/secrets").
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

	t.Run("Find a entry again", func(t *testing.T) {
		assert := assert.New(t)
		res := new(dao.EntryResult)

		_, err := request.Get(host+"/entries/"+entryID.String()).
			Set(gear.HeaderAuthorization, accessToken).
			JSON(res)
		assert.Nil(err)

		assert.Equal(entryID, res.ID)
		assert.Equal("test", res.Name)
		assert.Equal("", res.Category)
		assert.Equal(0, res.Priority)
		assert.True(strings.Contains(res.String(), `"shares":[]`))

		secret := res.Secrets[0]
		assert.Equal(secretID, secret.ID)
		assert.Equal("test secret", secret.Name)
		assert.Equal("test.com", secret.URL)
		assert.Equal("123456", secret.Pass)
	})

	t.Run("Find user entries again", func(t *testing.T) {
		assert := assert.New(t)
		res := new([]*dao.EntrySum)

		_, err := request.Get(host+"/entries").
			Set(gear.HeaderAuthorization, accessToken).
			JSON(res)
		assert.Nil(err)

		entry := (*res)[0]
		assert.Equal(entryID, entry.ID)
		assert.Equal("test", entry.Name)
		assert.Equal("", entry.Category)
		assert.Equal(0, entry.Priority)
		assert.False(strings.Contains(entry.String(), "secrets"))
		assert.False(strings.Contains(entry.String(), "shares"))
	})

	t.Run("Update a entry", func(t *testing.T) {
		assert := assert.New(t)
		res := new(dao.EntrySum)

		_, err := request.Put(host+"/entries/"+entryID.String()).
			Set(gear.HeaderAuthorization, accessToken).
			Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
			Send(map[string]interface{}{"name": "test1", "category": "银行卡", "priority": 1}).
			JSON(res)
		assert.Nil(err)

		assert.Equal(entryID, res.ID)
		assert.Equal("test1", res.Name)
		assert.Equal("银行卡", res.Category)
		assert.Equal(1, res.Priority)
		assert.True(res.Updated.After(res.Created))
	})

	t.Run("Delete a entry", func(t *testing.T) {
		assert := assert.New(t)

		res, err := request.Delete(host+"/entries/"+entryID.String()).
			Set(gear.HeaderAuthorization, accessToken).End()
		assert.Nil(err)
		assert.Equal(204, res.StatusCode)

		res, err = request.Get(host+"/entries/"+entryID.String()).
			Set(gear.HeaderAuthorization, accessToken).End()
		assert.Nil(err)
		assert.Equal(404, res.StatusCode)
	})

	t.Run("Restore a entry", func(t *testing.T) {
		assert := assert.New(t)
		res := new(dao.EntrySum)

		_, err := request.Put(host+"/entries/"+entryID.String()+"/restore").
			Set(gear.HeaderAuthorization, accessToken).
			JSON(res)
		assert.Nil(err)

		assert.Equal(entryID, res.ID)
		assert.Equal("test1", res.Name)
		assert.Equal("银行卡", res.Category)
		assert.Equal(1, res.Priority)

		res = new(dao.EntrySum)
		_, err = request.Get(host+"/entries/"+entryID.String()).
			Set(gear.HeaderAuthorization, accessToken).
			JSON(res)
		assert.Nil(err)

		assert.Equal(entryID, res.ID)
		assert.Equal("test1", res.Name)
		assert.Equal("银行卡", res.Category)
		assert.Equal(1, res.Priority)
	})
}

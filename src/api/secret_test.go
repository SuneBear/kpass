package api_test

import (
	"fmt"
	"testing"

	"github.com/DavidCai1993/request"
	"github.com/seccom/kpass/src"
	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/util"
	"github.com/stretchr/testify/assert"
	"github.com/teambition/gear"
)

func TestSecretAPI(t *testing.T) {
	app := src.New("", "test")
	srv := app.Start()
	defer srv.Close()

	host := "http://" + srv.Addr().String()
	userInfo := NewUser(host)

	entry := new(schema.EntrySum)
	_, err := request.Post(fmt.Sprintf(`%s/api/teams/%s/entries`, host, userInfo.TeamID)).
		Set(gear.HeaderAuthorization, userInfo.AccessToken).
		Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
		Send(map[string]interface{}{"name": "test"}).
		JSON(entry)
	assert.Nil(t, err)

	var secretID util.OID
	t.Run("Add a secret", func(t *testing.T) {
		assert := assert.New(t)
		res := new(schema.SecretResult)

		_, err := request.Post(host+"/api/entries/"+entry.ID.String()+"/secrets").
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
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

		_, err := request.Put(host+"/api/entries/"+entry.ID.String()+"/secrets/"+secretID.String()).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
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

		res, err := request.Delete(host+"/api/entries/"+entry.ID.String()+"/secrets/"+secretID.String()).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).End()
		assert.Nil(err)
		assert.Equal(204, res.StatusCode)

		res2 := new(schema.EntryResult)
		_, err = request.Get(host+"/api/entries/"+entry.ID.String()).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			JSON(res)
		assert.Nil(err)
		assert.True(len(res2.Secrets) == 0)
	})
}

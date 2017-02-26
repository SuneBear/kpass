package api_test

import (
	"strings"
	"testing"

	"fmt"

	"github.com/DavidCai1993/request"
	"github.com/seccom/kpass/src"
	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/util"
	"github.com/stretchr/testify/assert"
	"github.com/teambition/gear"
)

func TestEntryAPI(t *testing.T) {
	app := src.New("", "test")
	srv := app.Start()
	defer srv.Close()

	host := "http://" + srv.Addr().String()
	userInfo := NewUser(host)

	t.Run("Find with no content", func(t *testing.T) {
		assert := assert.New(t)
		res := new([]*schema.EntrySum)

		_, err := request.Get(fmt.Sprintf(`%s/api/teams/%s/entries`, host, userInfo.TeamID)).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			JSON(res)
		assert.Nil(err)
		assert.True(len(*res) == 0)
	})

	var entryID util.OID
	t.Run("Create a entry", func(t *testing.T) {
		assert := assert.New(t)
		res := new(schema.EntrySum)

		_, err := request.Post(fmt.Sprintf(`%s/api/teams/%s/entries`, host, userInfo.TeamID)).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
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
		res := new(schema.EntryResult)

		_, err := request.Get(host+"/api/entries/"+entryID.String()).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			JSON(res)
		assert.Nil(err)

		assert.Equal(entryID, res.ID)
		assert.Equal("test", res.Name)
		assert.Equal("", res.Category)
		assert.Equal(0, res.Priority)
		assert.True(strings.Contains(res.String(), `"secrets":[]`))
		assert.True(strings.Contains(res.String(), `"shares":[]`))
		assert.True(strings.Contains(res.String(), `"files":[]`))
	})

	var secretID util.OID
	t.Run("Add a secret", func(t *testing.T) {
		assert := assert.New(t)
		res := new(schema.SecretResult)

		_, err := request.Post(host+"/api/entries/"+entryID.String()+"/secrets").
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

	t.Run("Find a entry again", func(t *testing.T) {
		assert := assert.New(t)
		res := new(schema.EntryResult)

		_, err := request.Get(host+"/api/entries/"+entryID.String()).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
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
		res := new([]*schema.EntrySum)

		_, err := request.Get(fmt.Sprintf(`%s/api/teams/%s/entries`, host, userInfo.TeamID)).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
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
		res := new(schema.EntrySum)

		_, err := request.Put(host+"/api/entries/"+entryID.String()).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
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

		res, err := request.Delete(host+"/api/entries/"+entryID.String()).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).End()
		assert.Nil(err)
		assert.Equal(204, res.StatusCode)

		res, err = request.Get(host+"/api/entries/"+entryID.String()).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).End()
		assert.Nil(err)
		assert.Equal(404, res.StatusCode)
	})

	t.Run("Undelete the entry", func(t *testing.T) {
		assert := assert.New(t)
		res := new(schema.EntrySum)

		_, err := request.Post(host+"/api/entries/"+entryID.String()+":undelete").
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			JSON(res)
		assert.Nil(err)

		assert.Equal(entryID, res.ID)
		assert.Equal("test1", res.Name)
		assert.Equal("银行卡", res.Category)
		assert.Equal(1, res.Priority)

		res = new(schema.EntrySum)
		_, err = request.Get(host+"/api/entries/"+entryID.String()).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			JSON(res)
		assert.Nil(err)

		assert.Equal(entryID, res.ID)
		assert.Equal("test1", res.Name)
		assert.Equal("银行卡", res.Category)
		assert.Equal(1, res.Priority)
	})
}

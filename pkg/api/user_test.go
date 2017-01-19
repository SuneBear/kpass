package api_test

import (
	"testing"

	"github.com/DavidCai1993/request"
	"github.com/seccom/kpass/pkg"
	"github.com/seccom/kpass/pkg/auth"
	"github.com/seccom/kpass/pkg/schema"
	"github.com/seccom/kpass/pkg/util"
	"github.com/stretchr/testify/assert"
	"github.com/teambition/gear"
)

func TestUserAPI(t *testing.T) {
	app := pkg.New("", "test")
	srv := app.Start()
	defer srv.Close()

	host := "http://" + srv.Addr().String()
	id := "admin"
	pass := util.SHA256Sum(util.SHA256Sum("password"))

	t.Run("Join", func(t *testing.T) {
		assert := assert.New(t)
		user := &schema.UserResult{}

		_, err := request.Post(host+"/join").
			Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
			Send(map[string]interface{}{"id": id, "pass": pass}).
			JSON(user)
		assert.Nil(err)

		assert.Equal(id, user.ID)
		assert.NotNil(user.Created)
		assert.Equal(user.Created, user.Updated)
	})

	t.Run("Login", func(t *testing.T) {
		assert := assert.New(t)
		res := &struct {
			ExpiresIn   int    `json:"expires_in"`
			TokenType   string `json:"token_type"`
			AccessToken string `json:"access_token"`
		}{}

		_, err := request.Post(host+"/login").
			Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
			Send(map[string]interface{}{"username": id, "password": pass, "grant_type": "password"}).
			JSON(res)
		assert.Nil(err)

		assert.Equal(600, res.ExpiresIn)
		assert.Equal("Bearer", res.TokenType)

		claims, _ := auth.JWT().Decode(res.AccessToken)
		assert.Equal("admin", claims.Get("id").(string))
		assert.True(len(claims.Get("key").(string)) > 0)
	})
}

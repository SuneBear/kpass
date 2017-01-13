package userAPI_test

import (
	"testing"

	"github.com/DavidCai1993/request"
	"github.com/seccom/kpass/app"
	"github.com/seccom/kpass/app/crypto"
	"github.com/seccom/kpass/app/dao"
	"github.com/stretchr/testify/assert"
	"github.com/teambition/gear"
)

func TestAPIUser(t *testing.T) {
	srv := app.New().Start()
	defer srv.Close()

	host := "http://" + srv.Addr().String()
	pass := crypto.SHA256Sum("password")

	t.Run("user Join", func(t *testing.T) {
		assert := assert.New(t)

		res, err := request.Post(host+"/join").
			Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
			Send(map[string]interface{}{"id": "admin", "pass": pass}).
			JSON(&dao.UserResult{})
		assert.Nil(err)

		assert.Equal("admin", res.(*dao.UserResult).ID)
		assert.NotNil(res.(*dao.UserResult).Created)
		assert.Equal(res.(*dao.UserResult).Created, res.(*dao.UserResult).Updated)
	})
}

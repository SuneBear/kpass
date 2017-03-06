package model

import (
	"testing"
	"time"

	"github.com/seccom/kpass/src/auth"
	"github.com/seccom/kpass/src/service"
	"github.com/stretchr/testify/assert"
)

func TestUserModel(t *testing.T) {
	db, _ := service.NewDB("")
	userModel := new(User).Init(db)
	auth.Init(db.Salt, 20*time.Minute)

	t.Run("CheckID", func(t *testing.T) {
		assert := assert.New(t)

		err := userModel.CheckID("ad")
		assert.Equal(`invalid user id "ad"`, err.Error())

		err = userModel.CheckID("admin")
		assert.Nil(err)
	})

	t.Run("Create And CheckLogin", func(t *testing.T) {
		assert := assert.New(t)

		user, err := userModel.CheckLogin("admin", "password")
		assert.Nil(user)
		assert.NotNil(err)

		user, err = userModel.Create("admin", "password")
		assert.Nil(err)
		assert.Equal("admin", user.ID)
		assert.True(len(user.Pass) > 0)

		err = userModel.CheckID("admin")
		assert.NotNil(err)

		user, err = userModel.CheckLogin("admin", "password")
		assert.Equal("admin", user.ID)
		assert.Nil(err)

		user, err = userModel.CheckLogin("admin", "password1")
		assert.Nil(user)
		assert.NotNil(err)
	})

	t.Run("Find", func(t *testing.T) {
		assert := assert.New(t)

		user, err := userModel.Find("admin")
		assert.Nil(err)
		assert.Equal("admin", user.ID)

		user, err = userModel.Find("admin1")
		assert.Nil(user)
		assert.NotNil(err)
	})

	t.Run("Update", func(t *testing.T) {
		assert := assert.New(t)

		user, err := userModel.Find("admin")
		assert.Nil(err)
		assert.False(user.IsBlocked)

		user.IsBlocked = true
		assert.Nil(userModel.Update(user))

		user, err = userModel.Find("admin")
		assert.Nil(err)
		assert.True(user.IsBlocked)
	})

	t.Run("FindUsers", func(t *testing.T) {
		assert := assert.New(t)

		users, err := userModel.FindUsers("admin")
		assert.Nil(err)
		assert.Equal("admin", users[0].ID)
	})
}

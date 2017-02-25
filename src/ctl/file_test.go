package ctl_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"testing"
	"time"

	"github.com/DavidCai1993/request"
	"github.com/seccom/kpass/src"
	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/util"
	"github.com/stretchr/testify/assert"
	"github.com/teambition/gear"
)

var count int = int(time.Now().Unix())

type UserInfo struct {
	ID, Pass, AccessToken, TeamID string
}

func NewUser(host string) *UserInfo {
	count++
	info := &UserInfo{}
	info.ID = "user" + strconv.Itoa(count)
	info.Pass = util.SHA256Sum(util.SHA256Sum(util.RandPass(8, 2, 2)))
	_, err := request.Post(host+"/api/join").
		Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
		Send(map[string]interface{}{"id": info.ID, "pass": info.Pass}).
		End()

	if err != nil {
		panic(err)
	}

	res, err := request.Post(host+"/api/login").
		Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
		Send(map[string]interface{}{"username": info.ID, "password": info.Pass, "grant_type": "password"}).
		JSON()
	if err != nil {
		panic(err)
	}

	info.AccessToken = "Bearer " + (*res.(*map[string]interface{}))["access_token"].(string)

	teams := &[]*schema.TeamResult{}
	_, err = request.Get(host+"/api/teams").
		Set(gear.HeaderAuthorization, info.AccessToken).
		Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
		JSON(teams)
	if err != nil {
		panic(err)
	}

	info.TeamID = (*teams)[0].ID.String()
	res, err = request.Post(fmt.Sprintf(`%s/api/teams/%s/token`, host, info.TeamID)).
		Set(gear.HeaderAuthorization, info.AccessToken).
		Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
		Send(map[string]interface{}{"password": info.Pass, "grant_type": "password"}).
		JSON()
	if err != nil {
		panic(err)
	}
	info.AccessToken = "Bearer " + (*res.(*map[string]interface{}))["access_token"].(string)
	return info
}

func TestFileController(t *testing.T) {
	app := src.New("", "test")
	srv := app.Start()
	defer srv.Close()

	host := "http://" + srv.Addr().String()
	userInfo := NewUser(host)

	t.Run("Upload user avatar", func(t *testing.T) {
		assert := assert.New(t)
		user := &schema.UserResult{}
		assert.False(user.Avatar.Valid())

		_, err := request.Post(host+"/upload/avatar").
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			Attach("upload", "../../web/public/logo.png", "my avatar.png").
			JSON(user)
		assert.Nil(err)
		assert.True(user.Avatar.Valid())

		res, err := request.Get(fmt.Sprintf(
			`%s/download/%s?refType=user&refID=%s`, host, user.Avatar, user.ID)).End()
		assert.Nil(err)
		assert.Equal("image/png", res.Header.Get(gear.HeaderContentType))
		assert.Equal("inline; filename=my avatar.png", res.Header.Get(gear.HeaderContentDisposition))

		avatar, err := ioutil.ReadFile("../../web/public/logo.png")
		assert.Nil(err)

		file, err := res.Content()
		assert.Nil(err)
		assert.Equal(0, bytes.Compare(avatar, file))
		assert.Equal(strconv.Itoa(len(avatar)), res.Header.Get(gear.HeaderContentLength))
	})

	t.Run("Upload team logo", func(t *testing.T) {
		assert := assert.New(t)
		team := &schema.TeamResult{}
		_, err := request.Post(fmt.Sprintf(`%s/upload/team/%s/logo`, host, userInfo.TeamID)).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			Attach("upload", "../../web/public/logo.png", "logo.png").
			JSON(team)
		assert.Nil(err)
		assert.True(team.Logo.Valid())

		res, err := request.Get(fmt.Sprintf(
			`%s/download/%s?refType=team&refID=%s`, host, team.Logo, team.ID)).End()
		assert.Nil(err)
		assert.Equal("image/png", res.Header.Get(gear.HeaderContentType))
		assert.Equal("inline; filename=logo.png", res.Header.Get(gear.HeaderContentDisposition))

		avatar, err := ioutil.ReadFile("../../web/public/logo.png")
		assert.Nil(err)

		file, err := res.Content()
		assert.Nil(err)
		assert.Equal(0, bytes.Compare(avatar, file))
		assert.Equal(strconv.Itoa(len(avatar)), res.Header.Get(gear.HeaderContentLength))
	})

	t.Run("Upload file to the entry", func(t *testing.T) {
		assert := assert.New(t)
		res := &schema.EntrySum{}

		_, err := request.Post(fmt.Sprintf(`%s/api/teams/%s/entries`, host, userInfo.TeamID)).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
			Send(map[string]interface{}{"name": "test"}).
			JSON(res)
		assert.Nil(err)

		EntryID := res.ID
		res2 := &schema.EntryResult{}
		_, err = request.Get(fmt.Sprintf(`%s/api/entries/%s`, host, EntryID)).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			JSON(res2)
		assert.Nil(err)
		assert.Equal(EntryID, res2.ID)
		assert.Equal(0, len(res2.Files))

		file := &schema.FileResult{}
		_, err = request.Post(fmt.Sprintf(`%s/upload/entry/%s/file`, host, EntryID)).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			Attach("upload", "../../web/public/humans.txt", "humans.txt").
			JSON(file)
		assert.Nil(err)

		res3, err := request.Get(fmt.Sprintf(
			`%s/download/%s?refType=entry&refID=%s&signed=%s`, host, file.ID, EntryID, file.Signed)).End()
		assert.Nil(err)
		assert.Equal("text/plain; charset=utf-8", res3.Header.Get(gear.HeaderContentType))
		assert.Equal("attachment; filename=humans.txt", res3.Header.Get(gear.HeaderContentDisposition))

		txt, err := ioutil.ReadFile("../../web/public/humans.txt")
		assert.Nil(err)

		txt2, err := res3.Content()
		assert.Nil(err)
		assert.Equal(0, bytes.Compare(txt, txt2))
		assert.Equal(strconv.Itoa(len(txt)), res3.Header.Get(gear.HeaderContentLength))

		res4 := &schema.EntryResult{}
		_, err = request.Get(fmt.Sprintf(`%s/api/entries/%s`, host, EntryID)).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			JSON(res4)
		assert.Nil(err)
		assert.Equal(EntryID, res4.ID)
		assert.Equal(1, len(res4.Files))

		file = res4.Files[0]
		res5, err := request.Get(fmt.Sprintf(
			`%s/download/%s?refType=entry&refID=%s&signed=%s`, host, file.ID, EntryID, file.Signed)).End()
		assert.Nil(err)
		assert.Equal("text/plain; charset=utf-8", res5.Header.Get(gear.HeaderContentType))
		assert.Equal("attachment; filename=humans.txt", res5.Header.Get(gear.HeaderContentDisposition))

		txt3, err := res5.Content()
		assert.Nil(err)
		assert.Equal(0, bytes.Compare(txt, txt3))
		assert.Equal(strconv.Itoa(len(txt)), res5.Header.Get(gear.HeaderContentLength))

		res6, err := request.Delete(fmt.Sprintf(
			`%s/api/entries/%s/files/%s`, host, EntryID, file.ID)).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).End()
		assert.Nil(err)
		assert.Equal(204, res6.StatusCode)

		res7 := &schema.EntryResult{}
		_, err = request.Get(fmt.Sprintf(`%s/api/entries/%s`, host, EntryID)).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			JSON(res7)
		assert.Nil(err)
		assert.Equal(0, len(res7.Files))

		res8, err := request.Get(fmt.Sprintf(
			`%s/download/%s?refType=entry&refID=%s&signed=%s`, host, file.ID, EntryID, file.Signed)).End()
		assert.Nil(err)
		assert.Equal(404, res8.StatusCode)
	})
}

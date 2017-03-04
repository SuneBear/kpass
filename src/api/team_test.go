package api_test

import (
	"fmt"
	"testing"

	"github.com/DavidCai1993/request"
	"github.com/seccom/kpass/src"
	"github.com/seccom/kpass/src/schema"
	"github.com/stretchr/testify/assert"
	"github.com/teambition/gear"
)

func TestTeamAPI(t *testing.T) {
	app := src.New("", "test")
	srv := app.Start()
	defer srv.Close()

	host := "http://" + srv.Addr().String()
	userInfo := NewUser(host)
	userInfo2 := NewUser(host)

	team := new(schema.TeamResult)
	t.Run("create a team", func(t *testing.T) {
		assert := assert.New(t)

		_, err := request.Post(host+"/api/teams").
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
			Send(map[string]interface{}{"name": "KPass"}).
			JSON(team)

		assert.Nil(err)
		assert.Equal("KPass", team.Name)
		assert.Equal(userInfo.ID, team.UserID)
		assert.Equal("member", team.Visibility)
		assert.False(team.IsFrozen)
		assert.True(len(team.Members) == 1)
		assert.Equal(team.UserID, team.Members[0].ID)

		entry := new(schema.EntrySum)
		_, err = request.Post(fmt.Sprintf(`%s/api/teams/%s/entries`, host, team.ID)).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
			Send(map[string]interface{}{"name": "test"}).
			JSON(entry)
		assert.Nil(err)

		_, err = request.Post(host+"/api/entries/"+entry.ID.String()+"/secrets").
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
			Send(map[string]interface{}{"name": "test secret", "url": "test.com", "password": "123456"}).
			JSON()
		assert.Nil(err)
	})

	t.Run("invite a member and join", func(t *testing.T) {
		assert := assert.New(t)

		invite := new(struct {
			Code string `json:"code"`
		})
		_, err := request.Post(fmt.Sprintf(`%s/api/teams/%s/invite`, host, team.ID)).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
			Send(map[string]interface{}{"userID": userInfo2.ID}).
			JSON(invite)
		assert.Nil(err)
		assert.True(invite.Code != "")

		teamInfo := new(schema.TeamResult)
		_, err = request.Post(fmt.Sprintf(`%s/api/teams/join`, host)).
			Set(gear.HeaderAuthorization, userInfo2.AccessToken).
			Set(gear.HeaderContentType, gear.MIMEApplicationJSON).
			Send(map[string]interface{}{"code": invite.Code}).
			JSON(teamInfo)
		assert.Nil(err)
		assert.True(len(teamInfo.Members) == 2)
		assert.Equal(userInfo2.ID, teamInfo.Members[1].ID)

		entries := new([]*schema.EntrySum)
		_, err = request.Get(fmt.Sprintf(`%s/api/teams/%s/entries`, host, teamInfo.ID)).
			Set(gear.HeaderAuthorization, userInfo2.AccessToken).
			JSON(entries)
		assert.Nil(err)
		assert.Equal("test", (*entries)[0].Name)

		entry := new(schema.EntryResult)
		_, err = request.Get(host+"/api/entries/"+(*entries)[0].ID.String()).
			Set(gear.HeaderAuthorization, userInfo.AccessToken).
			JSON(entry)
		assert.Nil(err)
		assert.Equal("test", entry.Name)
		secret := entry.Secrets[0]
		assert.Equal("test secret", secret.Name)
		assert.Equal("test.com", secret.URL)
		assert.Equal("123456", secret.Pass)
	})
}

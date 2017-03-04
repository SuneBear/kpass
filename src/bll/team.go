package bll

import (
	"time"

	"github.com/seccom/kpass/src/auth"
	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/util"
	"github.com/teambition/gear"
)

// Team is Business Logic Layer for team
type Team struct {
	*Bll
}

// Create ...
func (b *Team) Create(userID, name, key, visibility string) (*schema.TeamResult, error) {
	teamPass := util.SHA256Sum(util.RandPass(20, 3, 5))
	res, err := b.Models.Team.Create(userID, teamPass, &schema.Team{
		Name:       name,
		UserID:     userID,
		Visibility: visibility,
		Members:    []string{userID},
	})

	if err != nil {
		return nil, err
	}

	if err = b.Models.Team.SavePass(res.ID, userID, key, teamPass); err != nil {
		return nil, err
	}
	return res, nil
}

// Invite ...
func (b *Team) Invite(ownerID, key, userID string, TeamID util.OID) (string, error) {
	team, err := b.Models.Team.Find(TeamID, false)
	if err != nil {
		return "", err
	}
	if team.UserID != ownerID {
		return "", &gear.Error{Code: 401, Msg: "not team owner"}
	}
	if team.HasMember(userID) {
		return "", &gear.Error{Code: 409, Msg: "already team member"}
	}
	if team.Visibility == "private" {
		return "", &gear.Error{Code: 403, Msg: "private team"}
	}
	owner, err := b.Models.User.Find(ownerID)
	if err != nil {
		return "", err
	}
	user, err := b.Models.User.Find(userID)
	if err != nil {
		return "", err
	}
	teamPass, err := b.Models.Team.GetPass(TeamID, ownerID, key)
	if err != nil {
		return "", err
	}
	teamPass, err = auth.EncryptText(auth.AESKey(owner.Pass, user.Pass), teamPass)
	if err != nil {
		return "", err
	}
	return auth.Sign(map[string]interface{}{"id": TeamID.String(), "code": teamPass}, time.Minute*20)
}

// Join ...
func (b *Team) Join(userID, key, token string) (*schema.TeamResult, error) {
	claims, err := auth.Verify(token)
	if err != nil {
		return nil, err
	}
	teamPass := claims.Get("code").(string)
	TeamID, err := util.ParseOID(claims.Get("id").(string))
	if err != nil {
		return nil, err
	}
	user, err := b.Models.User.Find(userID)
	if err != nil {
		return nil, err
	}
	team, err := b.Models.Team.Find(TeamID, false)
	if err != nil {
		return nil, err
	}
	if team.HasMember(userID) {
		return nil, &gear.Error{Code: 409, Msg: "already team member"}
	}
	owner, err := b.Models.User.Find(team.UserID)
	if err != nil {
		return nil, err
	}
	teamPass, err = auth.DecryptText(auth.AESKey(owner.Pass, user.Pass), teamPass)
	if err != nil {
		return nil, err
	}

	res, err := b.Models.Team.AddMember(team.UserID, userID, TeamID)
	if err != nil {
		return nil, err
	}
	err = b.Models.Team.SavePass(TeamID, userID, key, teamPass)
	return res, err
}

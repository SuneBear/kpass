package bll

import (
	"github.com/seccom/kpass/src/schema"
	"github.com/seccom/kpass/src/util"
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

	if err = b.Models.File.SaveTeamPass(res.ID, userID, key, teamPass); err != nil {
		return nil, err
	}
	return res, nil
}

package bll

import (
	"github.com/seccom/kpass/src/model"
	"github.com/seccom/kpass/src/service"
)

// Bll is Business Logic Layer with all models
type Bll struct {
	Models *model.All
}

// All ...
type All struct {
	Models *model.All
	Team   *Team
	Entry  *Entry
	Secret *Secret
}

// Init ...
func (a *All) Init(db *service.DB) *All {
	a.Models = new(model.All).Init(db)
	b := &Bll{a.Models}
	a.Team = &Team{b}
	a.Entry = &Entry{b}
	a.Secret = &Secret{b}
	return a
}

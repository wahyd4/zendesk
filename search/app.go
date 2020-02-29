package search

import (
	"github.com/wahyd4/zendesk/model"
)

// APP represents the application
type APP struct {
	organisations map[string]*model.Organisation
	users         map[string]*model.User
	tickets       map[string]*model.Ticket
}

// FindOrganisation find a organisation by organisation ID
func (app *APP) FindOrganisation(id string) *model.Organisation {
	return app.organisations[id]
}

// FindUser find a user by user ID
func (app *APP) FindUser(id string) *model.User {
	return app.users[id]
}

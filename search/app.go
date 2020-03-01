package search

import (
	"io/ioutil"

	"github.com/wahyd4/zendesk/index"
	"github.com/wahyd4/zendesk/model"
)

const (
	OrganisationsKey = "organisations"
	UsersKey         = "users"
	TicketsKey       = "tickets"
)

// APP represents the application
type APP struct {
	jsonContents  map[string][]byte
	organisations map[string]*model.Organisation
	users         map[string]*model.User
	tickets       map[string]*model.Ticket

	indexes       map[string]index.SearchIndex
	searchContext *SearchContext
}

// InitAPP takes data file paths and then initialise the application
func InitAPP(organisationsFile, usersFile, ticketsFile string) *APP {
	files := map[string]string{
		OrganisationsKey: organisationsFile,
		UsersKey:         usersFile,
		TicketsKey:       ticketsFile,
	}
	jsonContents := make(map[string][]byte)
	for fileType, file := range files {
		bytes, err := ioutil.ReadFile(file)
		if err != nil {
			panic("cannot init application due to unable to load:" + file)
		}
		jsonContents[fileType] = bytes
	}

	return &APP{
		jsonContents:  jsonContents,
		indexes:       make(map[string]index.SearchIndex),
		searchContext: &SearchContext{},
	}
}

// FindOrganisation find a organisation by organisation ID
func (app *APP) FindOrganisation(id string) *model.Organisation {
	return app.organisations[id]
}

// FindUser find a user by user ID
func (app *APP) FindUser(id string) *model.User {
	return app.users[id]
}

// FindTicket find a ticket by ticket ID
func (app *APP) FindTicket(id string) *model.Ticket {
	return app.tickets[id]
}

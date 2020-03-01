package search

import (
	"io/ioutil"

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

	organisationIndex map[string]map[string][]*model.Organisation
	// userIndex         map[string]map[string][]*model.User
	// ticketIndex       map[string]map[string][]*model.Ticket
}

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
		jsonContents:      jsonContents,
		organisationIndex: make(map[string]map[string][]*model.Organisation),
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

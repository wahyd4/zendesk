package search

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/wahyd4/zendesk/model"
	"github.com/wahyd4/zendesk/view"
)

// type dataHandler func(jsonContent []byte) error

func (app *APP) Parse(jsonFiles []string) error {
	// dataHandlers := map[string]dataHandler{
	// 	"organisations": app.LoadOrganisationsFromJSON,
	// 	"users":         app.LoadUsersFromJSON,
	// 	"tickets":       app.LoadTicketsFromJSON,
	// }

	for _, file := range jsonFiles {
		_, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read %s data file with error: %s", file, err.Error())
		}

	}
	return nil
}

func (app *APP) LoadOrganisationsFromJSON(jsonContent []byte) error {
	var organisations []*view.Organisation
	err := json.Unmarshal(jsonContent, &organisations)
	if err != nil {
		return fmt.Errorf("failed to load organisations from json: %s", err.Error())
	}
	app.organisations = convertOrganisationViews(organisations)
	return nil
}

func (app *APP) LoadUsersFromJSON(jsonContent []byte) error {
	var users []*view.User
	err := json.Unmarshal(jsonContent, &users)
	if err != nil {
		return fmt.Errorf("failed to load users from json: %s", err.Error())
	}
	app.users = app.convertUserViews(users)
	return nil
}

func (app *APP) LoadTicketsFromJSON(jsonContent []byte) error {
	var tickets []*view.Ticket
	err := json.Unmarshal(jsonContent, &tickets)
	if err != nil {
		return fmt.Errorf("failed to load tickets from json: %s", err.Error())
	}
	app.tickets = app.convertTicketViews(tickets)
	return nil
}

func convertOrganisationViews(views []*view.Organisation) map[string]*model.Organisation {
	models := make(map[string]*model.Organisation)
	for _, view := range views {
		models[stringID(view.ID)] = convertOrganisationView(view)
	}
	return models
}

func (app *APP) convertUserViews(views []*view.User) map[string]*model.User {
	models := make(map[string]*model.User)
	for _, view := range views {
		models[stringID(view.ID)] = app.convertUserView(view)
	}
	return models
}

func (app *APP) convertTicketViews(views []*view.Ticket) map[string]*model.Ticket {
	models := make(map[string]*model.Ticket)
	for _, view := range views {
		models[view.ID] = app.convertTicketView(view)
	}
	return models
}

func convertOrganisationView(view *view.Organisation) *model.Organisation {
	return &model.Organisation{
		ID:            view.ID,
		URL:           view.URL,
		ExternalID:    view.ExternalID,
		Name:          view.Name,
		DomainNames:   view.DomainNames,
		CreatedAt:     view.CreatedAt,
		Details:       view.Details,
		SharedTickets: view.SharedTickets,
		Tags:          view.Tags,
	}
}

func (app *APP) convertUserView(view *view.User) *model.User {
	return &model.User{
		ID:           view.ID,
		URL:          view.URL,
		ExternalID:   view.ExternalID,
		Name:         view.Name,
		Alias:        view.Alias,
		CreatedAt:    view.CreatedAt,
		Active:       view.Active,
		Verified:     view.Verified,
		Shared:       view.Shared,
		Locale:       view.Locale,
		Timezone:     view.Timezone,
		LastLoginAt:  view.LastLoginAt,
		Email:        view.Email,
		Phone:        view.Phone,
		Signature:    view.Signature,
		Organization: app.FindOrganisation(stringID(view.OrganizationID)),
		Tags:         view.Tags,
		Suspended:    view.Suspended,
		Role:         view.Role,
	}
}

func (app *APP) convertTicketView(view *view.Ticket) *model.Ticket {
	return &model.Ticket{
		ID:           view.ID,
		URL:          view.URL,
		ExternalID:   view.ExternalID,
		CreatedAt:    view.CreatedAt,
		Type:         view.Type,
		Subject:      view.Subject,
		Description:  view.Description,
		Priority:     view.Priority,
		Status:       view.Status,
		Submitter:    app.FindUser(stringID(view.SubmitterID)),
		Assignee:     app.FindUser(stringID(view.AssigneeID)),
		Organization: app.FindOrganisation(stringID(view.OrganizationID)),
		Tags:         view.Tags,
		HasIncidents: view.HasIncidents,
		DueAt:        view.DueAt,
		Via:          view.Via,
	}
}

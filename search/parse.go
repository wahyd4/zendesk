package search

import (
	"encoding/json"
	"fmt"

	"github.com/wahyd4/zendesk/model"
	"github.com/wahyd4/zendesk/view"
)

func (app *APP) Parse(jsonFiles []string) error {
	return nil
}

func (app *APP) LoadOrganisationsFromJSON(jsonContent []byte) error {
	var organisations []*view.Organisation
	err := json.Unmarshal(jsonContent, &organisations)
	if err != nil {
		return fmt.Errorf("failed to load organisations from json %s", err.Error())
	}
	app.organisations = convertOrganisationViews(organisations)
	return nil
}

func (app *APP) LoadUsersFromJSON(jsonContent []byte) error {
	var users []*view.User
	err := json.Unmarshal(jsonContent, &users)
	if err != nil {
		return fmt.Errorf("failed to load users from json %s", err.Error())
	}
	app.users = app.convertUserViews(users)
	return nil
}

func convertOrganisationViews(views []*view.Organisation) map[string]*model.Organisation {
	models := make(map[string]*model.Organisation)
	for _, view := range views {
		models[fmt.Sprintf("%d", view.ID)] = convertOrganisationView(view)
	}
	return models
}

func (app *APP) convertUserViews(views []*view.User) map[string]*model.User {
	models := make(map[string]*model.User)
	for _, view := range views {
		models[fmt.Sprintf("%d", view.ID)] = app.convertUserView(view)
	}
	return models
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
		Organization: app.FindOrganisation(fmt.Sprintf("%d", view.OrganizationID)),
		Tags:         view.Tags,
		Suspended:    view.Suspended,
		Role:         view.Role,
	}
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

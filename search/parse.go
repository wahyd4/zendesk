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

func convertOrganisationViews(views []*view.Organisation) map[string]*model.Organisation {
	models := make(map[string]*model.Organisation)
	for _, view := range views {
		models[fmt.Sprintf("%d", view.ID)] = convertOrganisationView(view)
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

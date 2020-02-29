package searchable

import (
	"encoding/json"
	"fmt"

	"github.com/wahyd4/zendesk/model"
	"github.com/wahyd4/zendesk/view"
)

type OrganisationSearchable struct {
}

func (org OrganisationSearchable) LoadFromJSON(jsonContent []byte) (interface{}, error) {
	var organisations []*view.Organisation
	err := json.Unmarshal(jsonContent, &organisations)
	if err != nil {
		return nil, fmt.Errorf("failed to load organisations from json %s", err.Error())
	}

	return org.convertViews(organisations), nil
}

func (org OrganisationSearchable) convertViews(views []*view.Organisation) []*model.Organisation {
	models := make([]*model.Organisation, 0)
	for _, view := range views {
		models = append(models, org.convertView(view))
	}
	return models
}

func (org OrganisationSearchable) convertView(view *view.Organisation) *model.Organisation {
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

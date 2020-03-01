package index

import "github.com/wahyd4/zendesk/model"

type OrganisationIndex struct {
	data map[string]map[string][]*model.Organisation
}

func (index OrganisationIndex) ListSearchableFields() []string {
	var result []string
	for field := range index.data {
		result = append(result, field)
	}
	return result
}

func (index OrganisationIndex) Search(field, value string) []string {
	organisations := index.data[field][value]

	var result []string
	for _, organisation := range organisations {
		result = append(result, organisation.Print())
	}
	return result
}

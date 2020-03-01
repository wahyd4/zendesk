package search

import (
	"encoding/json"
	"fmt"

	"github.com/wahyd4/zendesk/model"
)

const (
	idKey = "_id"
)

var (
	arrayKeys   = []string{"domain_names", "tags"}
	booleanKeys = []string{"shared_tickets"}
	integerKeys = []string{"_id"}
)

func (app *APP) BuildOrganisationIndex() error {
	bytes := app.jsonContents[OrganisationsKey]

	var organisations []map[string]interface{}

	err := json.Unmarshal(bytes, &organisations)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json to organisations array: %w", err)
	}

	if len(organisations) == 0 {
		return nil
	}
	searchIndex := app.organisationIndex
	organisationsTemplate := organisations[0]

	for fieldKey := range organisationsTemplate {

		for _, organisation := range organisations {
			organisationID := stringifyID(int(organisation[idKey].(float64)))

			if searchIndex[fieldKey] == nil {
				searchIndex[fieldKey] = make(map[string][]*model.Organisation)
			}

			fieldIndex := searchIndex[fieldKey]

			fieldValues := make([]string, 0)

			// process value based on different types
			if OneOfTheKeys(integerKeys, fieldKey) {
				fieldValues = append(fieldValues, stringifyID(int(organisation[fieldKey].(float64))))
			} else if OneOfTheKeys(booleanKeys, fieldKey) {
				fieldValues = append(fieldValues, fmt.Sprintf("%t", organisation[fieldKey].(bool)))
			} else if OneOfTheKeys(arrayKeys, fieldKey) {
				fieldValues = toStringSlice(organisation[fieldKey].([]interface{}))
			} else {
				fieldValues = append(fieldValues, organisation[fieldKey].(string))
			}

			// build fieldValue based indeX
			for _, fieldValue := range fieldValues {
				if fieldIndex[fieldValue] == nil {
					sameValueList := make([]*model.Organisation, 0)
					fieldIndex[fieldValue] = sameValueList
				}

				fieldIndex[fieldValue] = append(fieldIndex[fieldValue], app.FindOrganisation(organisationID))
			}

		}
	}
	return nil
}

func OneOfTheKeys(keys []string, target string) bool {
	for _, key := range keys {
		if key == target {
			return true
		}
	}
	return false
}

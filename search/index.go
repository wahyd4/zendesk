package search

import (
	"encoding/json"
	"fmt"

	"github.com/wahyd4/zendesk/model"
)

const (
	idKey = "_id"
)

type IndexBuilderFunc func() error

var (
	arrayKeys   = []string{"domain_names", "tags"}
	booleanKeys = []string{"shared_tickets", "active", "verified", "shared", "suspended", "has_incidents"}
	integerKeys = []string{"_id", "organization_id", "submitter_id", "assignee_id"}
)

// BuildIndexes build indexes for searching
func (app *APP) BuildIndexes() error {
	indexBuilders := map[string]IndexBuilderFunc{
		OrganisationsKey: app.buildOrganisationIndex,
		UsersKey:         app.buildUserIndex,
		TicketsKey:       app.buildTicketIndex,
	}

	for indexType, indexBuilderFunc := range indexBuilders {
		if err := indexBuilderFunc(); err != nil {
			return fmt.Errorf("failed to build index for %s with error: %w", indexType, err)
		}
	}

	return nil
}

func (app *APP) buildOrganisationIndex() error {
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

	// get first entity from json array as the template
	organisationsTemplate := organisations[0]

	for fieldKey := range organisationsTemplate {

		for _, organisation := range organisations {
			organisationID := stringifyID(int(organisation[idKey].(float64)))

			if searchIndex[fieldKey] == nil {
				searchIndex[fieldKey] = make(map[string][]*model.Organisation)
			}
			fieldIndex := searchIndex[fieldKey]

			fieldValues := extractFieldValues(organisation, fieldKey)

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

func (app *APP) buildUserIndex() error {
	bytes := app.jsonContents[UsersKey]

	var users []map[string]interface{}

	err := json.Unmarshal(bytes, &users)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json to user array: %w", err)
	}

	if len(users) == 0 {
		return nil
	}
	searchIndex := app.userIndex

	// get first entity from json array as the template
	entityTemplate := users[0]

	for fieldKey := range entityTemplate {

		for _, user := range users {
			userID := stringifyID(int(user[idKey].(float64)))

			if searchIndex[fieldKey] == nil {
				searchIndex[fieldKey] = make(map[string][]*model.User)
			}
			fieldIndex := searchIndex[fieldKey]

			fieldValues := extractFieldValues(user, fieldKey)

			// build fieldValue based indeX
			for _, fieldValue := range fieldValues {
				if fieldIndex[fieldValue] == nil {
					sameValueList := make([]*model.User, 0)
					fieldIndex[fieldValue] = sameValueList
				}

				fieldIndex[fieldValue] = append(fieldIndex[fieldValue], app.FindUser(userID))
			}

		}
	}
	return nil
}

func (app *APP) buildTicketIndex() error {
	bytes := app.jsonContents[TicketsKey]

	var tickets []map[string]interface{}

	err := json.Unmarshal(bytes, &tickets)
	if err != nil {
		return fmt.Errorf("failed to unmarshal json to ticket array: %w", err)
	}

	if len(tickets) == 0 {
		return nil
	}
	searchIndex := app.ticketIndex

	// get first entity from json array as the template
	entityTemplate := tickets[0]

	for fieldKey := range entityTemplate {

		for _, ticket := range tickets {
			ticketID := ticket[idKey].(string)

			if searchIndex[fieldKey] == nil {
				searchIndex[fieldKey] = make(map[string][]*model.Ticket)
			}
			fieldIndex := searchIndex[fieldKey]
			var fieldValues []string

			// due to the ticket id is string
			if fieldKey == idKey {
				fieldValues = []string{ticketID}
			} else {
				fieldValues = extractFieldValues(ticket, fieldKey)
			}

			// build fieldValue based indeX
			for _, fieldValue := range fieldValues {
				if fieldIndex[fieldValue] == nil {
					sameValueList := make([]*model.Ticket, 0)
					fieldIndex[fieldValue] = sameValueList
				}

				fieldIndex[fieldValue] = append(fieldIndex[fieldValue], app.FindTicket(ticketID))
			}

		}
	}
	return nil
}

func extractFieldValues(entity map[string]interface{}, fieldKey string) []string {
	fieldValues := make([]string, 0)

	if entity[fieldKey] == nil {
		return fieldValues
	}
	// process value based on different types
	if oneOfTheKeys(integerKeys, fieldKey) {
		fieldValues = append(fieldValues, stringifyID(int(entity[fieldKey].(float64))))
	} else if oneOfTheKeys(booleanKeys, fieldKey) {
		fieldValues = append(fieldValues, fmt.Sprintf("%t", entity[fieldKey].(bool)))
	} else if oneOfTheKeys(arrayKeys, fieldKey) {
		fieldValues = toStringSlice(entity[fieldKey].([]interface{}))
	} else {
		fieldValues = append(fieldValues, entity[fieldKey].(string))
	}

	return fieldValues
}

func oneOfTheKeys(keys []string, target string) bool {
	for _, key := range keys {
		if key == target {
			return true
		}
	}
	return false
}

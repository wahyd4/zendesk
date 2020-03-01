package index

import "github.com/wahyd4/zendesk/model"

type TicketIndex struct {
	data map[string]map[string][]*model.Ticket
}

func (index TicketIndex) ListSearchableFields() []string {
	var result []string
	for field := range index.data {
		result = append(result, field)
	}
	return result
}

func (index TicketIndex) Search(field, value string) []string {
	tickets := index.data[field][value]

	var result []string
	for _, ticket := range tickets {
		result = append(result, ticket.Print())
	}
	return result
}

package index

import (
	"github.com/wahyd4/zendesk/model"
)

type TicketIndex struct {
	Data map[string]map[string][]*model.Ticket
}

func (index TicketIndex) ListSearchableFields() []string {
	var result []string
	for field := range index.Data {
		result = append(result, field)
	}
	return result
}

func (index TicketIndex) Search(field, value string) interface{} {
	if len(index.Data[field][value]) == 0 {
		return nil
	}
	return index.Data[field][value]
}

package index

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/wahyd4/zendesk/model"
)

type UserIndex struct {
	Data map[string]map[string][]*model.User
}

func (index UserIndex) ListSearchableFields() []string {
	var result []string
	for field := range index.Data {
		result = append(result, field)
	}
	return result
}

func (index UserIndex) Search(field, value string) {
	users := index.Data[field][value]

	spew.Dump(users)
}

package index

import "github.com/wahyd4/zendesk/model"

type UserIndex struct {
	data map[string]map[string][]*model.User
}

func (index UserIndex) ListSearchableFields() []string {
	var result []string
	for field := range index.data {
		result = append(result, field)
	}
	return result
}

func (index UserIndex) Search(field, value string) []string {
	users := index.data[field][value]

	var result []string
	for _, user := range users {
		result = append(result, user.Print())
	}
	return result
}

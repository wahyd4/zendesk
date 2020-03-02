package search

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

// SearchContext holds the information of current search context
type SearchContext struct {
	dataType  string
	fieldName string
}

// UpdateDataType allows user select data type
func (app *APP) UpdateDataType(dataType string) {
	app.searchContext = &SearchContext{
		dataType: dataType,
	}
}

// UpdateFieldName updates search field name of a data type
func (app *APP) UpdateFieldName(fieldName string) {
	app.searchContext.fieldName = fieldName
}

// ListSearchableFields list all the searchable fields
func (app *APP) ListSearchableFields() []string {
	dataType := app.searchContext.dataType

	return app.indexes[dataType].ListSearchableFields()
}

// ListSearchableFields list all the searchable fields
func (app *APP) Search(value string) {
	dataType := app.searchContext.dataType
	field := app.searchContext.fieldName

	searchResult := app.indexes[dataType].Search(field, value)

	if searchResult == nil {
		fmt.Printf("Cannot find any matched result from %s with %s:%s \n\n", dataType, field, value)
		return
	}

	spew.Dump(searchResult)
}

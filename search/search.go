package search

// SearchContext holds the information of current search context
type SearchContext struct {
    dataType string
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
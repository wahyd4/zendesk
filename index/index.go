package index

type SearchIndex interface {
	ListSearchableFields() []string
	Search(field, value string)
}

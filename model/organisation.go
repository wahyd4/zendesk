package model

// Organisation represents a Organisation model
type Organisation struct {
	ID            int
	URL           string
	ExternalID    string
	Name          string
	DomainNames   []string
	CreatedAt     string
	Details       string
	SharedTickets bool
	Tags          []string
}

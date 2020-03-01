package model

import "fmt"

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

func (org *Organisation) Print() string {
	return fmt.Sprintf(
		`ID            %d
		URL           string
		ExternalID    string
		Name          string
		DomainNames   []string
		CreatedAt     string
		Details       string
		SharedTickets bool
		Tags          []string`, org.ID)
}

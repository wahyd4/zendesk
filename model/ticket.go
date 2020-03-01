package model

import "fmt"

// Ticket represents a ticket model
type Ticket struct {
	ID           string
	URL          string
	ExternalID   string
	CreatedAt    string
	Type         string
	Subject      string
	Description  string
	Priority     string
	Status       string
	Submitter    *User
	Assignee     *User
	Organization *Organisation
	Tags         []string
	HasIncidents bool
	DueAt        string
	Via          string
}

func (ticket *Ticket) Print() string {
	return fmt.Sprintf(
		`ID           %s
		URL          string
		ExternalID   string
		CreatedAt    string
		Type         string
		Subject      string
		Description  string
		Priority     string
		Status       string
		Submitter    *User
		Assignee     *User
		Organization *Organisation
		Tags         []string
		HasIncidents bool
		DueAt        string
		Via          string`, ticket.ID)
}

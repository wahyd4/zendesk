package model

// Ticket represents a ticket model
type Ticket struct {
	ID             string
	URL            string
	ExternalID     string
	CreatedAt      string
	Type           string
	Subject        string
	Description    string
	Priority       string
	Status         string
	Submitter      *User
	AssigneeID     *User
	OrganizationID *Organisation
	Tags           []string
	HasIncidents   bool
	DueAt          string
	Via            string
}

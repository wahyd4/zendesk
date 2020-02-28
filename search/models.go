package search

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

// User represents a user model
type User struct {
	ID           int
	URL          string
	ExternalID   string
	Name         string
	Alias        string
	CreatedAt    string
	Active       bool
	Verified     bool
	Shared       bool
	Locale       string
	Timezone     string
	LastLoginAt  string
	Email        string
	Phone        string
	Signature    string
	Organization *Organisation
	Tags         []string
	Suspended    bool
	Role         string
}

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

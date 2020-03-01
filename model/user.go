package model

import "fmt"

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

func (user *User) Print() string {
	return fmt.Sprintf(
		`ID           %d
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
		Role         string`, user.ID)
}

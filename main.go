package main

import (
	"fmt"

	"github.com/wahyd4/zendesk/search"
)

func main() {

	app := search.InitAPP("data/organizations.json", "data/users.json", "data/tickets.json")
	if err := app.Parse(); err != nil {
		panic("cannot load data: " + err.Error())
	}
	if err := app.BuildOrganisationIndex(); err != nil {
		panic("failed to build indexes: " + err.Error())
	}
	fmt.Println("---------------")
}

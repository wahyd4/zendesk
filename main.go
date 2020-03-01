package main

import (
	"github.com/wahyd4/zendesk/search"
)

func main() {

	app := search.InitAPP("data/organizations.json", "data/users.json", "data/tickets.json")
	if err := app.Parse(); err != nil {
		panic("cannot load data: " + err.Error())
	}
	if err := app.BuildIndexes(); err != nil {
		panic("failed to build indexes: " + err.Error())
	}

}

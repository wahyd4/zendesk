package main

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
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
	for {

		prompt := promptui.Select{
			Label: "Select a type of data to search with",
			Items: []string{search.OrganisationsKey, search.UsersKey, search.TicketsKey},
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Failed to process your input %v\n", err)
			os.Exit(1)
		}
		app.UpdateDataType(result)

		prompt = promptui.Select{
			Label: "Please select a field to search with",
			Items: app.ListSearchableFields(),
			Size:  20,
		}

		_, fieldName, err := prompt.Run()

		if err != nil {
			fmt.Printf("Failed to process your input %v\n", err)
			os.Exit(1)
		}
		app.UpdateFieldName(fieldName)

		input := promptui.Prompt{
			Label: "Please search by typing value",
		}

		value, _ := input.Run()
		fmt.Printf("----- The value you typed is: %s ----- \n\n", value)

		app.Search(value)
	}
}

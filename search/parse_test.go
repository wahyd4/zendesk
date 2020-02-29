package search

import (
	"reflect"
	"testing"

	"github.com/wahyd4/zendesk/model"
)

func TestAPP_LoadOrganisationsFromJSON(t *testing.T) {
	type fields struct {
		app *APP
	}
	type args struct {
		jsonContent []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantApp *APP
	}{
		{
			name: "can load organisations from json",
			fields: fields{
				app: &APP{},
			},
			args: args{
				jsonContent: []byte(`[{
					"_id": 101,
					"url": "http://initech.zendesk.com/api/v2/organizations/101.json",
					"external_id": "9270ed79-35eb-4a38-a46f-35725197ea8d",
					"name": "Enthaze",
					"domain_names": [
					"kage.com",
					"ecratic.com",
					"endipin.com",
					"zentix.com"
					],
					"created_at": "2016-05-21T11:10:28 -10:00",
					"details": "MegaCorp",
					"shared_tickets": false,
					"tags": [
					"Fulton",
					"West",
					"Rodriguez",
					"Farley"
					]}]`),
			},
			wantErr: false,
			wantApp: &APP{
				organisations: map[string]*model.Organisation{
					"101": &model.Organisation{
						ID:         101,
						URL:        "http://initech.zendesk.com/api/v2/organizations/101.json",
						ExternalID: "9270ed79-35eb-4a38-a46f-35725197ea8d",
						Name:       "Enthaze",
						DomainNames: []string{
							"kage.com",
							"ecratic.com",
							"endipin.com",
							"zentix.com",
						},
						CreatedAt:     "2016-05-21T11:10:28 -10:00",
						Details:       "MegaCorp",
						SharedTickets: false,
						Tags: []string{
							"Fulton",
							"West",
							"Rodriguez",
							"Farley",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := tt.fields.app
			if err := app.LoadOrganisationsFromJSON(tt.args.jsonContent); (err != nil) != tt.wantErr {
				t.Errorf("APP.LoadOrganisationsFromJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(app, tt.wantApp) {
				t.Errorf("LoadOrganisationsFromJSON() got = %v, want %v", app, tt.wantApp)
			}

		})
	}
}

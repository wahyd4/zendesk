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
		{
			name: "cannot load organisations from json due to json is not valid",
			fields: fields{
				app: &APP{},
			},
			args: args{
				jsonContent: []byte(`[{
					"_id": 101,
					"url" "http://initech.zendesk.com/api/v2/organizations/101.json"
					]}]`),
			},
			wantErr: true,
			wantApp: &APP{},
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

func TestAPP_LoadUsersFromJSON(t *testing.T) {
	organisation := &model.Organisation{
		ID: 119,
	}

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
			name: "can load users from json",
			fields: fields{
				app: &APP{
					organisations: map[string]*model.Organisation{
						"119": organisation,
					},
				},
			},
			args: args{
				jsonContent: []byte(`[{
					  "_id": 1,
					  "url": "http://initech.zendesk.com/api/v2/users/1.json",
					  "external_id": "74341f74-9c79-49d5-9611-87ef9b6eb75f",
					  "name": "Francisca Rasmussen",
					  "alias": "Miss Coffey",
					  "created_at": "2016-04-15T05:19:46 -10:00",
					  "active": true,
					  "verified": true,
					  "shared": false,
					  "locale": "en-AU",
					  "timezone": "Sri Lanka",
					  "last_login_at": "2013-08-04T01:03:27 -10:00",
					  "email": "coffeyrasmussen@flotonic.com",
					  "phone": "8335-422-718",
					  "signature": "Don't Worry Be Happy!",
					  "organization_id": 119,
					  "tags": [
						"Springville",
						"Sutton",
						"Hartsville/Hartley",
						"Diaperville"
					  ],
					  "suspended": true,
					  "role": "admin"
					}]`),
			},
			wantErr: false,
			wantApp: &APP{
				organisations: map[string]*model.Organisation{
					"119": organisation,
				},
				users: map[string]*model.User{
					"1": &model.User{
						ID:           1,
						URL:          "http://initech.zendesk.com/api/v2/users/1.json",
						ExternalID:   "74341f74-9c79-49d5-9611-87ef9b6eb75f",
						Name:         "Francisca Rasmussen",
						Alias:        "Miss Coffey",
						CreatedAt:    "2016-04-15T05:19:46 -10:00",
						Active:       true,
						Verified:     true,
						Shared:       false,
						Locale:       "en-AU",
						Timezone:     "Sri Lanka",
						LastLoginAt:  "2013-08-04T01:03:27 -10:00",
						Email:        "coffeyrasmussen@flotonic.com",
						Phone:        "8335-422-718",
						Signature:    "Don't Worry Be Happy!",
						Organization: organisation,
						Tags: []string{
							"Springville",
							"Sutton",
							"Hartsville/Hartley",
							"Diaperville",
						},
						Suspended: true,
						Role:      "admin",
					},
				},
			},
		},
		{
			name: "cannot load users from json due to json is not valid",
			fields: fields{
				app: &APP{},
			},
			args: args{
				jsonContent: []byte(`[{
					"_id": 101,
					"url" "http://initech.zendesk.com/api/v2/organizations/101.json"
					]}]`),
			},
			wantErr: true,
			wantApp: &APP{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := tt.fields.app
			if err := app.LoadUsersFromJSON(tt.args.jsonContent); (err != nil) != tt.wantErr {
				t.Errorf("APP.LoadUsersFromJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(app, tt.wantApp) {
				t.Errorf("LoadUsersFromJSON() got = %v, want %v", app, tt.wantApp)
			}

		})
	}
}

func TestAPP_LoadTicketsFromJSON(t *testing.T) {
	organisation := &model.Organisation{
		ID: 119,
	}

	user := &model.User{
		ID:           111,
		Organization: organisation,
	}

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
			name: "can load tickets from json",
			fields: fields{
				app: &APP{
					organisations: map[string]*model.Organisation{
						"119": organisation,
					},
					users: map[string]*model.User{
						"111": user,
					},
				},
			},
			args: args{
				jsonContent: []byte(`[{
					  "_id": "436bf9b0-1147-4c0a-8439-6f79833bff5b",
					  "url": "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
					  "external_id": "9210cdc9-4bee-485f-a078-35396cd74063",
					  "created_at": "2016-04-28T11:19:34 -10:00",
					  "type": "incident",
					  "subject": "A Catastrophe in Korea (North)",
					  "description": "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
					  "priority": "high",
					  "status": "pending",
					  "submitter_id": 111,
					  "assignee_id": 111,
					  "organization_id": 119,
					  "tags": [
						"Ohio",
						"Pennsylvania",
						"American Samoa",
						"Northern Mariana Islands"
					  ],
					  "has_incidents": false,
					  "due_at": "2016-07-31T02:37:50 -10:00",
					  "via": "web"
					}]`),
			},
			wantErr: false,
			wantApp: &APP{
				organisations: map[string]*model.Organisation{
					"119": organisation,
				},
				users: map[string]*model.User{
					"111": user,
				},
				tickets: map[string]*model.Ticket{
					"436bf9b0-1147-4c0a-8439-6f79833bff5b": &model.Ticket{
						ID:           "436bf9b0-1147-4c0a-8439-6f79833bff5b",
						URL:          "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
						ExternalID:   "9210cdc9-4bee-485f-a078-35396cd74063",
						CreatedAt:    "2016-04-28T11:19:34 -10:00",
						Type:         "incident",
						Subject:      "A Catastrophe in Korea (North)",
						Description:  "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
						Priority:     "high",
						Status:       "pending",
						Submitter:    user,
						Assignee:     user,
						Organization: organisation,
						Tags: []string{
							"Ohio",
							"Pennsylvania",
							"American Samoa",
							"Northern Mariana Islands",
						},
						HasIncidents: false,
						DueAt:        "2016-07-31T02:37:50 -10:00",
						Via:          "web",
					},
				},
			},
		},
		{
			name: "cannot load tickets from json due to json is not valid",
			fields: fields{
				app: &APP{},
			},
			args: args{
				jsonContent: []byte(`[{
					"_id": 101,
					"url" "http://initech.zendesk.com/api/v2/organizations/101.json"
					]}]`),
			},
			wantErr: true,
			wantApp: &APP{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := tt.fields.app
			if err := app.LoadTicketsFromJSON(tt.args.jsonContent); (err != nil) != tt.wantErr {
				t.Errorf("APP.LoadTicketsFromJSON() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !reflect.DeepEqual(app, tt.wantApp) {
				t.Errorf("LoadTicketsFromJSON() got = %v, want %v", app, tt.wantApp)
			}

		})
	}
}

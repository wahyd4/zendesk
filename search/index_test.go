package search

import (
	"reflect"
	"testing"

	"github.com/wahyd4/zendesk/index"
)

func Test_extractFieldValues(t *testing.T) {
	type args struct {
		entity   map[string]interface{}
		fieldKey string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "can extract values from int and convert to string",
			args: args{
				entity: map[string]interface{}{
					"_id": float64(123),
				},
				fieldKey: "_id",
			},
			want: []string{"123"},
		},
		{
			name: "can extract values from string ",
			args: args{
				entity: map[string]interface{}{
					"subject": "A Nuisance in Equatorial Guinea",
				},
				fieldKey: "subject",
			},
			want: []string{"A Nuisance in Equatorial Guinea"},
		},
		{
			name: "can extract values from boolean and convert to string",
			args: args{
				entity: map[string]interface{}{
					"verified": true,
				},
				fieldKey: "verified",
			},
			want: []string{"true"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := extractFieldValues(tt.args.entity, tt.args.fieldKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("extractFieldValues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPP_BuildIndexes(t *testing.T) {
	type fields struct {
		jsonContents map[string][]byte
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "can build index for the application",
			fields: fields{
				jsonContents: map[string][]byte{
					"organisations": []byte(`[{
						"_id": 999,
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
					"users": []byte(`[{
						"_id": 222,
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
						"organization_id": 111,
						"tags": [
						  "Springville",
						  "Sutton",
						  "Hartsville/Hartley",
						  "Diaperville"
						],
						"suspended": true,
						"role": "admin"
					  }]`),
					"tickets": []byte(`[{
						"_id": "436bf9b0-1147-4c0a-8439-6f79833bff5b",
						"url": "http://initech.zendesk.com/api/v2/tickets/436bf9b0-1147-4c0a-8439-6f79833bff5b.json",
						"external_id": "9210cdc9-4bee-485f-a078-35396cd74063",
						"created_at": "2016-04-28T11:19:34 -10:00",
						"type": "incident",
						"subject": "A Catastrophe in Korea (North)",
						"description": "Nostrud ad sit velit cupidatat laboris ipsum nisi amet laboris ex exercitation amet et proident. Ipsum fugiat aute dolore tempor nostrud velit ipsum.",
						"priority": "high",
						"status": "pending",
						"submitter_id": 222,
						"assignee_id": 222,
						"organization_id": 999,
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
			},
			wantErr: false,
		},
		{
			name: "cannot build index for the application, due to the json content is not valid",
			fields: fields{
				jsonContents: map[string][]byte{
					"organisations": []byte(`invalid json string`),
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &APP{
				jsonContents: tt.fields.jsonContents,
				indexes:      make(map[string]index.SearchIndex),
			}
			if err := app.BuildIndexes(); (err != nil) != tt.wantErr {
				t.Errorf("APP.BuildIndexes() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && len(app.indexes) == 0 {
				t.Errorf("expected have some index but got nothing")
			}
		})
	}
}

package search

import (
	"testing"
)

func TestInitAPP(t *testing.T) {
	type args struct {
		organisationsFile string
		usersFile         string
		ticketsFile       string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "can init the application",
			args: args{
				organisationsFile: "../data/organizations.json",
				usersFile:         "../data/users.json",
				ticketsFile:       "../data/tickets.json",
			},
			wantErr: false,
		},
		{
			name: "cannot init the application due to users json is not found",
			args: args{
				organisationsFile: "../data/organizations.json",
				usersFile:         "not_exist.json",
				ticketsFile:       "../data/tickets.json",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InitAPP(tt.args.organisationsFile, tt.args.usersFile, tt.args.ticketsFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitAPP() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && got == nil {
				t.Error("expect to init a app but didn't ")
			}

			if tt.wantErr && got != nil {
				t.Error("expect to receive some error but didn't")
			}
		})
	}
}

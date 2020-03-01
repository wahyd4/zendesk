package search

import (
	"reflect"
	"testing"
)

func TestAPP_UpdateDataType(t *testing.T) {
	type fields struct {
		app *APP
	}

	type args struct {
		dataType string
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantSearchContext *SearchContext
	}{
		{
			name: "can update search context data type ",
			fields: fields{
				app: &APP{
					searchContext: &SearchContext{},
				},
			},
			args: args{
				dataType: "users",
			},
			wantSearchContext: &SearchContext{
				dataType:  "users",
				fieldName: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := tt.fields.app

			app.UpdateDataType(tt.args.dataType)
			if !reflect.DeepEqual(app.searchContext, tt.wantSearchContext) {
				t.Errorf("UpdateDataType() got = %v, want %v", app.searchContext, tt.wantSearchContext)
			}
		})
	}
}

func TestAPP_UpdateFieldName(t *testing.T) {
	type fields struct {
		app *APP
	}

	type args struct {
		fieldName string
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantSearchContext *SearchContext
	}{
		{
			name: "can update search context field name",
			fields: fields{
				app: &APP{
					searchContext: &SearchContext{
						dataType: "tickets",
					},
				},
			},
			args: args{
				fieldName: "_id",
			},
			wantSearchContext: &SearchContext{
				dataType:  "tickets",
				fieldName: "_id",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := tt.fields.app

			app.UpdateFieldName(tt.args.fieldName)
			if !reflect.DeepEqual(app.searchContext, tt.wantSearchContext) {
				t.Errorf("UpdateFieldName() got = %v, want %v", app.searchContext, tt.wantSearchContext)
			}
		})
	}
}

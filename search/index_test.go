package search

import (
	"reflect"
	"testing"
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

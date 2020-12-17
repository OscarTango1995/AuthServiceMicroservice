package models

import (
	"reflect"
	"testing"
)

func TestUser_Map(t *testing.T) {
	type fields struct {
		ID    int
		Name  string
		Email string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success - convert User struct to map",
			fields: fields{
				ID:    0,
				Name:  "Ovais",
				Email: "abcdef@abc.com",
			},
			want: map[string]interface{}{
				"id":    0,
				"name":  "Ovais",
				"email": "abcdef@abc.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:    tt.fields.ID,
				Name:  tt.fields.Name,
				Email: tt.fields.Email,
			}
			if got := u.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_Names(t *testing.T) {
	type fields struct {
		ID    int
		Name  string
		Email string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: " success - names of student struct fields",
			fields: fields{
				ID:    12345,
				Name:  "Ovais",
				Email: "abc@abc.com",
			},
			want: []string{"id", "name", "email"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:    tt.fields.ID,
				Name:  tt.fields.Name,
				Email: tt.fields.Email,
			}
			if got := u.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}

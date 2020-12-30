package models

import (
	"reflect"
	"testing"
	"time"
)

func TestSession_Map(t *testing.T) {
	type fields struct {
		ID        int
		UserID    int
		Token     string
		CreatedAt time.Time
		ExpiresAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success - convert Session struct to map",
			fields: fields{
				ID:        0,
				UserID:    1,
				Token:     "abcdef",
				CreatedAt: time.Time{},
				ExpiresAt: time.Time{},
			},
			want: map[string]interface{}{
				"id":         0,
				"user_id":    1,
				"token":      "abcdef",
				"created_at": time.Time{},
				"expires_at": time.Time{},
			},
		},
		{
			name: "success - convert Session struct to map2",
			fields: fields{
				ID:        1,
				UserID:    3,
				Token:     "qwertyuiop",
				CreatedAt: time.Time{},
				ExpiresAt: time.Time{},
			},
			want: map[string]interface{}{
				"id":         1,
				"user_id":    3,
				"token":      "qwertyuiop",
				"created_at": time.Time{},
				"expires_at": time.Time{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				ID:        tt.fields.ID,
				UserID:    tt.fields.UserID,
				Token:     tt.fields.Token,
				CreatedAt: tt.fields.CreatedAt,
				ExpiresAt: tt.fields.ExpiresAt,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_Names(t *testing.T) {
	type fields struct {
		ID        int
		UserID    int
		Token     string
		CreatedAt time.Time
		ExpiresAt time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: " success - names of session struct fields",
			fields: fields{
				ID:        12345,
				UserID:    1,
				Token:     "abcdef",
				CreatedAt: time.Time{},
				ExpiresAt: time.Time{},
			},
			want: []string{"id", "user_id", "token", "created_at", "expires_at"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Session{
				ID:        tt.fields.ID,
				UserID:    tt.fields.UserID,
				Token:     tt.fields.Token,
				CreatedAt: tt.fields.CreatedAt,
				ExpiresAt: tt.fields.ExpiresAt,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}

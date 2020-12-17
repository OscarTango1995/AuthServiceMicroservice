package models

import (
	"github.com/fatih/structs"
	"time"
)

type Session struct {
	ID        int       `json:"id" structs:"id" bson:"_id" db:"id"`
	UserID    int       `json:"user_id" structs:"user_id" bson:"user_id" db:"user_id"`
	Token     string    `json:"token" structs:"token" bson:"token" db:"token"`
	CreatedAt time.Time `json:"created_at" structs:"created_at" bson:"created_at" db:"created_at"`
	ExpiresAt time.Time `json:"expires_at" structs:"expires_at" bson:"expires_at" db:"expires_at"`
}

// Structure to a map representation
func (s *Session) Map() map[string]interface{} {
	return structs.Map(s)
}

// Names returns the field names of Session model
func (s *Session) Names() []string {
	fields := structs.Fields(s)
	names := make([]string, len(fields))

	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}

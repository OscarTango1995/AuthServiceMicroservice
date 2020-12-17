package models

import (
	"github.com/fatih/structs"
)

type User struct {
	ID    int    `json:"id" structs:"id" bson:"_id" db:"id"`
	Name  string `json:"name" structs:"name" bson:"name" db:"name"`
	Email string `json:"email" structs:"email" bson:"email" db:"email"`
}

// Structure to a map representation
func (u *User) Map() map[string]interface{} {
	return structs.Map(u)
}

// Names returns the field names of User model
func (u *User) Names() []string {
	fields := structs.Fields(u)
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

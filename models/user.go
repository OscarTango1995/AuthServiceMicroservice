package models

import (
	"time"

	"github.com/fatih/structs"
)

type User struct {
	ID        string    `json:"id" structs:"id" bson:"_id" db:"id"`
	Name      string    `json:"name" structs:"name" bson:"name" db:"name"`
	Email     string    `json:"email" structs:"email" bson:"email" db:"email"`
	Password  string    `json:"password" structs:"password" bson:"password" db:"password"`
	CreatedAt time.Time `json:"created_at" structs:"created_at" bson:"created_at" db:"created_at"`
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

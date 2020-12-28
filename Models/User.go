package Models

import (
	"github.com/fatih/structs"
)

type User struct {
	id int
	name string
	email string
}

// Structure to a map representation
func (u *User) Map() map[string]interface{} {
	return structs.Map(u)
}
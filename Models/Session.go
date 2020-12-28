package Models

import (
	"github.com/fatih/structs"
)

type Session struct {
	id int
	user_id string
	token string
	created_at string
	expires_at string
}

// Structure to a map representation
func (s *Session) Map() map[string]interface{} {
	return structs.Map(s)
}
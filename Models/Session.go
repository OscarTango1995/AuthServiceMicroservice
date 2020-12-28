package Models

import (
	"github.com/fatih/structs"
	"time"
)

type Session struct {
	id int
	user_id string
	token string
	created_at time.Time
	expires_at time.Time
}

// Structure to a map representation
func (s *Session) Map() map[string]interface{} {
	return structs.Map(s)
}
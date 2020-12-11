// Author: dat.nguyen

package store

//Users -
type Users struct {
	UserID string `gorm:"primary_key" json:"user_id"`
	Role   int    `json:"role"`
	Email  string `json:"email"`
	TeamID string `json:"team_id"`
}

//Role enums -
const (
	AdminHub = iota + 1
	AdminTeam
	Manager
	User
	Guess
)

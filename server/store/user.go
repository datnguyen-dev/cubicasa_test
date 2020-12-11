// Author: dat.nguyen

package store

//Users -
type Users struct {
	UserID string `gorm:"primary_key"`
	Role   int
	Email  string
	TeamID string
}

//Role enums -
const (
	AdminHub = iota + 1
	AdminTeam
	Manager
	User
	Guess
)

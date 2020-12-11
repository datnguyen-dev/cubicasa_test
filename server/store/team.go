// Author: dat.nguyen

package store

//Team -
type Team struct {
	TeamID string `gorm:"primary_key"`
	Name   string
	Type   string
	HubID  string
}

//TeamType -
const (
	Develop      = "DE"
	Test         = "TE"
	ProductOwner = "PO"
)

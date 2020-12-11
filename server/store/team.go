// Author: dat.nguyen

package store

//Team -
type Team struct {
	TeamID string `gorm:"primary_key" json:"team_id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	HubID  string `json:"hub_id"`
}

//TeamType -
const (
	Develop      = "DE"
	Test         = "TE"
	ProductOwner = "PO"
)

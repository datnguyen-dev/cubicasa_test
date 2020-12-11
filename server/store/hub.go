// Author: dat.nguyen

package store

//Hub -
type Hub struct {
	HubID       string `gorm:"primary_key" json:"hub_id"`
	Name        string `json:"name"`
	GeoLocation string `json:"geo_location"`
}

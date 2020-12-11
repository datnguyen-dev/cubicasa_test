// Author: dat.nguyen

package store

//Hub -
type Hub struct {
	HubID       string `gorm:"primary_key"`
	Name        string
	GeoLocation string
}

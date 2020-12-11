package postgresdb

import (
	"testing"

	"datnguyen.cubicasa.test/store"
	"github.com/stretchr/testify/assert"
)

func TestHub(t *testing.T) {
	s, tea := getConnection()
	defer tea()

	//add
	hub := &store.Hub{
		HubID:       "",
		Name:        "Hub 01",
		GeoLocation: "101.10100020;543.0888900",
	}
	id, err := s.Hubs().AddHub(hub)
	assert.True(t, id != "")
	assert.True(t, err == nil)

	//update
	hub.Name = "Change hub to 02"
	val, err := s.Hubs().UpdateHub(hub)
	assert.True(t, err == nil)
	assert.True(t, val)

	//delete
	val, err = s.Hubs().DeleteHub(hub.HubID)
	assert.True(t, err == nil)
	assert.True(t, val)
}

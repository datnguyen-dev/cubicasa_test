package postgresdb

import (
	"testing"

	"datnguyen.cubicasa.test/store"
	"github.com/stretchr/testify/assert"
)

func TestTeam(t *testing.T) {
	s, tea := getConnection()
	defer tea()

	//add
	team := &store.Team{
		TeamID: "",
		Name:   "Team 01",
		Type:   "Dev",
		HubID:  "",
	}
	id, err := s.Teams().AddTeam(team)
	assert.True(t, id != "")
	assert.True(t, err == nil)

	//update
	team.Name = "Change team to 02"
	val, err := s.Teams().UpdateTeam(team)
	assert.True(t, err == nil)
	assert.True(t, val)

	//delete
	val, err = s.Teams().DeleteTeam(team.TeamID)
	assert.True(t, err == nil)
	assert.True(t, val)
}

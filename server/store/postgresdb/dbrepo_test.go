package postgresdb

import (
	"testing"

	"datnguyen.cubicasa.test/store"
	"github.com/stretchr/testify/assert"
)

func TestRepo(t *testing.T) {
	s, teadown := getConnection()
	defer teadown()
	res := s.DBRepo().CheckDB()
	assert.True(t, res)

	//create db by script
	res, err := s.DBRepo().Install()
	assert.True(t, err == nil)
	assert.True(t, res)

	//insert hub
	hub := &store.Hub{
		HubID:       "",
		Name:        "Hub 01",
		GeoLocation: "101.10100020;543.0888900",
	}
	idHub, err := s.Hubs().AddHub(hub)
	assert.True(t, idHub != "")
	assert.True(t, err == nil)

	//insert team
	team := &store.Team{
		TeamID: "",
		Name:   "Team 01",
		Type:   "Dev",
		HubID:  "",
	}
	idTeam, err := s.Teams().AddTeam(team)
	assert.True(t, idTeam != "")
	assert.True(t, err == nil)

	//insert User
	user := &store.Users{
		UserID: "",
		Email:  "User@01.com",
		Role:   0,
	}
	idUser, err := s.Users().AddUser(user)
	assert.True(t, idUser != "")
	assert.True(t, err == nil)

	//join team to hub
	val, err := s.DBRepo().JoinTeamIntoHub(idTeam, idHub)
	assert.True(t, err == nil)
	assert.True(t, val)

	//join user to team
	val, err = s.DBRepo().JoinUserIntoTeam(idUser, idTeam, 1)
	assert.True(t, err == nil)
	assert.True(t, val)

	//search hub name
	hubs, err := s.DBRepo().SearchHubByName("01")
	assert.True(t, err == nil)
	assert.True(t, hubs != nil)

	//search team name
	hubs, err = s.DBRepo().SearchTeamByName("01")
	assert.True(t, err == nil)
	assert.True(t, hubs != nil)

	//delete user
	val, err = s.Users().DeleteUser(idUser)
	assert.True(t, err == nil)
	assert.True(t, val)

	//delete team
	val, err = s.Teams().DeleteTeam(idTeam)
	assert.True(t, err == nil)
	assert.True(t, val)

	//delete hub
	val, err = s.Hubs().DeleteHub(idHub)
	assert.True(t, err == nil)
	assert.True(t, val)
}

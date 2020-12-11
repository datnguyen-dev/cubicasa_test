package postgresdb

import (
	"testing"

	"datnguyen.cubicasa.test/store"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	s, tea := getConnection()
	defer tea()

	//add
	user := &store.Users{
		UserID: "",
		Email:  "User@01.com",
		Role:   1,
	}
	id, err := s.Users().AddUser(user)
	assert.True(t, id != "")
	assert.True(t, err == nil)

	//update
	user.Email = "User02@02.com"
	val, err := s.Users().UpdateUser(user)
	assert.True(t, err == nil)
	assert.True(t, val)

	//delete
	val, err = s.Users().DeleteUser(user.UserID)
	assert.True(t, err == nil)
	assert.True(t, val)
}

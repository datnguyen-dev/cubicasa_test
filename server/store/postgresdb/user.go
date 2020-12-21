package postgresdb

import (
	"fmt"

	"datnguyen.cubicasa.test/common"
	"datnguyen.cubicasa.test/store"
	"github.com/jinzhu/gorm"
)

type users struct {
	db *gorm.DB
}

//AddUser - Add user to database
func (u *users) AddUser(user *store.Users) (string, error) {
	if user.UserID == "" {
		user.UserID = common.GenUID()
	} else {
		var us *store.Users
		u.db.First(&us, "user_id = ?", user.UserID)
		if us != nil {
			return "", fmt.Errorf("Duplicate")
		}
	}
	user.Role = 0
	user.TeamID = ""
	u.db.Create(user)
	return user.UserID, nil
}

//UpdateUser - Update user existed
func (u *users) UpdateUser(user *store.Users) (bool, error) {
	if user.UserID == "" {
		return false, fmt.Errorf("NotFound")
	}
	var us store.Users
	u.db.First(&us, "user_id = ?", user.UserID)
	if &us == nil {
		return false, fmt.Errorf("NotFound")
	}

	u.db.Model(us).Updates(user)
	return true, nil
}

//DeleteUser - Delete user existed
func (u *users) DeleteUser(id string) (bool, error) {
	if id == "" {
		return false, fmt.Errorf("NotFound")
	}
	var us store.Users
	u.db.First(&us, "user_id = ?", id)
	if &us == nil {
		return false, fmt.Errorf("NotFound")
	}
	u.db.Delete(us, "user_id = ?", id)
	return true, nil
}

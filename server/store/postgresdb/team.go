package postgresdb

import (
	"fmt"

	"datnguyen.cubicasa.test/common"
	"datnguyen.cubicasa.test/store"
	"github.com/jinzhu/gorm"
)

type team struct {
	db *gorm.DB
}

func (t *team) AddTeam(team *store.Team) (string, error) {
	if team.TeamID == "" {
		team.TeamID = common.GenUID()
	} else {
		var te *store.Team
		t.db.First(&te, "team_id = ?", team.TeamID)
		if te != nil {
			return "", fmt.Errorf("Duplicate")
		}
	}
	team.HubID = ""
	t.db.Create(team)
	return team.TeamID, nil
}

func (t *team) UpdateTeam(team *store.Team) (bool, error) {
	if team.TeamID == "" {
		return false, fmt.Errorf("NotFound")
	}
	var te store.Team
	t.db.First(&te, "team_id = ?", team.TeamID)
	if &te == nil {
		return false, fmt.Errorf("NotFound")
	}
	t.db.Model(te).Updates(team)
	return true, nil
}

func (t *team) DeleteTeam(id string) (bool, error) {
	if id == "" {
		return false, fmt.Errorf("NotFound")
	}
	var te store.Team
	t.db.First(&te, "team_id = ?", id)
	if &te == nil {
		return false, fmt.Errorf("NotFound")
	}
	t.db.Delete(te, "team_id = ?", id)
	return true, nil
}

package postgresdb

import (
	"fmt"
	"io/ioutil"

	"datnguyen.cubicasa.test/store"
	"github.com/jinzhu/gorm"
)

type dbrepo struct {
	db   *gorm.DB
	file string
}

func (r *dbrepo) Install() (bool, error) {
	cmd, err := ioutil.ReadFile(r.file)
	if err != nil {
		return false, err
	}
	rows := r.db.Exec(fmt.Sprintf("%s", cmd))
	if rows.Error != nil {
		return false, rows.Error
	}
	return true, nil
}

func (r *dbrepo) CheckDB() bool {
	cmd := `SELECT table_schema FROM information_schema.tables 
		WHERE  table_name = ?`
	var schema string
	row := r.db.Raw(cmd, "hubs").Row()
	row.Scan(&schema)
	return schema != ""
}

func (r *dbrepo) SearchHubByName(hubName string) (interface{}, error) {
	type rest struct {
		hub   *store.Hub
		teams []*store.Team
	}
	var res []rest
	var hub []*store.Hub

	r.db.Table("hubs").Where("name like ?", "%"+hubName+"%").Find(&hub)
	if len(hub) > 0 {
		for _, h := range hub {
			var te []*store.Team
			r.db.Table("teams").Where("hub_id = ?", h.HubID).Find(&te)
			res = append(res, rest{
				hub:   h,
				teams: te,
			})
		}
	}
	return res, nil
}

func (r *dbrepo) SearchTeamByName(teamName string) (interface{}, error) {
	type rest struct {
		hub   *store.Hub
		teams []*store.Team
	}
	var res []rest

	var team []*store.Team
	r.db.Table("teams").Where("name like ?", "%"+teamName+"%").Find(&team)
	if len(team) > 0 {
		lst := make(map[string][]*store.Team)

		for _, t := range team {
			if t.HubID != "" {
				lst[t.HubID] = append(lst[t.HubID], t)
			}
		}
		for k, v := range lst {
			var hu store.Hub
			r.db.Table("hubs").Where("hub_id = ?", k).First(&hu)
			res = append(res, rest{
				hub:   &hu,
				teams: v,
			})
		}
	}
	return res, nil
}

func (r *dbrepo) JoinTeamIntoHub(teamID string, hubID string) (bool, error) {
	var team store.Team
	r.db.Table("teams").Where("team_id = ?", teamID).Scan(&team)
	if &team != nil {
		var hub store.Hub
		r.db.Table("hubs").Where("hub_id = ?", hubID).Scan(&hub)
		if &hub != nil {
			row := r.db.Exec("UPDATE teams SET hub_id = ?", hubID)
			return row.RowsAffected > 0, nil
		}
	}
	return false, fmt.Errorf("NotFound")
}

func (r *dbrepo) JoinUserIntoTeam(userID string, teamID string, roleID int) (bool, error) {
	var user store.Users
	r.db.Table("users").Where("user_id = ?", userID).Scan(&user)
	if &user != nil {
		var team store.Team
		r.db.Table("teams").Where("team_id = ?", teamID).Scan(&team)
		if &team != nil {
			row := r.db.Exec("UPDATE users SET team_id = ?, role = ?", teamID, roleID)
			return row.RowsAffected > 0, nil
		}
	}
	return false, fmt.Errorf("NotFound")
}

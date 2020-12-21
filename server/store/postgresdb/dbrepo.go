package postgresdb

import (
	"fmt"
	"io/ioutil"
	"strings"

	"datnguyen.cubicasa.test/store"
	"github.com/jinzhu/gorm"
)

type dbrepo struct {
	db   *gorm.DB
	file string
}

//Install - Initial DB with script file from the PATH return true/false
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

//CheckDB - Check DB exsited or not return true/false
func (r *dbrepo) CheckDB() bool {
	cmd := `SELECT table_schema FROM information_schema.tables 
		WHERE  table_name = ?`
	var schema string
	row := r.db.Raw(cmd, "hubs").Row()
	row.Scan(&schema)
	return schema != ""
}

//SearchHubByName - Search hub via Name of hub and return Hub and Team reference
func (r *dbrepo) SearchHubByName(hubName string) (interface{}, error) {
	type rest struct {
		Hub   store.Hub    `json:"hub"`
		Teams []store.Team `json:"teams"`
	}
	var res []rest
	var hub []store.Hub

	hubName = strings.ReplaceAll(hubName, `"`, ``)
	r.db.Table("hubs").Where("name like ?", fmt.Sprintf("%s%v%s", "%", hubName, "%")).Find(&hub)

	if len(hub) > 0 {
		for _, h := range hub {
			var te []store.Team
			r.db.Table("teams").Where("hub_id = ?", h.HubID).Find(&te)
			res = append(res, rest{
				Hub:   h,
				Teams: te,
			})
		}
	}
	return res, nil
}

// SearchTeamByName - Search team via Name of team return list Hubs and Teams references
func (r *dbrepo) SearchTeamByName(teamName string) (interface{}, error) {
	type rest struct {
		Hub   store.Hub    `json:"hub"`
		Teams []store.Team `json:"teams"`
	}
	var res []rest

	var team []store.Team
	teamName = strings.ReplaceAll(teamName, `"`, ``)
	r.db.Table("teams").Where("name like ?", fmt.Sprintf("%s%v%s", "%", teamName, "%")).Find(&team)
	if len(team) > 0 {
		lst := make(map[string][]store.Team)
		for _, t := range team {
			if t.HubID != "" {
				lst[t.HubID] = append(lst[t.HubID], t)
			}
		}
		for k, v := range lst {
			var hu store.Hub
			r.db.Table("hubs").Where("hub_id = ?", k).First(&hu)
			res = append(res, rest{
				Hub:   hu,
				Teams: v,
			})
		}
	}
	return res, nil
}

//JoinTeanIntoHub - Join Team into Hub with HubID
func (r *dbrepo) JoinTeamIntoHub(teamID string, hubID string) (bool, error) {
	var team store.Team
	r.db.Table("teams").Where("team_id = ?", teamID).Scan(&team)
	if team != (store.Team{}) {
		var hub store.Hub
		r.db.Table("hubs").Where("hub_id = ?", hubID).Scan(&hub)
		if hub != (store.Hub{}) {
			row := r.db.Exec("UPDATE teams SET hub_id = ?", hubID)
			return row.RowsAffected > 0, nil
		}
	}
	return false, fmt.Errorf("NotFound")
}

//JoinUserIntoTeam - Join User into Team with teamid and roleID
func (r *dbrepo) JoinUserIntoTeam(userID string, teamID string, roleID int) (bool, error) {
	var user store.Users
	r.db.Table("users").Where("user_id = ?", userID).Scan(&user)
	if user != (store.Users{}) {
		var team store.Team
		r.db.Table("teams").Where("team_id = ?", teamID).Scan(&team)
		if team != (store.Team{}) {
			row := r.db.Exec("UPDATE users SET team_id = ?, role = ?", teamID, roleID)
			return row.RowsAffected > 0, nil
		}
	}
	return false, fmt.Errorf("NotFound")
}

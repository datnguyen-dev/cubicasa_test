package api

import (
	"fmt"
	"net/http"
	"strconv"

	"datnguyen.cubicasa.test/store"
)

type process struct {
	*Handler
}

func (p *process) checkDB(w http.ResponseWriter, r *http.Request) {
	res := p.store.DBRepo().CheckDB()
	p.render(w, http.StatusOK, struct {
		DBExisted bool `json:"db_existed"`
	}{res})
}

func (p *process) install(w http.ResponseWriter, r *http.Request) {
	res := p.store.DBRepo().CheckDB()
	if !res || res {
		_, err := p.store.DBRepo().Install()
		if err != nil {
			p.render(w, http.StatusBadRequest, struct {
				Message string `json:"message"`
			}{err.Error()})
			return
		}
		p.render(w, http.StatusOK, struct {
			Message string `json:"message"`
		}{"OK"})
		return
	}
	p.render(w, http.StatusBadRequest, struct {
		Message string `json:"message"`
	}{"DB Existed!"})
}

func (p *process) createHub(w http.ResponseWriter, r *http.Request) {
	var request store.Hub
	err := p.parseRequest(r, &request)
	if err != nil {
		p.render(w, http.StatusBadRequest, "invalid request body "+err.Error())
		return
	}
	fmt.Print(request)
	id, err := p.store.Hubs().AddHub(&request)
	if err != nil {
		p.render(w, http.StatusBadRequest, "add hub error "+err.Error())
		return
	}
	p.render(w, http.StatusOK, struct {
		ID string `json:"id"`
	}{id})
}

func (p *process) createTeam(w http.ResponseWriter, r *http.Request) {
	var request store.Team
	err := p.parseRequest(r, &request)
	if err != nil {
		p.render(w, http.StatusBadRequest, "invalid request body "+err.Error())
		return
	}
	id, err := p.store.Teams().AddTeam(&request)
	if err != nil {
		p.render(w, http.StatusBadRequest, "add team error "+err.Error())
		return
	}
	p.render(w, http.StatusOK, struct {
		ID string `json:"id"`
	}{id})
}

func (p *process) createUser(w http.ResponseWriter, r *http.Request) {
	var request store.Users
	err := p.parseRequest(r, &request)
	if err != nil {
		p.render(w, http.StatusBadRequest, "invalid request body "+err.Error())
		return
	}
	id, err := p.store.Users().AddUser(&request)
	if err != nil {
		p.render(w, http.StatusBadRequest, "add user error "+err.Error())
		return
	}
	p.render(w, http.StatusOK, struct {
		ID string `json:"id"`
	}{id})
}

func (p *process) updateHub(w http.ResponseWriter, r *http.Request) {
	var request store.Hub
	err := p.parseRequest(r, &request)
	if err != nil {
		p.render(w, http.StatusBadRequest, "invalid request body "+err.Error())
		return
	}
	res, err := p.store.Hubs().UpdateHub(&request)
	if err != nil {
		p.render(w, http.StatusBadRequest, "update hub error "+err.Error())
		return
	}
	p.render(w, http.StatusOK, struct {
		Status bool `json:"status"`
	}{res})
}

func (p *process) updateTeam(w http.ResponseWriter, r *http.Request) {
	var request store.Team
	err := p.parseRequest(r, &request)
	if err != nil {
		p.render(w, http.StatusBadRequest, "invalid request body "+err.Error())
		return
	}
	res, err := p.store.Teams().UpdateTeam(&request)
	if err != nil {
		p.render(w, http.StatusBadRequest, "update team error "+err.Error())
		return
	}
	p.render(w, http.StatusOK, struct {
		Status bool `json:"status"`
	}{res})
}

func (p *process) updateUser(w http.ResponseWriter, r *http.Request) {
	var request store.Users
	err := p.parseRequest(r, &request)
	if err != nil {
		p.render(w, http.StatusBadRequest, "invalid request body "+err.Error())
		return
	}
	res, err := p.store.Users().UpdateUser(&request)
	if err != nil {
		p.render(w, http.StatusBadRequest, "update user error "+err.Error())
		return
	}
	p.render(w, http.StatusOK, struct {
		Status bool `json:"status"`
	}{res})
}

func (p *process) deleteHub(w http.ResponseWriter, r *http.Request) {
	id := p.urlParam(r, "id")
	if len(id) == 0 {
		p.render(w, http.StatusBadRequest, "param error")
		return
	}
	res, err := p.store.Hubs().DeleteHub(id)
	if err != nil {
		p.render(w, http.StatusBadRequest, "delete hub error "+err.Error())
		return
	}
	p.render(w, http.StatusOK, struct {
		Status bool `json:"status"`
	}{res})
}

func (p *process) deleteTeam(w http.ResponseWriter, r *http.Request) {
	id := p.urlParam(r, "id")
	if len(id) == 0 {
		p.render(w, http.StatusBadRequest, "param error")
		return
	}
	res, err := p.store.Teams().DeleteTeam(id)
	if err != nil {
		p.render(w, http.StatusBadRequest, "delete team error "+err.Error())
		return
	}
	p.render(w, http.StatusOK, struct {
		Status bool `json:"status"`
	}{res})
}

func (p *process) deleteUser(w http.ResponseWriter, r *http.Request) {
	id := p.urlParam(r, "id")
	if len(id) == 0 {
		p.render(w, http.StatusBadRequest, "param error")
		return
	}
	res, err := p.store.Users().DeleteUser(id)
	if err != nil {
		p.render(w, http.StatusBadRequest, "delete user error "+err.Error())
		return
	}
	p.render(w, http.StatusOK, struct {
		Status bool `json:"status"`
	}{res})
}

func (p *process) searchHub(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		p.render(w, http.StatusBadRequest, "param error")
		return
	}
	res, err := p.store.DBRepo().SearchHubByName(name)
	if err != nil {
		p.render(w, http.StatusBadRequest, "search hub error "+err.Error())
		return
	}

	p.render(w, http.StatusOK, struct {
		Data interface{} `json:"Data"`
	}{res})
}

func (p *process) searchTeam(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if len(name) == 0 {
		p.render(w, http.StatusBadRequest, "param error")
		return
	}
	res, err := p.store.DBRepo().SearchTeamByName(name)
	if err != nil {
		p.render(w, http.StatusBadRequest, "search team error "+err.Error())
		return
	}
	p.render(w, http.StatusOK, struct {
		Data interface{} `json:"Data"`
	}{res})
}

func (p *process) joinTeam(w http.ResponseWriter, r *http.Request) {
	teamid := p.urlParam(r, "teamid")
	if len(teamid) == 0 {
		p.render(w, http.StatusBadRequest, "param teamid error")
		return
	}
	hubid := p.urlParam(r, "hubid")
	if len(hubid) == 0 {
		p.render(w, http.StatusBadRequest, "param hubid error")
		return
	}
	res, err := p.store.DBRepo().JoinTeamIntoHub(teamid, hubid)
	if err != nil {
		p.render(w, http.StatusBadRequest, "join team error "+err.Error())
		return
	}
	p.render(w, http.StatusOK, struct {
		Message bool `json:"message"`
	}{res})
}

func (p *process) joinUser(w http.ResponseWriter, r *http.Request) {
	userid := p.urlParam(r, "userid")
	if len(userid) == 0 {
		p.render(w, http.StatusBadRequest, "param userid error")
		return
	}
	teamid := p.urlParam(r, "teamid")
	if len(teamid) == 0 {
		p.render(w, http.StatusBadRequest, "param teamid error")
		return
	}
	role, err := strconv.Atoi(p.urlParam(r, "role"))
	if err != nil {
		p.render(w, http.StatusBadRequest, "param role error")
		return
	}
	res, err := p.store.DBRepo().JoinUserIntoTeam(userid, teamid, role)
	if err != nil {
		p.render(w, http.StatusBadRequest, "join user error "+err.Error())
		return
	}
	p.render(w, http.StatusOK, struct {
		Message bool `json:"message"`
	}{res})
}

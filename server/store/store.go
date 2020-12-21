// Author: dat.nguyen

package store

//Store -
type Store interface {
	Users() UserStore
	Teams() TeamStore
	Hubs() HubStore
	DBRepo() DBRepoStore
}

//UserStore - Store all Functionality of User
type UserStore interface {
	AddUser(*Users) (string, error)
	UpdateUser(*Users) (bool, error)
	DeleteUser(string) (bool, error)
}

//TeamStore - Store all Functionaluty of Team
type TeamStore interface {
	AddTeam(*Team) (string, error)
	UpdateTeam(*Team) (bool, error)
	DeleteTeam(string) (bool, error)
}

//HubStore - Store all functionality of Hub
type HubStore interface {
	AddHub(*Hub) (string, error)
	UpdateHub(*Hub) (bool, error)
	DeleteHub(string) (bool, error)
}

//DBRepoStore - Store all functionality of connection to Team/Hub/User
type DBRepoStore interface {
	Install() (bool, error)
	CheckDB() bool
	SearchHubByName(string) (interface{}, error)
	SearchTeamByName(string) (interface{}, error)
	JoinTeamIntoHub(string, string) (bool, error)
	JoinUserIntoTeam(string, string, int) (bool, error)
}

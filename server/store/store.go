// Author: dat.nguyen

package store

//Store -
type Store interface {
	Users() UserStore
	Teams() TeamStore
	Hubs() HubStore
	DBRepo() DBRepoStore
}

//UserStore -
type UserStore interface {
	AddUser(*Users) (string, error)
	UpdateUser(*Users) (bool, error)
	DeleteUser(string) (bool, error)
}

//TeamStore -
type TeamStore interface {
	AddTeam(*Team) (string, error)
	UpdateTeam(*Team) (bool, error)
	DeleteTeam(string) (bool, error)
}

//HubStore -
type HubStore interface {
	AddHub(*Hub) (string, error)
	UpdateHub(*Hub) (bool, error)
	DeleteHub(string) (bool, error)
}

//DBRepoStore -
type DBRepoStore interface {
	Install() (bool, error)
	CheckDB() bool
	SearchHubByName(string) (interface{}, error)
	SearchTeamByName(string) (interface{}, error)
	JoinTeamIntoHub(string, string) (bool, error)
	JoinUserIntoTeam(string, string, int) (bool, error)
}

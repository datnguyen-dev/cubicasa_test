package postgresdb

import (
	"datnguyen.cubicasa.test/store"
	"github.com/jinzhu/gorm"
)

//Store - define implement Store
type Store struct {
	users  *users
	teams  *team
	hubs   *hub
	dbrepo *dbrepo
}

//Users -
func (s *Store) Users() store.UserStore {
	return s.users
}

//Teams -
func (s *Store) Teams() store.TeamStore {
	return s.teams
}

//Hubs -
func (s *Store) Hubs() store.HubStore {
	return s.hubs
}

//DBRepo -
func (s *Store) DBRepo() store.DBRepoStore {
	return s.dbrepo
}

// Check valid interface implementation
var _ store.Store = (*Store)(nil)

//InitDB - Init connection to database
func InitDB(db *gorm.DB, file string) *Store {
	db.AutoMigrate(&store.Users{})
	db.AutoMigrate(&store.Hub{})
	db.AutoMigrate(&store.Team{})
	s := &Store{
		users:  &users{db: db},
		teams:  &team{db: db},
		hubs:   &hub{db: db},
		dbrepo: &dbrepo{db: db, file: file},
	}
	return s
}

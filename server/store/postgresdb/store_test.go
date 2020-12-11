package postgresdb

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234567890"
	dbname   = "cubicasa_db"
)

func getConnection() (*Store, func()) {
	var db *gorm.DB
	cmdPath := "../../setup/script.sql"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	s := InitDB(db, cmdPath)
	teawdown := func() {
		db.Exec("DROP TABLE Users, Hubs, Teams")
		defer db.Close()
	}
	return s, teawdown
}

package main

import (
	"fmt"
	"net/http"

	"datnguyen.cubicasa.test/api"
	"datnguyen.cubicasa.test/config"
	"datnguyen.cubicasa.test/store/postgresdb"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	AppName := "cubicasa_test"
	ConfigFile := fmt.Sprintf("%s.conf", AppName)
	cnf, err := config.InitConfig(ConfigFile)
	if err != nil {
		panic(err)
	}
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cnf.Postgres.Host, cnf.Postgres.Port, cnf.Postgres.UserName,
		cnf.Postgres.Password, cnf.Postgres.DBName, cnf.Postgres.SSL)

	var db *gorm.DB
	db, err = gorm.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	//init database
	s := postgresdb.InitDB(db, cnf.Postgres.SetupFile)
	apiHandler := api.New(s)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Mount("/", apiHandler)
	fmt.Printf("starting the server: %s", "3100")
	http.ListenAndServe(":3100", r)
}

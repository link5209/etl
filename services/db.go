package services

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var db *sqlx.DB

func init() {
	// conn db
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"))

	var err error
	db, err = sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(10)

	fmt.Println("init db connection success!")
}

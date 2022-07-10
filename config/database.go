package config

import (
	"database/sql"
	"fmt"
	"os"
)

var Db *sql.DB

func Connect() error {
	var err error
	Db, err = sql.Open("postgres", os.Getenv("DB_SERVER_URL"))
	if err != nil {
		fmt.Println("error config", err)
		return err
	}
	if err = Db.Ping(); err != nil {
		return err
	}

	return nil
}

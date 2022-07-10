package config

import (
	"fmt"

	"github.com/aryaputra777/rest/utils"
)

func Createtableuser() error {
	var err error
	query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS users ( 	id serial NOT NULL,
		username varchar(200) NULL,
		"name" varchar(200) NULL,
		"password" varchar(200) NULL,
		picture text NULL,
		id_role int4 NULL,
		email varchar(100) NULL,
		phone varchar(20) NULL,
		status bool NULL,
		created_at timestamp NULL,
		updated_at timestamp NULL )`)

	_, err = Db.Exec(query)
	if err != nil {
		return err
	}

	bycrypt, _ := utils.HashPassword("admin")
	resp, _ := Db.Query("SELECT  a.id, a.username, a.name, a.email, a.status, a.phone, a.id_role FROM users as a WHERE username=$1 or email=$2", "admin", "admin@gmail.com")
	if resp == nil {
		insertUser := fmt.Sprintf(`
	INSERT INTO users (username, name, email,status,phone, password, id_role)
	VALUES ('admin', 'admin', 'admin@gmail.com', true, '08888', '%s', 1)
	`, bycrypt)
		_, err = Db.Exec(insertUser)
	}

	return nil
}

package models

import "github.com/hanbarfe/rest_example/db"

func CreateUser(user User) (id int64, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO users (name, email, isadmin) VALUES ($1, $2, $3) RETURNING id`

	err = conn.QueryRow(sql, user.Name, user.Email, user.IsAdmin).Scan(&id)

	return
}

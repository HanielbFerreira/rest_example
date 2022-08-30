package models

import (
	"log"

	"github.com/hanbarfe/rest_example/db"
)

func GetUserById(id int64) (user User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `SELECT * FROM users WHERE id=$1`

	row := conn.QueryRow(sql, id)

	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.IsAdmin)

	return
}

func GetAllUsers() (users []User, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `SELECT * FROM users`

	rows, err := conn.Query(sql)

	for rows.Next() {
		var user User

		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.IsAdmin)

		if err != nil {
			log.Fatal("Um erro aconteceu", err)
			continue
		}

		users = append(users, user)
	}

	return
}

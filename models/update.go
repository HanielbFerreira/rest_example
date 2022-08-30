package models

import "github.com/hanbarfe/rest_example/db"

func UpdateUserById(id int64, user User) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE users SET name=$1, email=$2, isadmin=$3 WHERE id=$4`, user.Name, user.Email, user.IsAdmin, user.ID)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}

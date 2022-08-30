package models

import "github.com/hanbarfe/rest_example/db"

func DeleteUserById(id int64) (int64, error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE from users WHERE id=$1`, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}

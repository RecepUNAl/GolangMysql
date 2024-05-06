package account

import (
	"database/sql"
	"goMysql/types"
)

var tableName string = "users"

func SingupModel(post *types.Singup, db *sql.DB) error {
	query := ""
	insert, err := db.Query(query)
	if err != nil {
		return err
	}
	defer insert.Close()
	return nil
}

func SinginModel(post *types.Singin, db *sql.DB) error {
	return nil
}

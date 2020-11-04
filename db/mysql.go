package db

import "database/sql"

type MySql struct {
	Connection *sql.DB
}

func (db *MySql) GetIDFromToken(Token string) (uint, error) {
	stmt, err := db.Connection.Prepare(`select user_id from user_tokens where token = ?`)
	if err != nil {
		return 0, err
	}

	var id uint
	err = stmt.QueryRow(Token).Scan(&id)

	if err != nil {
		return 0, err
	}
	return id, nil
}

func (db)
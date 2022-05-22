package models

import (
	"database/sql"
	"houseduino-be/constants"
)

func GetTestSqlx(db *sql.DB) string {

	esito := ""

	db.Exec(constants.TEST_SET)
	err := db.QueryRow(constants.TEST_GET).Scan(&esito)
	if err != nil {
		PrintErrorLog("Test", err)
	}
	return esito
}

package config

import (
	"fmt"
	"houseduino-be/constants"
	"houseduino-be/lang"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDBSqlx() *sqlx.DB {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", constants.DBHOST, constants.DBPORT, constants.DBUSERNAME, constants.DBPASSWORD, constants.DBNAME)
	PrintLog("Config", psqlconn)
	// open database
	db, err := sqlx.Open(constants.DBTYPE, psqlconn)
	if err != nil {
		PrintErrorLog("Config", err)
		panic(err.Error())
	} else {
		PrintLog("Config", "DB ok!")
	}

	return db
}

func ErrHandler(err error) string {
	var errmessage string
	if os.Getenv("DEBUG") == "true" {
		errmessage = err.Error()
	} else {
		errmessage = lang.Get("something_went_wrong")
	}
	return errmessage
}

func PrintErrorLog(area string, err error) {
	log.Printf("%s -> Errore: %s", area, ErrHandler(err))

}
func PrintLog(area string, message string) {
	log.Printf("%s -> %s", area, message)

}
func PrintStringErrorLog(area string, messaggio string) {
	log.Printf("%s -> Errore: %s", area, messaggio)

}

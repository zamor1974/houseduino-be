package models

import (
	"database/sql"
	"fmt"
	"houseduino-be/constants"
	"houseduino-be/lang"
	"log"
	"time"
)

// swagger:model Rain
type Rain struct {
	// Id of rain value
	// in: int64
	Id int64 `json:"id"`
	// Value of rain
	// in: int
	Value int `json:"valore"`
	// Timestamp of insert
	// in: time
	DateInsert time.Time `json:"data_inserimento"`
}

type Rains []Rain

type ReqAddRain struct {
	// Value of the rain
	// in: int
	Value int `json:"valore" validate:"required"`
}

// swagger:parameters addRain
type ReqRainBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/ReqAddRain"
	//  required: true
	Body ReqAddRain `json:"body"`
}

func GetRainsSqlx(db *sql.DB) *Rains {
	rains := Rains{}
	rows, err := db.Query(constants.RAIN_GET)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Rain
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			PrintErrorLog("Pioggia", err)
		}
		rains = append(rains, p)
	}
	return &rains
}
func GetLastRainSqlx(db *sql.DB) *Rains {
	rains := Rains{}
	rows, err := db.Query(constants.RAIN_GET_LAST)
	if err != nil {
		PrintErrorLog("Pioggia", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Rain
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			PrintErrorLog("Pioggia", err)
		}
		rains = append(rains, p)
	}
	return &rains
}
func GetRainsLastHourSqlx(db *sql.DB) *Rains {
	rains := Rains{}

	tFine := time.Now()
	dataFine := tFine.Format("2006-01-02 15:04:05")

	tInizio := time.Now().Add(time.Duration(-1) * time.Hour)
	dataInizio := tInizio.Format("2006-01-02 15:04:05")

	sqlStatement := fmt.Sprintf(constants.RAIN_GET_LAST_HOUR, dataInizio, dataFine)

	rows, err := db.Query(sqlStatement)
	if err != nil {
		PrintErrorLog("Pioggia", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Rain
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			PrintErrorLog("Pioggia", err)
		}
		rains = append(rains, p)
	}

	if len(rains) == 0 {
		elemento := GetLastRainSqlx(db)
		rains = append(rains, *elemento...)
	}
	return &rains
}

// PostRainSqlx insert rain value
func PostRainSqlx(db *sql.DB, reqrain *ReqAddRain) (*Rain, string) {

	value := reqrain.Value

	var rain Rain

	lastInsertId := 0

	//sqlStatement := fmt.Sprintf("insert into 'pioggia' ('valore','data_inserimento') values (%d,CURRENT_TIMESTAMP) RETURNING id", value)
	sqlStatement := fmt.Sprintf(constants.RAIN_POST_DATA, value)

	err := db.QueryRow(sqlStatement).Scan(&lastInsertId)

	if err != nil {
		PrintErrorLog("Pioggia", err)
		return &rain, ErrHandler(err)
	}

	sqlStatement1 := fmt.Sprintf("SELECT id,valore,data_inserimento FROM pioggia where id = %d", lastInsertId)
	rows, err := db.Query(sqlStatement1)

	if err != nil {
		PrintErrorLog("Pioggia", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Rain
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			PrintErrorLog("Pioggia", err)
		}
		rain = p
	}
	if err != nil {
		PrintErrorLog("Pioggia", err)
		return &rain, lang.Get("no_result")
	}
	return &rain, ""
}
func GetRainShowDataSqlx(db *sql.DB, recordNumber int) *Rains {
	rains := Rains{}

	sqlStatement := fmt.Sprintf(constants.RAIN_GET_SHOWDATA, recordNumber)
	rows, err := db.Query(sqlStatement)
	if err != nil {
		PrintErrorLog("Pioggia", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Rain
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			PrintErrorLog("Pioggia", err)
		}
		rains = append(rains, p)
	}

	return &rains
}

package models

import (
	"database/sql"
	"fmt"
	"houseduino-be/constants"
	"houseduino-be/lang"
	"log"
	"time"
)

// swagger:model Temperature
type Temperature struct {
	// Id of Temperature value
	// in: int64
	Id int64 `json:"id"`
	// Value of Temperature
	// in: float
	Value float32 `json:"valore"`
	// Timestamp of insert
	// in: time
	DateInsert time.Time `json:"data_inserimento"`
}

type Temperatures []Temperature

type ReqAddTemperature struct {
	// Value of the Temperature
	// in: float
	Value float32 `json:"valore" validate:"required"`
}

// swagger:parameters addTemperature
type ReqTemperatureBody struct {
	// - name: body
	//  in: body
	//  description: Temperature
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/ReqAddTemperature"
	//  required: true
	Body ReqAddTemperature `json:"body"`
}

func GetTemperaturesSqlx(db *sql.DB) *Temperatures {
	temperatures := Temperatures{}
	rows, err := db.Query(constants.TEMPERATURE_GET)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Temperature
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		temperatures = append(temperatures, p)
	}
	return &temperatures
}
func GetLastTemperatureSqlx(db *sql.DB) *Temperatures {
	temperatures := Temperatures{}
	rows, err := db.Query(constants.TEMPERATURE_GET_LAST)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Temperature
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		temperatures = append(temperatures, p)
	}
	return &temperatures
}
func GetTemperaturesLastHourSqlx(db *sql.DB) *Temperatures {
	temperatures := Temperatures{}

	tFine := time.Now()
	dataFine := tFine.Format("2006-01-02 15:04:05")

	tInizio := time.Now().Add(time.Duration(-1) * time.Hour)
	dataInizio := tInizio.Format("2006-01-02 15:04:05")

	sqlStatement := fmt.Sprintf(constants.TEMPERATURE_GET_LAST_HOUR, dataInizio, dataFine)

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Temperature
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		temperatures = append(temperatures, p)
	}

	if len(temperatures) == 0 {
		elemento := GetLastTemperatureSqlx(db)
		temperatures = append(temperatures, *elemento...)
	}
	return &temperatures
}

// PostTemperatureSqlx insert Temperature value
func PostTemperatureSqlx(db *sql.DB, reqTemperature *ReqAddTemperature) (*Temperature, string) {

	value := reqTemperature.Value

	var temperature Temperature

	lastInsertId := 0

	//sqlStatement := fmt.Sprintf("insert into 'pioggia' ('valore','data_inserimento') values (%d,CURRENT_TIMESTAMP) RETURNING id", value)
	sqlStatement := fmt.Sprintf(constants.TEMPERATURE_POST_DATA, value)

	err := db.QueryRow(sqlStatement).Scan(&lastInsertId)

	if err != nil {
		return &temperature, ErrHandler(err)
	}

	sqlStatement1 := fmt.Sprintf("SELECT id,valore,data_inserimento FROM temperatura where id = %d", lastInsertId)
	rows, err := db.Query(sqlStatement1)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Temperature
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		temperature = p
	}
	if err != nil {
		return &temperature, lang.Get("no_result")
	}
	return &temperature, ""
}
func GetTemperatureShowDataSqlx(db *sql.DB, recordNumber int) *Temperatures {
	temperatures := Temperatures{}

	sqlStatement := fmt.Sprintf(constants.TEMPERATURE_GET_SHOWDATA, recordNumber)
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Temperature
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		temperatures = append(temperatures, p)
	}

	return &temperatures
}

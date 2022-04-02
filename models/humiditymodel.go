package models

import (
	"database/sql"
	"fmt"
	"houseduino-be/constants"
	"houseduino-be/lang"
	"log"
	"time"
)

// swagger:model Humidity
type Humidity struct {
	// Id of rain value
	// in: int64
	Id int64 `json:"id"`
	// Value of Humidity
	// in: int
	Value float32 `json:"valore"`
	// Timestamp of insert
	// in: time
	DateInsert time.Time `json:"data_inserimento"`
}

type Humidities []Humidity

type ReqAddHumidity struct {
	// Value of the Humidity
	// in: int
	Value float32 `json:"valore" validate:"required"`
}

// swagger:parameters addHumidity
type ReqHumidityBody struct {
	// - name: body
	//  in: body
	//  description: Humidity
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/ReqAddHumidity"
	//  required: true
	Body ReqAddHumidity `json:"body"`
}

func GetHumiditiesSqlx(db *sql.DB) *Humidities {
	humidities := Humidities{}
	rows, err := db.Query(constants.HUMIDITY_GET)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Humidity
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		humidities = append(humidities, p)
	}
	return &humidities
}
func GetLastHumiditySqlx(db *sql.DB) *Humidities {
	humidities := Humidities{}
	rows, err := db.Query(constants.HUMIDITY_GET_LAST)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Humidity
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		humidities = append(humidities, p)
	}
	return &humidities
}
func GetHumiditiesLastHourSqlx(db *sql.DB) *Humidities {
	humidities := Humidities{}

	tFine := time.Now()
	dataFine := tFine.Format("2006-01-02 15:04:05")

	tInizio := time.Now().Add(time.Duration(-1) * time.Hour)
	dataInizio := tInizio.Format("2006-01-02 15:04:05")

	sqlStatement := fmt.Sprintf(constants.HUMIDITY_GET_LAST_HOUR, dataInizio, dataFine)

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Humidity
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		humidities = append(humidities, p)
	}

	if len(humidities) == 0 {
		elemento := GetLastHumiditySqlx(db)
		humidities = append(humidities, *elemento...)
	}
	return &humidities
}

// PostHumiditySqlx insert Humidity value
func PostHumiditySqlx(db *sql.DB, reqHumidity *ReqAddHumidity) (*Humidity, string) {

	value := reqHumidity.Value

	var humidity Humidity

	lastInsertId := 0

	//sqlStatement := fmt.Sprintf("insert into 'pioggia' ('valore','data_inserimento') values (%d,CURRENT_TIMESTAMP) RETURNING id", value)
	sqlStatement := fmt.Sprintf(constants.HUMIDITY_POST_DATA, value)
	log.Println(sqlStatement)
	err := db.QueryRow(sqlStatement).Scan(&lastInsertId)

	if err != nil {
		return &humidity, ErrHandler(err)
	}

	rows, err := db.Query(constants.HUMIDITY_GET_LAST)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Humidity
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		humidity = p
	}
	if err != nil {
		return &humidity, lang.Get("no_result")
	}
	return &humidity, ""
}
func GetHumidityShowDataSqlx(db *sql.DB, recordNumber int) *Humidities {
	humidities := Humidities{}

	sqlStatement := fmt.Sprintf(constants.HUMIDITY_GET_SHOWDATA, recordNumber)
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Humidity
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		humidities = append(humidities, p)
	}

	return &humidities
}

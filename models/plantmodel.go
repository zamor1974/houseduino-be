package models

import (
	"database/sql"
	"fmt"
	"houseduino-be/constants"
	"houseduino-be/lang"
	"log"
	"strconv"
	"time"
)

// swagger:model Plant
type Plant struct {
	// Id of plant
	// in: int64
	Id int64 `json:"id"`
	// Value of Name
	// in: string
	Value string `json:"nome"`
	// Timestamp of insert
	// in: time
	DateInsert time.Time `json:"data_inserimento"`
}
type Plants []Plant

// swagger:model PlantHumidity
type PlantHumidity struct {
	// Id of rain value
	// in: int64
	Id int64 `json:"id"`
	// Id of plant
	// in: int64
	IdPlant int64 `json:"id_plant"`
	// Value of Humidity
	// in: int
	Value float32 `json:"valore"`
	// Timestamp of insert
	// in: time
	DateInsert time.Time `json:"data_inserimento"`
}

type PlantHumidities []PlantHumidity

// swagger:model PlantValue
type PlantValue struct {
	// Id of rain value
	// in: int64
	Id int64 `json:"id"`
	// Name of plant
	// in: string
	Name string `json:"nome"`
	// Value of Humidity
	// in: int
	Value float32 `json:"valore"`
	// Timestamp of insert
	// in: time
	DateInsert time.Time `json:"data_inserimento"`
}

type ReqAddPlantHumidity struct {
	// Id of plant
	// in: int64
	IdPlant int64 `json:"id_plant"`
	// Value of the PlantHumidity
	// in: int
	Value float32 `json:"valore" validate:"required"`
}

// swagger:parameters addPlantHumidity
type ReqPlantHumidityBody struct {
	// - name: body
	//  in: body
	//  description: PlantHumidity
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/ReqAddPlantHumidity"
	//  required: true
	Body ReqAddPlantHumidity `json:"body"`
}

func GetPlantHumiditiesSqlx(db *sql.DB, idPlant string) *PlantHumidities {
	humidities := PlantHumidities{}

	sqlStatement := fmt.Sprintf(constants.PLANT_HUMIDITY_GET, idPlant)

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p PlantHumidity
		if err := rows.Scan(&p.Id, &p.IdPlant, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		humidities = append(humidities, p)
	}
	return &humidities
}
func GetLastPlantHumiditySqlx(db *sql.DB, idPlant string) *PlantHumidities {
	humidities := PlantHumidities{}
	sqlStatement := fmt.Sprintf(constants.PLANT_HUMIDITY_GET_LAST, idPlant)
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p PlantHumidity
		if err := rows.Scan(&p.Id, &p.IdPlant, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		humidities = append(humidities, p)
	}
	return &humidities
}
func GetPlantHumiditiesLastHourSqlx(db *sql.DB, idPlant string) *PlantHumidities {
	humidities := PlantHumidities{}

	tFine := time.Now()
	dataFine := tFine.Format("2006-01-02 15:04:05")

	tInizio := time.Now().Add(time.Duration(-1) * time.Hour)
	dataInizio := tInizio.Format("2006-01-02 15:04:05")

	sqlStatement := fmt.Sprintf(constants.PLANT_HUMIDITY_GET_LAST_HOUR, idPlant, dataInizio, dataFine)

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p PlantHumidity
		if err := rows.Scan(&p.Id, &p.IdPlant, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		humidities = append(humidities, p)
	}

	if len(humidities) == 0 {
		elemento := GetLastPlantHumiditySqlx(db, idPlant)
		humidities = append(humidities, *elemento...)
	}
	return &humidities
}
func GetPlantLastSqlx(db *sql.DB, idPlant string) *PlantValue {
	valore := PlantValue{}

	sqlStatement := fmt.Sprintf(constants.PLANT_HUMIDITY_GET_LAST_VALUE, idPlant)

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p PlantValue
		if err := rows.Scan(&p.Id, &p.Name, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		valore = p
	}

	if valore.Id == 0 {
		var elemento PlantValue
		elemento.Id = 0
		elemento.Name = ""
		elemento.Value = -1

		valore = elemento
	}
	return &valore
}

// PostHumiditySqlx insert Humidity value
func PostPlantHumiditySqlx(db *sql.DB, reqPlantHumidity *ReqAddPlantHumidity) (*PlantHumidity, string) {

	value := reqPlantHumidity.Value
	idPlant := strconv.FormatInt(reqPlantHumidity.IdPlant, 10)
	var humidity PlantHumidity

	lastInsertId := 0

	//sqlStatement := fmt.Sprintf("insert into 'pioggia' ('valore','data_inserimento') values (%d,CURRENT_TIMESTAMP) RETURNING id", value)
	sqlStatement := fmt.Sprintf(constants.PLANT_HUMIDITY_POST_DATA, idPlant, value)
	//log.Println(sqlStatement)
	err := db.QueryRow(sqlStatement).Scan(&lastInsertId)

	if err != nil {
		log.Println(err)
		return &humidity, ErrHandler(err)
	}
	sqlStatement2 := fmt.Sprintf(constants.PLANT_HUMIDITY_GET_LAST, idPlant)
	//log.Println(sqlStatement2)
	rows, err := db.Query(sqlStatement2)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p PlantHumidity
		if err := rows.Scan(&p.Id, &p.IdPlant, &p.Value, &p.DateInsert); err != nil {
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
func GetPlantHumidityShowDataSqlx(db *sql.DB, idPlant string, recordNumber int) *PlantHumidities {
	humidities := PlantHumidities{}

	sqlStatement := fmt.Sprintf(constants.PLANT_HUMIDITY_GET_SHOWDATA, idPlant, recordNumber, idPlant)
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p PlantHumidity
		if err := rows.Scan(&p.Id, &p.IdPlant, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		humidities = append(humidities, p)
	}

	return &humidities
}
func GetPlantAllSqlx(db *sql.DB) *Plants {
	humidities := Plants{}

	sqlStatement := fmt.Sprintf(constants.PLANT_GET)

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Plant
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		humidities = append(humidities, p)
	}
	return &humidities
}

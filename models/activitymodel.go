package models

import (
	"database/sql"
	"fmt"
	"houseduino-be/constants"
	"houseduino-be/lang"
	"time"
)

// swagger:model Activity
type Activity struct {
	// Id of Activity value
	// in: int64
	Id int64 `json:"id"`
	// Value of Activity
	// in: int
	Value int `json:"valore"`
	// Timestamp of insert
	// in: time
	DateInsert time.Time `json:"data_inserimento"`
}

type Activities []Activity

type IsActive bool

type ReqAddActivity struct {
	// Value of the Activity
	// in: int
	Value int `json:"valore" validate:"required"`
}

// swagger:parameters addActivity
type ReqActivityBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/ReqAddActivity"
	//  required: true
	Body ReqAddActivity `json:"body"`
}

func GetIsActiveSqlx(db *sql.DB) bool {
	var contatore int64

	err := db.QueryRow(constants.ACTIVITY_ISACTIVE).Scan(&contatore)
	if err != nil {
		PrintErrorLog("Attività", err)
	}

	return contatore > 0
}

func GetActivitiesSqlx(db *sql.DB) *Activities {
	activities := Activities{}

	rows, err := db.Query(constants.ACTIVITY_GET)

	PrintLog("Attività", constants.ACTIVITY_GET)

	if err != nil {
		PrintErrorLog("Attività", err)
	} else {

		defer rows.Close()

		for rows.Next() {
			var p Activity
			if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
				PrintErrorLog("Attività", err)
			}

			activities = append(activities, p)
		}
	}

	return &activities
}
func GetLastActivitySqlx(db *sql.DB) *Activities {
	activities := Activities{}
	rows, err := db.Query(constants.ACTIVITY_GET_LAST)
	if err != nil {
		PrintErrorLog("Attività", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Activity
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			PrintErrorLog("Attività", err)
		}
		activities = append(activities, p)
	}
	return &activities
}
func GetActivitiesLastHourSqlx(db *sql.DB) *Activities {
	activities := Activities{}

	tFine := time.Now()
	dataFine := tFine.Format("2006-01-02 15:04:05")

	tInizio := time.Now().Add(time.Duration(-1) * time.Hour)
	dataInizio := tInizio.Format("2006-01-02 15:04:05")

	sqlStatement := fmt.Sprintf(constants.ACTIVITY_GET_LASTHOUR, dataInizio, dataFine)

	rows, err := db.Query(sqlStatement)
	if err != nil {
		PrintErrorLog("Attività", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Activity
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			PrintErrorLog("Attività", err)
		}
		activities = append(activities, p)
	}

	if len(activities) == 0 {
		elemento := GetLastActivitySqlx(db)
		activities = append(activities, *elemento...)
	}
	return &activities
}

// PostActivitySqlx insert Activity value
func PostActivitySqlx(db *sql.DB, reqrain *ReqAddActivity) (*Activity, string) {

	//value := reqrain.Value

	var activity Activity

	lastInsertId := 0

	//sqlStatement := fmt.Sprintf("insert into 'pioggia' ('valore','data_inserimento') values (%d,CURRENT_TIMESTAMP) RETURNING id", value)
	sqlStatement := "insert into attivita (data_inserimento) values (CURRENT_TIMESTAMP) RETURNING id"
	//log.Println(sqlStatement)
	err := db.QueryRow(sqlStatement).Scan(&lastInsertId)

	if err != nil {
		PrintErrorLog("Attività", err)
		return &activity, ErrHandler(err)
	}

	sqlStatement1 := fmt.Sprintf("SELECT id,0,data_inserimento FROM attivita where id = %d", lastInsertId)
	rows, err := db.Query(sqlStatement1)

	if err != nil {
		PrintErrorLog("Attività", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Activity
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			PrintErrorLog("Attività", err)
		}
		activity = p
	}
	if err != nil {
		PrintErrorLog("Attività", err)
		return &activity, lang.Get("no_result")
	}
	return &activity, ""
}

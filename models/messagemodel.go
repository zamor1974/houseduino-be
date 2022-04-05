package models

import (
	"database/sql"
	"fmt"
	"houseduino-be/constants"
	"houseduino-be/lang"
	"log"
	"time"
)

// swagger:model Message
type Message struct {
	// Id of Message value
	// in: int64
	Id int64 `json:"id"`
	// Value of Message
	// in: int
	Value string `json:"messaggio"`
	// Timestamp of insert
	// in: time
	DateInsert time.Time `json:"data_inserimento"`
}

type Messages []Message

type ReqAddMessage struct {
	// Value of the Message
	// in: string
	Value string `json:"messaggio" validate:"required"`
}

// swagger:parameters addMessage
type ReqMessageBody struct {
	// - name: body
	//  in: body
	//  description: Message
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/ReqAddMessage"
	//  required: true
	Body ReqAddMessage `json:"body"`
}

func GetLastMessageSqlx(db *sql.DB) *Messages {
	messages := Messages{}
	rows, err := db.Query(constants.MESSAGE_GET_LAST)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Message
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		messages = append(messages, p)
	}
	return &messages
}

func GetMessagesLastHourSqlx(db *sql.DB) *Messages {
	messages := Messages{}

	tFine := time.Now()
	dataFine := tFine.Format("2006-01-02 15:04:05")

	tInizio := time.Now().Add(time.Duration(-1) * time.Hour)
	dataInizio := tInizio.Format("2006-01-02 15:04:05")

	sqlStatement := fmt.Sprintf(constants.MESSAGE_GET_LASTHOUR, dataInizio, dataFine)

	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Message
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		messages = append(messages, p)
	}

	if len(messages) == 0 {
		elemento := GetMessagessSqlx(db)
		messages = append(messages, *elemento...)
	}
	return &messages
}
func GetMessagessSqlx(db *sql.DB) *Messages {
	messages := Messages{}
	rows, err := db.Query(constants.MESSAGE_GET_LAST)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Message
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			log.Fatal(err)
		}
		messages = append(messages, p)
	}
	return &messages
}

// PostMessageSqlx insert Message value
func PostMessageSqlx(db *sql.DB, reqmessage *ReqAddMessage) (*Message, string) {

	value := reqmessage.Value

	var message Message

	lastInsertId := 0

	//sqlStatement := fmt.Sprintf("insert into 'pioggia' ('valore','data_inserimento') values (%d,CURRENT_TIMESTAMP) RETURNING id", value)
	sqlStatement := fmt.Sprintf(constants.MESSAGE_POST_DATA, value)
	//log.Println(sqlStatement)
	err := db.QueryRow(sqlStatement).Scan(&lastInsertId)

	if err != nil {
		return &message, ErrHandler(err)
	}

	rows, err := db.Query(constants.MESSAGE_GET_LAST)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var p Message
		if err := rows.Scan(&p.Id, &p.Value, &p.DateInsert); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		message = p
	}
	if err != nil {
		return &message, lang.Get("no_result")
	}
	return &message, ""
}

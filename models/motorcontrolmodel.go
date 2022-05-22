package models

import (
	"database/sql"
	"fmt"
	"houseduino-be/constants"
)

// swagger:model MotorStatus
type MotorStatus struct {
	// Value of Name
	// in: bool
	Value bool `json:"nome"`
}

func GetMotorStatus(db *sql.DB) *MotorStatus {
	var esito int32
	esito = -1

	sqlStatement := fmt.Sprintf(constants.MOTOR_GET, "1")

	err := db.QueryRow(sqlStatement).Scan(&esito)
	if err != nil {
		PrintErrorLog("Test", err)
	}
	risultato := MotorStatus{}
	pianta := GetPlantSqlx(db, "1")

	PrintLog("Motor", fmt.Sprintf("Pianta %s -> valore: %d", pianta.Value, esito))

	if esito < 35 {
		risultato.Value = true
	} else if esito > 70 {
		risultato.Value = false
	} else if esito == -1 {
		risultato.Value = false
	} else {
		risultato.Value = false
	}

	return &risultato
}

func GetMotorStatus2(db *sql.DB, idPlant string) *MotorStatus {
	var esito int32
	esito = -1

	sqlStatement := fmt.Sprintf(constants.MOTOR_GET, idPlant)

	err := db.QueryRow(sqlStatement).Scan(&esito)
	if err != nil {
		PrintErrorLog("Test", err)
	}
	risultato := MotorStatus{}
	pianta := GetPlantSqlx(db, idPlant)

	PrintLog("Motor", fmt.Sprintf("Pianta %s -> valore: %d", pianta.Value, esito))

	if esito == -1 {
		risultato.Value = false
	} else if esito > 0 && esito < 35 {
		risultato.Value = true
	} else if esito > 70 {
		risultato.Value = false
	} else {
		risultato.Value = false
	}

	return &risultato
}

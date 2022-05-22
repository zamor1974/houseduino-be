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
	humidities := MotorStatus{}

	sqlStatement := fmt.Sprintf(constants.MOTOR_GET_ALL)

	rows, err := db.Query(sqlStatement)
	if err != nil {
		PrintErrorLog("Motore", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p MotorStatus
		if err := rows.Scan(&p.Value); err != nil {
			PrintErrorLog("Motore", err)
		}
		return &p
	}
	return &humidities
}

func GetMotorStatus2(db *sql.DB, idPlant string) *MotorStatus {
	humidities := MotorStatus{}

	sqlStatement := fmt.Sprintf(constants.MOTOR_GET, idPlant)

	rows, err := db.Query(sqlStatement)
	if err != nil {
		PrintErrorLog("Motore", err)
	}
	defer rows.Close()

	for rows.Next() {
		var p MotorStatus
		if err := rows.Scan(&p.Value); err != nil {
			PrintErrorLog("Motore", err)
		}
		return &p
	}
	return &humidities
}

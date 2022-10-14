package models

import (
	"database/sql"
	"houseduino-be/constants"
	"strconv"
)

// swagger:model WeatherPrevision
const (
	sunny         string = "wi-day-sunny"
	rain_mix             = "wi-day-rain-mix"
	rain                 = "wi-rain"
	storm_showers        = "wi-storm-showers"
	cloudy               = "wi-cloudy"
	snow                 = "wi-snow"
)

// swagger:model Weather
type Weather struct {
	// Value of temperature
	// in: float32
	Temperature float32 `json:"temperature"`
	// Value of humidity
	// in: float32
	Humidity float32 `json:"humidity"`
	// Value of pressure
	// in: float32
	Pressure float32 `json:"pressure"`
	// Weather prevision
	// in: string
	WeatherPrevision string `json:"weather_prevision"`
	// Weather description
	// in: string
	WeatherDescription string `json:"weather_description"`

	//Last update
	//in: string
	LastUpdate string `json:"last_update"`
}

type Prevision struct {
	PressureType   string
	TemperatureMin float32
	TemperatureMax float32
	PressureMin    float32
	PressureMax    float32
}

func GetWeatherSqlx(db *sql.DB) *Weather {
	weather := Weather{}

	var temperature = GetLastTemperature2Sqlx(db)
	var humidity = GetLastHumidity2Sqlx(db)
	var pressure = GetLastPressure2Sqlx(db)
	var lastUpdate = GetLastActivityDatetimeSqlx(db)
	//var prevision = GetPrevision(db)

	weather.Temperature = temperature.Value
	weather.Humidity = humidity.Value
	weather.Pressure = pressure.Value
	weather.LastUpdate = lastUpdate

	if weather.Pressure > 1014 {
		weather.WeatherPrevision = sunny
		weather.WeatherDescription = "Sole"
	}
	if weather.Pressure <= 1014 && weather.Pressure > 1008 {
		weather.WeatherPrevision = cloudy
		weather.WeatherDescription = "Variabile"
	}
	if weather.Pressure <= 1008 && weather.Pressure > 1005 {
		weather.WeatherPrevision = rain
		weather.WeatherDescription = "Pioggia"
	}
	if weather.Pressure <= 1005 {
		weather.WeatherPrevision = storm_showers
		weather.WeatherDescription = "Temporale"
	}

	return &weather
}

func GetPrevision(db *sql.DB) Prevision {
	prevision := Prevision{}

	rows, err := db.Query(constants.PREVISION_GET)
	if err != nil {
		PrintErrorLog("Meteo", err)
	}
	defer rows.Close()

	for rows.Next() {
		var descr string
		var valore string
		if err := rows.Scan(&descr, &valore); err != nil {
			PrintErrorLog("Meteo", err)
		}
		switch descr {
		case "TIPO PRESSIONE":
			prevision.PressureType = valore
		case "PRESSIONE MINIMA":
			floatNum, err := strconv.ParseFloat(valore, 32)
			if err != nil {
				PrintErrorLog("Meteo", err)
			}
			prevision.PressureMin = float32(floatNum)
			break
		case "PRESSIONE MASSIMA":
			floatNum, err := strconv.ParseFloat(valore, 32)
			if err != nil {
				PrintErrorLog("Meteo", err)
			}
			prevision.PressureMax = float32(floatNum)
			break
		case "TEMPERATURA MINIMA":
			floatNum, err := strconv.ParseFloat(valore, 32)
			if err != nil {
				PrintErrorLog("Meteo", err)
			}
			prevision.TemperatureMin = float32(floatNum)
			break
		case "TEMPERATURA MASSIMA":
			floatNum, err := strconv.ParseFloat(valore, 32)
			if err != nil {
				PrintErrorLog("Meteo", err)
			}
			prevision.TemperatureMax = float32(floatNum)
			break
		}

	}

	return prevision
}

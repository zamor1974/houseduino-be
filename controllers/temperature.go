package controllers

import (
	"encoding/json"
	"fmt"
	"houseduino-be/lang"
	"houseduino-be/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:model GetTemperatures
type GetTemperatures struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string               `json:"message"`
	Data    *models.Temperatures `json:"data"`
}

// swagger:model GetTemperature
type GetTemperature struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string `json:"message"`
	// Umidity value
	Data *models.Temperature `json:"data"`
}

// swagger:route GET /temperature/all temperature temperatureAll
// Get Temperature list
//
// responses:
//  401: CommonError
//  200: GetTemperatures
func (h *BaseHandlerSqlx) GetTemperaturesSqlx(w http.ResponseWriter, r *http.Request) {
	response := GetTemperatures{}

	temperatures := models.GetTemperaturesSqlx(h.db.DB)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = temperatures

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// swagger:route GET /temperature/lasthour temperature temperatureLastHour
// Get list of last hour of temperature values .... or the last value inserted
//
// responses:
//  401: CommonError
//  200: GetTemperatures
func (h *BaseHandlerSqlx) GetTemperaturesLastHourSqlx(w http.ResponseWriter, r *http.Request) {
	response := GetTemperatures{}

	temperatures := models.GetTemperaturesLastHourSqlx(h.db.DB)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = temperatures

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// swagger:route POST /temperature/insert temperature addTemperature
// Create a new temperature value
//
// responses:
//  401: CommonError
//  200: GetTemperature
func (h *BaseHandlerSqlx) PostTemperatureSqlx(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	response := GetTemperature{}

	decoder := json.NewDecoder(r.Body)
	var reqTemperature *models.ReqAddTemperature
	err := decoder.Decode(&reqTemperature)

	if err != nil {
		json.NewEncoder(w).Encode(ErrHandler(lang.Get("invalid_request")))
		return
	}

	rain, errmessage := models.PostTemperatureSqlx(h.db.DB, reqTemperature)
	if errmessage != "" {
		json.NewEncoder(w).Encode(ErrHandler(errmessage))
		return
	}

	response.Status = 1
	response.Message = lang.Get("insert_success")
	response.Data = rain
	json.NewEncoder(w).Encode(response)
}

// swagger:route GET /temperature/showdata/{recordNumber} temperature temperatureShowdata
// Get list of recordNumber of temperature values
//
//     Parameters:
//       + name: recordNumber
//         in: path
//         description: maximum numnber of results to return
//         required: true
//         type: integer
//         format: int32
// responses:
//  401: CommonError
//  200: GetTemperatures
func (h *BaseHandlerSqlx) GetTemperatureShowDataSqlx(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	recordNumber, err := strconv.Atoi(vars["recordNumber"])
	if err != nil {
		// handle error
		fmt.Println(err)

	}

	response := GetTemperatures{}

	temperatures := models.GetTemperatureShowDataSqlx(h.db.DB, recordNumber)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = temperatures

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

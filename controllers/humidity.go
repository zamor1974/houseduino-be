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

// swagger:model GetHumidities
type GetHumidities struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string             `json:"message"`
	Data    *models.Humidities `json:"data"`
}

// swagger:model GetHumidity
type GetHumidity struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string `json:"message"`
	// Umidity value
	Data *models.Humidity `json:"data"`
}

// swagger:route GET /humidity/all humidity humidityAll
// Get humidity list
//
// responses:
//  401: CommonError
//  200: GetHumidities
func (h *BaseHandlerSqlx) GetHumiditiesSqlx(w http.ResponseWriter, r *http.Request) {
	response := GetHumidities{}

	humidities := models.GetHumiditiesSqlx(h.db.DB)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = humidities

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// swagger:route GET /humidity/lasthour humidity humiditylastHour
// Get list of last hour of humidity values .... or the last value inserted
//
// responses:
//  401: CommonError
//  200: GetHumidities
func (h *BaseHandlerSqlx) GetHumiditiesLastHourSqlx(w http.ResponseWriter, r *http.Request) {
	response := GetHumidities{}

	companies := models.GetHumiditiesLastHourSqlx(h.db.DB)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = companies

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// swagger:route POST /humidity/insert humidity addHumidity
// Create a new humidity value
//
// responses:
//  401: CommonError
//  200: GetHumidity
func (h *BaseHandlerSqlx) PostHumiditySqlx(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	response := GetHumidity{}

	decoder := json.NewDecoder(r.Body)
	var reqHumidity *models.ReqAddHumidity
	err := decoder.Decode(&reqHumidity)

	if err != nil {
		json.NewEncoder(w).Encode(ErrHandler(lang.Get("invalid_request")))
		return
	}

	rain, errmessage := models.PostHumiditySqlx(h.db.DB, reqHumidity)
	if errmessage != "" {
		json.NewEncoder(w).Encode(ErrHandler(errmessage))
		return
	}

	response.Status = 1
	response.Message = lang.Get("insert_success")
	response.Data = rain
	json.NewEncoder(w).Encode(response)
}

// swagger:route GET /humidity/showdata/{recordNumber} humidity humidityShowdata
// Get list of recordNumber of humidity values
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
//  200: GetHumidities
func (h *BaseHandlerSqlx) GetHumidityShowDataSqlx(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	recordNumber, err := strconv.Atoi(vars["recordNumber"])
	if err != nil {
		// handle error
		fmt.Println(err)

	}

	response := GetHumidities{}

	humidities := models.GetHumidityShowDataSqlx(h.db.DB, recordNumber)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = humidities

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

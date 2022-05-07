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

// swagger:model GetPlantHumidities
type GetPlantHumidities struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string                  `json:"message"`
	Data    *models.PlantHumidities `json:"data"`
}

// swagger:model GetPlantHumidity
type GetPlantHumidity struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string `json:"message"`
	// PlantUmidity value
	Data *models.PlantHumidity `json:"data"`
}

// swagger:route GET /plant/humidity/all/{id_plant} plant plantHumidityAll
// Get plantHumidity list
//
//     Parameters:
//       + name: id_plant
//         in: path
//         description: id of the plant
//         required: true
//         type: string
//         format: string
// responses:
//  401: CommonError
//  200: GetPlantHumidities
func (h *BaseHandlerSqlx) GetPlantHumiditiesSqlx(w http.ResponseWriter, r *http.Request) {
	response := GetPlantHumidities{}
	vars := mux.Vars(r)

	id_plant := vars["id_plant"]
	humidities := models.GetPlantHumiditiesSqlx(h.db.DB, id_plant)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = humidities

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// swagger:route GET /plant/humidity/lasthour/{id_plant} plant plantHumiditylastHour
// Get list of last hour of plant humidity values .... or the last value inserted
//
//     Parameters:
//       + name: id_plant
//         in: path
//         description: id of the plant
//         required: true
//         type: string
//         format: string
// responses:
//  401: CommonError
//  200: GetPlantHumidities
func (h *BaseHandlerSqlx) GetPlantHumiditiesLastHourSqlx(w http.ResponseWriter, r *http.Request) {
	response := GetPlantHumidities{}
	vars := mux.Vars(r)

	id_plant := vars["id_plant"]

	companies := models.GetPlantHumiditiesLastHourSqlx(h.db.DB, id_plant)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = companies

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// swagger:route POST /plant/humidity/insert plant addPlantHumidity
// Create a new plant humidity value
//
// responses:
//  401: CommonError
//  200: GetPlantHumidity
func (h *BaseHandlerSqlx) PostPlantHumiditySqlx(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	response := GetPlantHumidity{}

	decoder := json.NewDecoder(r.Body)

	var reqHumidity *models.ReqAddPlantHumidity
	err := decoder.Decode(&reqHumidity)

	if err != nil {
		json.NewEncoder(w).Encode(ErrHandler(lang.Get("invalid_request")))
		return
	}

	rain, errmessage := models.PostPlantHumiditySqlx(h.db.DB, reqHumidity)
	if errmessage != "" {
		json.NewEncoder(w).Encode(ErrHandler(errmessage))
		return
	}

	response.Status = 1
	response.Message = lang.Get("insert_success")
	response.Data = rain
	json.NewEncoder(w).Encode(response)
}

// swagger:route GET /plant/humidity/showdata/{id_plant}/{recordNumber} plant plantHumidityShowdata
// Get list of recordNumber of plant humidity values
//
//     Parameters:
//       + name: id_plant
//         in: path
//         description: id of the plant
//         required: true
//         type: string
//         format: string
//       + name: recordNumber
//         in: path
//         description: maximum numnber of results to return
//         required: true
//         type: integer
//         format: int32
// responses:
//  401: CommonError
//  200: GetHumidities
func (h *BaseHandlerSqlx) GetPlantHumidityShowDataSqlx(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id_plant := vars["id_plant"]

	recordNumber, err := strconv.Atoi(vars["recordNumber"])
	if err != nil {
		// handle error
		fmt.Println(err)

	}

	response := GetPlantHumidities{}

	humidities := models.GetPlantHumidityShowDataSqlx(h.db.DB, id_plant, recordNumber)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = humidities

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

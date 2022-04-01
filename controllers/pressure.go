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

// swagger:model GetPressures
type GetPressures struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string            `json:"message"`
	Data    *models.Pressures `json:"data"`
}

// swagger:model GetPressure
type GetPressure struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string `json:"message"`
	// Umidity value
	Data *models.Pressure `json:"data"`
}

// swagger:route GET /pressure/all pressure pressureAll
// Get Pressure list
//
// responses:
//  401: CommonError
//  200: GetPressures
func (h *BaseHandlerSqlx) GetPressuresSqlx(w http.ResponseWriter, r *http.Request) {
	response := GetPressures{}

	pressures := models.GetPressuresSqlx(h.db.DB)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = pressures

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// swagger:route GET /pressure/lasthour pressure pressureLastHour
// Get list of last hour of pressure values .... or the last value inserted
//
// responses:
//  401: CommonError
//  200: GetPressures
func (h *BaseHandlerSqlx) GetPressuresLastHourSqlx(w http.ResponseWriter, r *http.Request) {
	response := GetPressures{}

	pressures := models.GetPressuresLastHourSqlx(h.db.DB)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = pressures

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// swagger:route POST /pressure/insert pressure pressurePost
// Create a new pressure value
//
// responses:
//  401: CommonError
//  200: GetPressure
func (h *BaseHandlerSqlx) PostPressureSqlx(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	response := GetPressure{}

	decoder := json.NewDecoder(r.Body)
	var reqPressure *models.ReqAddPressure
	err := decoder.Decode(&reqPressure)
	fmt.Println(err)

	if err != nil {
		json.NewEncoder(w).Encode(ErrHandler(lang.Get("invalid_request")))
		return
	}

	rain, errmessage := models.PostPressureSqlx(h.db.DB, reqPressure)
	if errmessage != "" {
		json.NewEncoder(w).Encode(ErrHandler(errmessage))
		return
	}

	response.Status = 1
	response.Message = lang.Get("insert_success")
	response.Data = rain
	json.NewEncoder(w).Encode(response)
}

// swagger:route GET /pressure/showdata/{recordNumber} pressure pressureShowdata
// Get list of recordNumber of pressure values
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
//  200: GetPressures
func (h *BaseHandlerSqlx) GetPressureShowDataSqlx(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	recordNumber, err := strconv.Atoi(vars["recordNumber"])
	if err != nil {
		// handle error
		fmt.Println(err)

	}

	response := GetPressures{}

	pressures := models.GetPressureShowDataSqlx(h.db.DB, recordNumber)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = pressures

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

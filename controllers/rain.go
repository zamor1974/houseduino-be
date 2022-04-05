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

// swagger:model GetRains
type GetRains struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string        `json:"message"`
	Data    *models.Rains `json:"data"`
}

// swagger:model GetRain
type GetRain struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string `json:"message"`
	// Companies for this user
	Data *models.Rain `json:"data"`
}

// swagger:route GET /rain/all rain rainAll
// Get rains list
//
// responses:
//  401: CommonError
//  200: GetRains
func (h *BaseHandlerSqlx) GetRainsSqlx(w http.ResponseWriter, r *http.Request) {
	response := GetRains{}

	companies := models.GetRainsSqlx(h.db.DB)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = companies

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// swagger:route GET /rain/lasthour rain rainLastHour
// Get list of last hour of rain .... or the last value inserted
//
// responses:
//  401: CommonError
//  200: GetRains
func (h *BaseHandlerSqlx) GetRainsLastHourSqlx(w http.ResponseWriter, r *http.Request) {
	response := GetRains{}

	companies := models.GetRainsLastHourSqlx(h.db.DB)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = companies

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// swagger:route POST /rain/insert rain addRain
// Create a new rain value
//
// responses:
//  401: CommonError
//  200: GetRain
func (h *BaseHandlerSqlx) PostRainSqlx(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	response := GetRain{}

	decoder := json.NewDecoder(r.Body)
	var reqRain *models.ReqAddRain
	err := decoder.Decode(&reqRain)

	if err != nil {
		json.NewEncoder(w).Encode(ErrHandler(lang.Get("invalid_request")))
		return
	}

	rain, errmessage := models.PostRainSqlx(h.db.DB, reqRain)
	if errmessage != "" {
		json.NewEncoder(w).Encode(ErrHandler(errmessage))
		return
	}

	response.Status = 1
	response.Message = lang.Get("insert_success")
	response.Data = rain
	json.NewEncoder(w).Encode(response)
}

// swagger:route GET /rain/showdata/{recordNumber} rain rainShowdata
// Get list of recordNumber of rain values
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
//  200: GetRains
func (h *BaseHandlerSqlx) GetRainShowDataSqlx(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	recordNumber, err := strconv.Atoi(vars["recordNumber"])
	if err != nil {
		// handle error
		fmt.Println(err)

	}

	response := GetRains{}

	rains := models.GetRainShowDataSqlx(h.db.DB, recordNumber)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = rains

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

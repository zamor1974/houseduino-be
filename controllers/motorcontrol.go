package controllers

import (
	"encoding/json"
	"houseduino-be/lang"
	"houseduino-be/models"
	"net/http"

	"github.com/gorilla/mux"
)

// swagger:model GetMotorStatus
type GetMotorStatus struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string              `json:"message"`
	Data    *models.MotorStatus `json:"data"`
}

// swagger:route GET /motor/status/{id_plant} motor motorStatusId
// Get status of the motor for the plant id
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
//  200: GetMotorStatus
func (h *BaseHandlerSqlx) GetMotorStatus(w http.ResponseWriter, r *http.Request) {
	response := GetMotorStatus{}
	vars := mux.Vars(r)

	id_plant := vars["id_plant"]
	valore := models.GetMotorStatus2(h.db.DB, id_plant)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = valore

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// swagger:route GET /motor/status/all motor motorStatus
// Get status of motor for all the plants
//
// responses:
//  401: CommonError
//  200: GetMotorStatus
func (h *BaseHandlerSqlx) GetMotorStatusSqlx(w http.ResponseWriter, r *http.Request) {
	response := GetMotorStatus{}

	valore := models.GetMotorStatus(h.db.DB)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = valore

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

package controllers

import (
	"encoding/json"
	"houseduino-be/lang"
	"houseduino-be/models"
	"net/http"
)

// swagger:model GetWeather
type GetWeather struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string          `json:"message"`
	Data    *models.Weather `json:"data"`
}

// swagger:route GET /weather/now weather weatherNow
// Get weather data
//
// responses:
//  401: CommonError
//  200: GetWeather
func (h *BaseHandlerSqlx) GetWeatherSqlx(w http.ResponseWriter, r *http.Request) {
	response := GetWeather{}

	companies := models.GetWeatherSqlx(h.db.DB)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = companies

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

package controllers

import (
	"encoding/json"
	"houseduino-be/lang"
	"houseduino-be/models"
	"net/http"
)

// swagger:model GetMessages
type GetMessages struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string           `json:"message"`
	Data    *models.Messages `json:"data"`
}

// swagger:model GetMessage
type GetMessage struct {
	// Status of the error
	// in: int64
	Status int64 `json:"status"`
	// Message of the response
	// in: string
	Message string `json:"message"`
	// Message value
	Data *models.Message `json:"data"`
}

// swagger:route GET /message/lasthour message messageLastHour
// Get list of last hour of Messages .... or the last value inserted
//
// responses:
//  401: CommonError
//  200: GetMessages
func (h *BaseHandlerSqlx) GetMessagesLastHourSqlx(w http.ResponseWriter, r *http.Request) {
	response := GetMessages{}

	messages := models.GetMessagesLastHourSqlx(h.db.DB)

	response.Status = 1
	response.Message = lang.Get("success")
	response.Data = messages

	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// swagger:route POST /message/insert message addMessage
// Create a new Message value
//
// responses:
//  401: CommonError
//  200: GetMessage
func (h *BaseHandlerSqlx) PostMessageSqlx(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	response := GetMessage{}

	decoder := json.NewDecoder(r.Body)
	var reqMessage *models.ReqAddMessage
	err := decoder.Decode(&reqMessage)

	if err != nil {
		json.NewEncoder(w).Encode(ErrHandler(lang.Get("invalid_request")))
		return
	}

	activity, errmessage := models.PostMessageSqlx(h.db.DB, reqMessage)
	if errmessage != "" {
		json.NewEncoder(w).Encode(ErrHandler(errmessage))
		return
	}

	response.Status = 1
	response.Message = lang.Get("insert_success")
	response.Data = activity
	json.NewEncoder(w).Encode(response)
}
